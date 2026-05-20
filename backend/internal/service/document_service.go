package service

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"time"

	"log"
	"github.com/google/uuid"
	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/worker"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"image"
	"image/color"
	"image/draw"
	"image/jpeg"
	"image/png"
	"encoding/json"
	"strings"
	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"
)

type DocumentService struct {
	cfg      config.Config
	repo     repository.Querier
	storage  *infra.StorageService
	asynq    *asynq.Client
	notifSvc *NotificationService
}

func NewDocumentService(cfg config.Config, repo repository.Querier, storage *infra.StorageService, asynqClient *asynq.Client, notifSvc *NotificationService) *DocumentService {
	return &DocumentService{cfg: cfg, repo: repo, storage: storage, asynq: asynqClient, notifSvc: notifSvc}
}

type UploadDocumentParams struct {
	Title        string
	Description  string
	FileName     string
	FileSize     int64
	MimeType     string
	Content      io.Reader
	OwnerID      uuid.UUID
	CompanyID    uuid.UUID
	BranchID     uuid.UUID
	DepartmentID uuid.UUID
	BatchID      uuid.UUID
	TypeID       uuid.UUID
	Sensitivity  string
	Urgency      string
	DocumentDate string
	PageCount    int
	Status       string
}

type UpdateDocumentParams struct {
	ID           uuid.UUID
	Title        string
	Description  string
	TypeID       uuid.UUID
	Sensitivity  string
	Urgency      string
	DocumentDate string
	Metadata     map[string]interface{}
	FileContent  io.Reader
	FileName     string
	FileSize     int64
	MimeType     string
	Status       string
}

func (s *DocumentService) UploadDocument(ctx context.Context, p UploadDocumentParams) (repository.Document, error) {
	log.Printf("[DocumentService] Uploading document: %s (Size: %d, Owner: %s)", p.FileName, p.FileSize, p.OwnerID)

	// 1. Generate unique file path
	objectName := fmt.Sprintf("%s/%s_%s", p.DepartmentID, uuid.New().String(), p.FileName)

	// 2. Check encryption setting from system_settings
	encryptEnabled := false
	encSetting, err := s.repo.GetSystemSetting(ctx, repository.GetSystemSettingParams{
		Category: "storage",
		Key:      "encryption_enabled",
	})
	if err == nil && encSetting.Value.String == "true" {
		encryptEnabled = true
		log.Printf("[DocumentService] AES-256 Encryption (SSE-S3) is ENABLED for this upload")
	}

	// 3. Upload to MinIO if content exists
	if p.Content != nil {
		_, err = s.storage.Upload(ctx, objectName, p.Content, p.FileSize, p.MimeType, encryptEnabled)
		if err != nil {
			log.Printf("[DocumentService] Error uploading to storage: %v", err)
			return repository.Document{}, err
		}
	}

	log.Printf("[DocumentService] Storage upload successful, creating DB record for %s", p.FileName)

	// Prepare metadata JSON
	metadata := map[string]interface{}{
		"urgency":       p.Urgency,
		"document_date": p.DocumentDate,
		"page_count":    p.PageCount,
	}
	metadataBytes, _ := json.Marshal(metadata)

	// 3. Save to DB
	docStatus := p.Status
	if docStatus == "" {
		docStatus = "processing"
	}

	doc, err := s.repo.CreateDocument(ctx, repository.CreateDocumentParams{
		Title:        p.Title,
		Description:  pgtype.Text{String: p.Description, Valid: p.Description != ""},
		FileName:     pgtype.Text{String: p.FileName, Valid: p.FileName != ""},
		FilePath:     pgtype.Text{String: objectName, Valid: objectName != ""},
		FileSize:     pgtype.Int8{Int64: p.FileSize, Valid: p.FileSize > 0},
		MimeType:     pgtype.Text{String: p.MimeType, Valid: p.MimeType != ""},
		OwnerID:      p.OwnerID,
		CompanyID:    p.CompanyID,
		BranchID:     p.BranchID,
		DepartmentID: p.DepartmentID,
		Status:       docStatus,
		BatchID:      pgtype.UUID{Bytes: p.BatchID, Valid: p.BatchID != uuid.Nil},
		Sensitivity:  pgtype.Text{String: strings.ToLower(p.Sensitivity), Valid: p.Sensitivity != ""},
		Metadata:     metadataBytes,
		TypeID:       pgtype.UUID{Bytes: p.TypeID, Valid: p.TypeID != uuid.Nil},
	})
	if err != nil {
		log.Printf("[DocumentService] Error creating document record in DB: %v", err)
		return repository.Document{}, err
	}

	log.Printf("[DocumentService] DB record created successfully: %s (Status: %s)", doc.ID, doc.Status)

	// If it's a draft, stop here (no OCR, no approval)
	if doc.Status == "draft" {
		return doc, nil
	}

	// 4. Enqueue OCR Task
	if s.asynq == nil {
		log.Printf("[DocumentService] WARNING: OCR task skipped - Redis/Asynq client is not initialized")
		return doc, nil
	}

	log.Printf("[DocumentService] Enqueuing OCR task for doc: %s", doc.ID)

	task, err := worker.NewDocumentOCRTask(doc.ID)
	if err != nil {
		log.Printf("[DocumentService] Error creating OCR task: %v", err)
		return doc, fmt.Errorf("failed to create OCR task: %v", err)
	}

	_, err = s.asynq.Enqueue(task)
	if err != nil {
		log.Printf("[DocumentService] Error enqueuing OCR task to Redis: %v", err)
		// We make it non-critical for now to avoid 500 if Redis is buggy
		return doc, nil
	}
	
	log.Printf("[DocumentService] OCR task enqueued successfully for doc: %s", doc.ID)
	
	// 5. Create Approval Workflow Task
	// We automatically assign an approval task to the Department Head
	dept, err := s.repo.GetDepartment(ctx, p.DepartmentID)
	if err == nil && dept.HeadID.Valid {
		log.Printf("[DocumentService] Creating approval task for head of department: %s", dept.HeadID.Bytes)
		_, err = s.repo.CreateApprovalTask(ctx, repository.CreateApprovalTaskParams{
			EntityType: "document_upload",
			EntityID:   doc.ID,
			ApproverID: dept.HeadID.Bytes,
			Level:      1,
		})
		if err != nil {
			log.Printf("[DocumentService] Error creating approval task: %v", err)
		} else {
			// Trigger Notification for Approver
			owner, _ := s.repo.GetUserByID(ctx, doc.OwnerID)
			approveLink := fmt.Sprintf("%s/approvals/%s", s.cfg.AppURL, doc.ID)
			
			meta := map[string]string{
				"docTitle":       doc.Title,
				"departmentName": dept.Name,
				"ownerName":      owner.FullName,
				"approveLink":    approveLink,
			}
			metaJSON, _ := json.Marshal(meta)

			_, _ = s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
				UserID:     dept.HeadID.Bytes,
				Title:      "Permintaan Persetujuan Dokumen",
				Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' memerlukan persetujuan Anda.", doc.Title), Valid: true},
				Type:       "doc-pending-approval",
				EntityType: pgtype.Text{String: "document", Valid: true},
				EntityID:   pgtype.UUID{Bytes: doc.ID, Valid: true},
				Channel:    pgtype.Text{String: "email", Valid: true},
				Metadata:   metaJSON,
			})
		}
	} else {
		log.Printf("[DocumentService] WARNING: No department head found for approval task. Document ID: %s", doc.ID)
	}

	return doc, nil
}


func (s *DocumentService) GetWatermarkedPDF(ctx context.Context, docID uuid.UUID, userFullName string, position string) ([]byte, string, error) {
	log.Printf("[DocumentService] Accessing watermarked PDF: %s (Requested by: %s)", docID, userFullName)

	// 1. Get Doc from DB
	doc, err := s.repo.GetDocument(ctx, docID)
	if err != nil {
		log.Printf("[DocumentService] Error getting document from DB: %v", err)
		return nil, "", err
	}

	// 2. Download from MinIO
	reader, err := s.storage.Download(ctx, doc.FilePath.String)
	if err != nil {
		log.Printf("[DocumentService] Error downloading from storage: %v", err)
		return nil, "", err
	}
	defer reader.Close()

	// 3. Read content
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("[DocumentService] Error reading content: %v", err)
		return nil, "", err
	}

	// 4. Fetch Watermark Settings
	wmNode, err := s.repo.GetIntegrationNodeByType(ctx, "WATERMARK")
	wmType := "text"
	wmText := "CONFIDENTIAL - {user} - {date}"
	wmOpacity := 0.3
	wmPos := "diagonal"

	if err == nil {
		var wmConfig map[string]interface{}
		if err := json.Unmarshal(wmNode.ConfigJson, &wmConfig); err == nil {
			if t, ok := wmConfig["type"].(string); ok { wmType = t }
			if t, ok := wmConfig["text"].(string); ok { wmText = t }
			if o, ok := wmConfig["opacity"].(float64); ok { wmOpacity = o }
			if p, ok := wmConfig["position"].(string); ok { wmPos = p }
		}
	}

	// Replace variables
	watermarkText := strings.ReplaceAll(wmText, "{user}", strings.ToUpper(userFullName))
	watermarkText = strings.ReplaceAll(watermarkText, "{date}", time.Now().Format("2006-01-02"))

	// 4. Apply Watermark based on MimeType (with fallback for octet-stream)
	effectiveMimeType := doc.MimeType.String
	if effectiveMimeType == "application/octet-stream" || effectiveMimeType == "" {
		fileName := strings.ToLower(doc.FileName.String)
		if strings.HasSuffix(fileName, ".pdf") {
			effectiveMimeType = "application/pdf"
		} else if strings.HasSuffix(fileName, ".jpg") || strings.HasSuffix(fileName, ".jpeg") || strings.HasSuffix(fileName, ".png") {
			effectiveMimeType = "image/jpeg" // trigger image watermarking
		}
	}

	if effectiveMimeType != "application/pdf" {
		log.Printf("[DocumentService] Applying image watermark for: %s (MimeType: %s, Type: %s)", docID, effectiveMimeType, wmType)
		watermarkedContent, err := s.applyImageWatermark(content, watermarkText, effectiveMimeType, wmOpacity)
		if err != nil {
			log.Printf("[DocumentService] Warning: Failed to apply image watermark, returning raw: %v", err)
			return content, doc.MimeType.String, nil
		}
		return watermarkedContent, doc.MimeType.String, nil
	}

	// 5. Apply Watermark using pdfcpu for PDFs
	// Configure pdfcpu string based on settings
	rot := "45"
	if wmPos == "center" { rot = "0" }
	
	wmDesc := fmt.Sprintf("font:Helvetica, points:24, scale:0.5, op:%.1f, rot:%s", wmOpacity, rot)
	wm, err := api.TextWatermark(watermarkText, wmDesc, true, false, types.POINTS)
	if err != nil {
		log.Printf("[DocumentService] Error creating watermark: %v", err)
		return nil, "", err
	}

	// Output buffer
	var out bytes.Buffer
	err = api.AddWatermarks(bytes.NewReader(content), &out, nil, wm, nil)
	if err != nil {
		log.Printf("[DocumentService] Error applying watermark: %v", err)
		// Fallback to raw content if watermarking fails but it's a PDF
		return content, doc.MimeType.String, nil
	}

	log.Printf("[DocumentService] Watermarked PDF generated: %s", docID)
	return out.Bytes(), doc.MimeType.String, nil
}

type BatchDetails struct {
	Batch     repository.ProcessingBatch `json:"batch"`
	Documents []repository.Document      `json:"documents"`
}

func (s *DocumentService) CreateBatch(ctx context.Context, userID uuid.UUID, totalFiles int) (repository.ProcessingBatch, error) {
	return s.repo.CreateBatch(ctx, repository.CreateBatchParams{
		UserID:     userID,
		TotalFiles: int32(totalFiles),
	})
}

func (s *DocumentService) GetBatchDetails(ctx context.Context, batchID uuid.UUID) (*BatchDetails, error) {
	batch, err := s.repo.GetBatch(ctx, batchID)
	if err != nil {
		return nil, err
	}

	docs, err := s.repo.GetDocumentsByBatch(ctx, pgtype.UUID{Bytes: batchID, Valid: true})
	if err != nil {
		return nil, err
	}

	return &BatchDetails{
		Batch:     batch,
		Documents: docs,
	}, nil
}
func (s *DocumentService) ListOCRHistory(ctx context.Context, page, pageSize int) ([]repository.ListOCRJobsRow, int64, error) {
	offset := (page - 1) * pageSize
	rows, err := s.repo.ListOCRJobs(ctx, repository.ListOCRJobsParams{
		Limit:  int32(pageSize),
		Offset: int32(offset),
	})
	if err != nil {
		return nil, 0, err
	}

	total, err := s.repo.CountOCRJobs(ctx)
	if err != nil {
		return nil, 0, err
	}

	return rows, total, nil
}

func (s *DocumentService) GetDocument(ctx context.Context, id uuid.UUID) (repository.GetDocumentWithDetailsRow, error) {
	return s.repo.GetDocumentWithDetails(ctx, id)
}

func (s *DocumentService) ListRecentDocuments(ctx context.Context, limit int) ([]repository.ListRecentDocumentsRow, error) {
	return s.repo.ListRecentDocuments(ctx, repository.ListRecentDocumentsParams{
		Limit:  int32(limit),
		Offset: 0,
	})
}

func (s *DocumentService) ListRecentDocumentsByOwner(ctx context.Context, ownerID uuid.UUID, limit int) ([]repository.ListRecentDocumentsByOwnerRow, error) {
	return s.repo.ListRecentDocumentsByOwner(ctx, repository.ListRecentDocumentsByOwnerParams{
		OwnerID: ownerID,
		Limit:   int32(limit),
		Offset:  0,
	})
}

func (s *DocumentService) GetOCRJob(ctx context.Context, docID uuid.UUID) (repository.OcrJob, error) {
	return s.repo.GetOCRJobByEntity(ctx, repository.GetOCRJobByEntityParams{
		EntityType: "document",
		EntityID:   docID,
	})
}
func (s *DocumentService) GetUser(ctx context.Context, userID uuid.UUID) (repository.GetUserByIDRow, error) {
	return s.repo.GetUserByID(ctx, userID)
}

func (s *DocumentService) GetDepartment(ctx context.Context, id uuid.UUID) (repository.Department, error) {
	return s.repo.GetDepartment(ctx, id)
}

func (s *DocumentService) GetBranch(ctx context.Context, id uuid.UUID) (repository.Branch, error) {
	return s.repo.GetBranch(ctx, id)
}

func (s *DocumentService) applyImageWatermark(content []byte, text string, mimeType string, opacity float64) ([]byte, error) {
	// 1. Decode image
	src, _, err := image.Decode(bytes.NewReader(content))
	if err != nil {
		return nil, err
	}

	bounds := src.Bounds()
	w, h := bounds.Dx(), bounds.Dy()

	// 2. Create a new RGBA image
	rgba := image.NewRGBA(bounds)
	draw.Draw(rgba, bounds, src, bounds.Min, draw.Src)

	// 3. Configure font color with dynamic opacity
	alpha := uint8(255 * opacity)
	col := color.RGBA{255, 255, 255, alpha} // Semi-transparent white
	
	d := &font.Drawer{
		Dst:  rgba,
		Src:  image.NewUniform(col),
		Face: basicfont.Face7x13,
	}

	// 4. Draw multiple watermarks in a grid
	// Adjust spacing based on image size
	stepX := w / 3
	if stepX < 300 { stepX = 300 }
	stepY := h / 4
	if stepY < 300 { stepY = 300 }

	for x := 100; x < w; x += stepX {
		for y := 100; y < h; y += stepY {
			d.Dot = fixed.P(x, y)
			d.DrawString(text)
		}
	}

	// 5. Encode back to bytes
	var buf bytes.Buffer
	var encodeErr error
	if mimeType == "image/png" {
		encodeErr = png.Encode(&buf, rgba)
	} else {
		encodeErr = jpeg.Encode(&buf, rgba, &jpeg.Options{Quality: 85})
	}

	if encodeErr != nil {
		return nil, encodeErr
	}

	return buf.Bytes(), nil
}
func (s *DocumentService) GetLoanHistory(ctx context.Context, docID uuid.UUID) ([]repository.GetDocumentLoanHistoryRow, error) {
	return s.repo.GetDocumentLoanHistory(ctx, docID)
}
func (s *DocumentService) ApproveDocument(ctx context.Context, docID uuid.UUID, notes string) error {
	// 1. Update workflow status
	if err := s.repo.ApproveTask(ctx, repository.ApproveTaskParams{
		EntityID:     docID,
		DecisionNote: pgtype.Text{String: notes, Valid: notes != ""},
	}); err != nil {
		return err
	}

	// 2. Update document status to 'approved' (waiting physical) and physical_status to 'pending'
	err := s.repo.UpdateDocumentStatus(ctx, repository.UpdateDocumentStatusParams{
		ID:     docID,
		Status: "approved",
	})
	if err != nil {
		return err
	}

	if err := s.repo.UpdateDocumentPhysicalStatus(ctx, repository.UpdateDocumentPhysicalStatusParams{
		ID:             docID,
		PhysicalStatus: pgtype.Text{String: "pending", Valid: true},
	}); err != nil {
		return err
	}

	// 3. Trigger Notifications (Async)
	go func() {
		bgCtx := context.Background()
		doc, err := s.repo.GetDocument(bgCtx, docID)
		if err != nil {
			log.Printf("[DocumentService] Error getting doc for approval notification: %v", err)
			return
		}
		owner, _ := s.repo.GetUserByID(bgCtx, doc.OwnerID)

		// 3.1 Notify Owner
		metaOwner := map[string]string{
			"docTitle": doc.Title,
			"notes":    notes,
		}
		metaOwnerJSON, _ := json.Marshal(metaOwner)
		_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
			UserID:     doc.OwnerID,
			Title:      "Dokumen Disetujui",
			Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' Anda telah disetujui.", doc.Title), Valid: true},
			Type:       "doc-approved",
			EntityType: pgtype.Text{String: "document", Valid: true},
			EntityID:   pgtype.UUID{Bytes: docID, Valid: true},
			Channel:    pgtype.Text{String: "email", Valid: true},
			Metadata:   metaOwnerJSON,
		})

		// 3.2 Notify DC Admins & Superadmin (Ready for Intake)
		roles := []string{"superadmin", "admin doc controller", "kepala doc controller"}
		admins, err := s.repo.ListUsersByRoles(bgCtx, roles)
		if err == nil {
			for _, admin := range admins {
				metaDC := map[string]string{
					"docTitle":  doc.Title,
					"ownerName": owner.FullName,
				}
				metaDCJSON, _ := json.Marshal(metaDC)
				_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
					UserID:     admin.ID,
					Title:      "Dokumen Siap Intake",
					Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' telah disetujui. Silakan lakukan physical intake.", doc.Title), Valid: true},
					Type:       "doc-approved-dc",
					EntityType: pgtype.Text{String: "document", Valid: true},
					EntityID:   pgtype.UUID{Bytes: docID, Valid: true},
					Channel:    pgtype.Text{String: "email", Valid: true},
					Metadata:   metaDCJSON,
				})
			}
		}
	}()

	return nil
}

func (s *DocumentService) RejectDocument(ctx context.Context, docID uuid.UUID, reason string, notes string) error {
	// 1. Update workflow status
	if err := s.repo.RejectTask(ctx, repository.RejectTaskParams{
		EntityID:        docID,
		DecisionNote:    pgtype.Text{String: notes, Valid: notes != ""},
		RejectionReason: pgtype.Text{String: reason, Valid: reason != ""},
	}); err != nil {
		return err
	}

	// 2. Update document status to rejected
	err := s.repo.UpdateDocumentStatus(ctx, repository.UpdateDocumentStatusParams{
		ID:     docID,
		Status: "rejected",
	})
	if err != nil {
		return err
	}

	// 3. Trigger Notification for Owner (Async)
	go func() {
		bgCtx := context.Background()
		doc, err := s.repo.GetDocument(bgCtx, docID)
		if err != nil {
			log.Printf("[DocumentService] Error getting doc for rejection notification: %v", err)
			return
		}

		meta := map[string]string{
			"docTitle": doc.Title,
			"notes":    fmt.Sprintf("%s. %s", reason, notes),
		}
		metaJSON, _ := json.Marshal(meta)
		_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
			UserID:     doc.OwnerID,
			Title:      "Dokumen Ditolak",
			Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' Anda ditolak oleh Manager.", doc.Title), Valid: true},
			Type:       "doc-rejected",
			EntityType: pgtype.Text{String: "document", Valid: true},
			EntityID:   pgtype.UUID{Bytes: docID, Valid: true},
			Channel:    pgtype.Text{String: "email", Valid: true},
			Metadata:   metaJSON,
		})
	}()

	return nil
}
func (s *DocumentService) BulkApproveDocuments(ctx context.Context, ids []uuid.UUID) error {
	for _, id := range ids {
		if err := s.ApproveDocument(ctx, id, "Bulk approved"); err != nil {
			return err
		}
	}
	return nil
}

func (s *DocumentService) BulkRejectDocuments(ctx context.Context, ids []uuid.UUID) error {
	for _, id := range ids {
		if err := s.RejectDocument(ctx, id, "Bulk Action", "Bulk rejected"); err != nil {
			return err
		}
	}
	return nil
}

func (s *DocumentService) UpdateDocument(ctx context.Context, params UpdateDocumentParams) (repository.Document, error) {
	// 1. Get existing document to merge metadata
	existing, err := s.repo.GetDocument(ctx, params.ID)
	if err != nil {
		return repository.Document{}, err
	}

	metadata := make(map[string]interface{})
	if len(existing.Metadata) > 0 {
		json.Unmarshal(existing.Metadata, &metadata)
	}

	// 2. Update metadata fields
	metadata["urgency"] = params.Urgency
	metadata["document_date"] = params.DocumentDate
	for k, v := range params.Metadata {
		metadata[k] = v
	}

	metaJSON, _ := json.Marshal(metadata)

	// 3. Handle file replacement if provided
	filePath := ""
	if params.FileContent != nil {
		objectName := fmt.Sprintf("%s/%s_%s", existing.DepartmentID, uuid.New().String(), params.FileName)
		_, err := s.storage.Upload(ctx, objectName, params.FileContent, params.FileSize, params.MimeType, false)
		if err != nil {
			return repository.Document{}, err
		}
		filePath = objectName

		// Cleanup: Delete old file if path is different or even if same to be safe
		if existing.FilePath.String != "" && existing.FilePath.String != filePath {
			s.storage.Delete(ctx, existing.FilePath.String)
		}
	}

	// 4. Determine status: if file replaced, set to 'processing' for OCR, otherwise 'pending'
	docStatus := params.Status
	if docStatus == "" {
		docStatus = "pending"
		if filePath != "" {
			docStatus = "processing"
		}
	}

	// 5. Perform update
	doc, err := s.repo.UpdateDocument(ctx, repository.UpdateDocumentParams{
		ID:          params.ID,
		Title:       params.Title,
		Description: pgtype.Text{String: params.Description, Valid: params.Description != ""},
		TypeID:      pgtype.UUID{Bytes: params.TypeID, Valid: true},
		Sensitivity: pgtype.Text{String: params.Sensitivity, Valid: params.Sensitivity != ""},
		Metadata:    metaJSON,
		FileName:    params.FileName,
		FilePath:    filePath,
		FileSize:    params.FileSize,
		MimeType:    params.MimeType,
		Status:      docStatus,
	})
	if err != nil {
		return repository.Document{}, err
	}

	// If it's still a draft, stop here (no OCR, no approval)
	if doc.Status == "draft" {
		return doc, nil
	}

	// 6. Handle OCR if file replaced
	if filePath != "" && s.asynq != nil {
		log.Printf("[DocumentService] Enqueuing OCR task for revised doc: %s", doc.ID)
		task, _ := worker.NewDocumentOCRTask(doc.ID)
		s.asynq.Enqueue(task)
	}

	// 7. Create Approval Workflow Task for Department Head
	dept, err := s.repo.GetDepartment(ctx, doc.DepartmentID)
	if err == nil && dept.HeadID.Valid {
		log.Printf("[DocumentService] Creating revision approval task for head of department: %s", dept.HeadID.Bytes)
		_, err = s.repo.CreateApprovalTask(ctx, repository.CreateApprovalTaskParams{
			EntityType: "document_upload", // Use standard type
			EntityID:   doc.ID,
			ApproverID: dept.HeadID.Bytes,
			Level:      1,
		})
		if err != nil {
			log.Printf("[DocumentService] Error creating new approval task for revision: %v", err)
			return doc, fmt.Errorf("failed to create approval task: %v", err)
		}

		// Trigger Notification for Approver (Revision)
		owner, _ := s.repo.GetUserByID(ctx, doc.OwnerID)
		approveLink := fmt.Sprintf("%s/approvals/%s", s.cfg.AppURL, doc.ID)
		
		meta := map[string]string{
			"docTitle":       doc.Title,
			"departmentName": dept.Name,
			"ownerName":      owner.FullName,
			"approveLink":    approveLink,
		}
		metaJSON, _ := json.Marshal(meta)

		_, _ = s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
			UserID:     dept.HeadID.Bytes,
			Title:      "Revisi Dokumen Memerlukan Persetujuan",
			Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' telah direvisi dan memerlukan persetujuan ulang.", doc.Title), Valid: true},
			Type:       "doc-pending-approval",
			EntityType: pgtype.Text{String: "document", Valid: true},
			EntityID:   pgtype.UUID{Bytes: doc.ID, Valid: true},
			Channel:    pgtype.Text{String: "email", Valid: true},
			Metadata:   metaJSON,
		})
	} else {
		log.Printf("[DocumentService] WARNING: No department head found for revision approval. Document ID: %s", doc.ID)
	}

	// 8. Log activity
	s.repo.CreateActivityLog(ctx, repository.CreateActivityLogParams{
		UserID:     existing.OwnerID,
		Action:     "REVISION_SUBMITTED",
		EntityType: "document",
		EntityID:   pgtype.UUID{Bytes: doc.ID, Valid: true},
		Details:    []byte(`{"note": "Document resubmitted after revision"}`),
	})

	return doc, nil
}

func (s *DocumentService) GetRawFile(ctx context.Context, path string) ([]byte, string, error) {
	reader, err := s.storage.Download(ctx, path)
	if err != nil {
		return nil, "", err
	}
	defer reader.Close()

	data, err := io.ReadAll(reader)
	if err != nil {
		return nil, "", err
	}

	// Simple mime type detection based on extension
	mime := "application/octet-stream"
	if strings.HasSuffix(strings.ToLower(path), ".jpg") || strings.HasSuffix(strings.ToLower(path), ".jpeg") {
		mime = "image/jpeg"
	} else if strings.HasSuffix(strings.ToLower(path), ".png") {
		mime = "image/png"
	} else if strings.HasSuffix(strings.ToLower(path), ".pdf") {
		mime = "application/pdf"
	}

	return data, mime, nil
}

func (s *DocumentService) GetSystemSettingsByCategory(ctx context.Context, category string) ([]repository.SystemSetting, error) {
	return s.repo.GetSystemSettingsByCategory(ctx, category)
}

func (s *DocumentService) UpsertSystemSetting(ctx context.Context, arg repository.UpsertSystemSettingParams) (repository.SystemSetting, error) {
	return s.repo.UpsertSystemSetting(ctx, arg)
}

type TreeNode struct {
	ID       string     `json:"id"`
	Name     string     `json:"name"`
	Type     string     `json:"type"`
	Expanded bool       `json:"expanded"`
	Active   bool       `json:"active"`
	Count    int        `json:"count,omitempty"`
	Children []TreeNode `json:"children,omitempty"`
}

func (s *DocumentService) GetExplorerTree(ctx context.Context) ([]TreeNode, error) {
	// 1. Get Hierarchy Config
	setting, err := s.repo.GetSystemSetting(ctx, repository.GetSystemSettingParams{
		Category: "explorer",
		Key:      "folder_path",
	})
	path := "company,branch,department,year,rack,box,ordner"
	if err == nil {
		path = setting.Value.String
	}
	levels := strings.Split(path, ",")

	// 2. Get Counts
	counts, err := s.repo.GetDocumentHierarchyCounts(ctx)
	if err != nil {
		return nil, err
	}

	// 3. Fetch names for all entities
	companies, _ := s.repo.ListCompanies(ctx)
	branches, _ := s.repo.ListAllBranchesGlobal(ctx)
	departments, _ := s.repo.ListAllDepartments(ctx)
	racks, _ := s.repo.ListAllRacksGlobal(ctx)
	boxes, _ := s.repo.ListAllBoxesGlobal(ctx)
	ordners, _ := s.repo.ListAllOrdnersGlobal(ctx)
	types, _ := s.repo.ListDocumentTypes(ctx)

	// Build name maps
	compMap := make(map[uuid.UUID]string)
	for _, c := range companies { compMap[c.ID] = c.Name }
	branchMap := make(map[uuid.UUID]string)
	for _, b := range branches { branchMap[b.ID] = b.Name }
	deptMap := make(map[uuid.UUID]string)
	for _, d := range departments { deptMap[d.ID] = d.Name }
	rackMap := make(map[uuid.UUID]string)
	for _, r := range racks { rackMap[r.ID] = r.Name }
	boxMap := make(map[uuid.UUID]string)
	for _, b := range boxes { boxMap[b.ID] = b.Name }
	ordnerMap := make(map[uuid.UUID]string)
	for _, o := range ordners { ordnerMap[o.ID] = o.Name }
	typeMap := make(map[uuid.UUID]string)
	for _, t := range types { typeMap[t.ID] = t.Name }

	// 4. Build Tree
	// We use a map to store nodes at each level to avoid duplicates under the same parent
	type nodeKey struct {
		level    int
		id       string
		parentID string
	}
	nodeCache := make(map[nodeKey]*TreeNode)

	root := &TreeNode{ID: "root", Name: "ARSIP", Type: "root", Expanded: true, Children: []TreeNode{}}

	for _, c := range counts {
		currentNode := root
		for i, level := range levels {
			var id, name string
			switch level {
			case "company":
				if c.CompanyID == uuid.Nil { continue }
				id = c.CompanyID.String()
				name = compMap[c.CompanyID]
			case "branch":
				if c.BranchID == uuid.Nil { continue }
				id = c.BranchID.String()
				name = branchMap[c.BranchID]
			case "department":
				if c.DepartmentID == uuid.Nil { continue }
				id = c.DepartmentID.String()
				name = deptMap[c.DepartmentID]
			case "year":
				if c.DocYear == 0 { continue }
				id = fmt.Sprintf("year-%d", c.DocYear)
				name = fmt.Sprintf("%d", c.DocYear)
			case "rack":
				if !c.RackID.Valid { continue }
				id = uuid.UUID(c.RackID.Bytes).String()
				name = rackMap[c.RackID.Bytes]
			case "box":
				if !c.BoxID.Valid { continue }
				id = uuid.UUID(c.BoxID.Bytes).String()
				name = boxMap[c.BoxID.Bytes]
			case "ordner":
				if !c.OrdnerID.Valid { continue }
				id = uuid.UUID(c.OrdnerID.Bytes).String()
				name = ordnerMap[c.OrdnerID.Bytes]
			case "type":
				if !c.TypeID.Valid { continue }
				id = uuid.UUID(c.TypeID.Bytes).String()
				name = typeMap[c.TypeID.Bytes]
			}

			if id == "" { continue }

			key := nodeKey{level: i, id: id, parentID: currentNode.ID}
			if node, ok := nodeCache[key]; ok {
				node.Count += int(c.DocCount)
				currentNode = node
			} else {
				newNode := &TreeNode{
					ID:    id,
					Name:  name,
					Type:  level,
					Count: int(c.DocCount),
				}
				nodeCache[key] = newNode
				currentNode.Children = append(currentNode.Children, *newNode)
				// Since we append a copy, we need to get the pointer to the copy in the slice
				currentNode = &currentNode.Children[len(currentNode.Children)-1]
				// Update cache to point to the one in the slice
				nodeCache[key] = currentNode
			}
		}
	}

	return root.Children, nil
}

func (s *DocumentService) FilterDocuments(ctx context.Context, params repository.SearchDocumentsParams) ([]repository.SearchDocumentsRow, error) {
	return s.repo.SearchDocuments(ctx, params)
}

// ============================================================================
// Loan Request Service Methods
// ============================================================================

type CreateLoanRequestParams struct {
	UserID       uuid.UUID
	DocumentIDs  []uuid.UUID
	Purpose      string
	DurationDays int32
	Notes        string
	PickupDate   string
	TimeSlot     string
}

func (s *DocumentService) CreateLoanRequest(ctx context.Context, p CreateLoanRequestParams) (repository.LoanRequest, error) {
	log.Printf("[DocumentService] Creating loan request for user %s with %d documents", p.UserID, len(p.DocumentIDs))

	// 1. Auto-detect User's org hierarchy
	user, err := s.repo.GetUserByID(ctx, p.UserID)
	if err != nil {
		return repository.LoanRequest{}, fmt.Errorf("failed to get user: %v", err)
	}

	var companyID, branchID, departmentID pgtype.UUID
	if user.DepartmentID.Valid {
		departmentID = user.DepartmentID
		dept, err := s.repo.GetDepartment(ctx, user.DepartmentID.Bytes)
		if err == nil {
			branchID = pgtype.UUID{Bytes: dept.BranchID, Valid: true}
			branch, err := s.repo.GetBranch(ctx, dept.BranchID)
			if err == nil {
				companyID = pgtype.UUID{Bytes: branch.CompanyID, Valid: true}
			}
		}
	}

	// 2. Generate request number: LOAN-YYYY-MM-NNNNN
	nextSeq, err := s.repo.GetNextLoanRequestNo(ctx)
	if err != nil {
		nextSeq = 1
	}
	requestNo := fmt.Sprintf("LOAN-%s-%05d", time.Now().Format("2006-01"), nextSeq)

	// 3. Create loan request
	loanReq, err := s.repo.CreateLoanRequest(ctx, repository.CreateLoanRequestParams{
		RequestNo:    requestNo,
		UserID:       p.UserID,
		Purpose:      p.Purpose,
		DurationDays: p.DurationDays,
		Notes:        pgtype.Text{String: p.Notes, Valid: p.Notes != ""},
		CompanyID:    companyID,
		BranchID:     branchID,
		DepartmentID: departmentID,
	})
	if err != nil {
		log.Printf("[DocumentService] Error creating loan request: %v", err)
		return repository.LoanRequest{}, fmt.Errorf("failed to create loan request: %v", err)
	}

	log.Printf("[DocumentService] Loan request created: %s (%s)", loanReq.ID, requestNo)

	// 4. Create loan request items
	for _, docID := range p.DocumentIDs {
		_, err := s.repo.CreateLoanRequestItem(ctx, repository.CreateLoanRequestItemParams{
			LoanRequestID: loanReq.ID,
			DocumentID:    docID,
			Category:      pgtype.Text{String: "general", Valid: true},
			Sensitivity:   pgtype.Text{String: "internal", Valid: true},
			Method:        pgtype.Text{String: "physical", Valid: true},
		})
		if err != nil {
			log.Printf("[DocumentService] Error adding loan item (doc: %s): %v", docID, err)
			// Continue to next item, don't fail the whole request
		}
	}

	// 5. Create Approval Task for Department Head
	if departmentID.Valid {
		dept, err := s.repo.GetDepartment(ctx, departmentID.Bytes)
		if err == nil && dept.HeadID.Valid {
			log.Printf("[DocumentService] Creating loan approval task for dept head: %s", dept.HeadID.Bytes)
			_, err = s.repo.CreateApprovalTask(ctx, repository.CreateApprovalTaskParams{
				EntityType: "loan_request",
				EntityID:   loanReq.ID,
				ApproverID: dept.HeadID.Bytes,
				Level:      1,
			})
			if err != nil {
				log.Printf("[DocumentService] Error creating approval task: %v", err)
			} else {
				// 6. Notify Approver
				meta := map[string]string{
					"requestNo": requestNo,
					"userName":  user.FullName,
					"purpose":   p.Purpose,
				}
				metaJSON, _ := json.Marshal(meta)

				_, _ = s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
					UserID:     dept.HeadID.Bytes,
					Title:      "Permintaan Peminjaman Dokumen",
					Body:       pgtype.Text{String: fmt.Sprintf("%s mengajukan peminjaman %d dokumen (%s)", user.FullName, len(p.DocumentIDs), requestNo), Valid: true},
					Type:       "loan-pending-approval",
					EntityType: pgtype.Text{String: "loan_request", Valid: true},
					EntityID:   pgtype.UUID{Bytes: loanReq.ID, Valid: true},
					Channel:    pgtype.Text{String: "email", Valid: true},
					Metadata:   metaJSON,
				})
			}
		}
	}

	return loanReq, nil
}

func (s *DocumentService) ListUserLoanRequests(ctx context.Context, userID uuid.UUID) ([]repository.ListUserLoanRequestsRow, error) {
	return s.repo.ListUserLoanRequests(ctx, userID)
}

func (s *DocumentService) ListAllLoanRequests(ctx context.Context, limit int32) ([]repository.ListAllLoanRequestsRow, error) {
	return s.repo.ListAllLoanRequests(ctx, limit)
}

func (s *DocumentService) GetLoanRequest(ctx context.Context, id uuid.UUID) (repository.GetLoanRequestRow, error) {
	return s.repo.GetLoanRequest(ctx, id)
}

func (s *DocumentService) GetLoanRequestItems(ctx context.Context, loanRequestID uuid.UUID) ([]repository.GetLoanRequestItemsRow, error) {
	return s.repo.GetLoanRequestItems(ctx, loanRequestID)
}

func (s *DocumentService) ApproveLoanRequest(ctx context.Context, loanID uuid.UUID, approverID uuid.UUID) error {
	loan, err := s.repo.GetLoanRequest(ctx, loanID)
	if err != nil {
		return err
	}

	if loan.Status == "pending" || loan.Status == "returned" {
		if err := s.repo.ApproveLoanRequestL1(ctx, repository.ApproveLoanRequestL1Params{
			ID:           loanID,
			L1ApprovedBy: pgtype.UUID{Bytes: approverID, Valid: true},
		}); err != nil {
			return err
		}
	} else if loan.Status == "l1_approved" {
		if err := s.repo.UpdateLoanRequestStatus(ctx, repository.UpdateLoanRequestStatusParams{
			ID:     loanID,
			Status: "active",
		}); err != nil {
			return err
		}
	}

	// Update approval workflow
	_ = s.repo.ApproveTask(ctx, repository.ApproveTaskParams{
		EntityID:     loanID,
		DecisionNote: pgtype.Text{String: "Loan request approved", Valid: true},
	})

	// Notify requester
	go func() {
		bgCtx := context.Background()
		loan, err := s.repo.GetLoanRequest(bgCtx, loanID)
		if err != nil {
			return
		}
		meta := map[string]string{"requestNo": loan.RequestNo}
		metaJSON, _ := json.Marshal(meta)

		_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
			UserID:     loan.UserID,
			Title:      "Peminjaman Disetujui",
			Body:       pgtype.Text{String: fmt.Sprintf("Permohonan peminjaman %s telah disetujui.", loan.RequestNo), Valid: true},
			Type:       "loan-approved",
			EntityType: pgtype.Text{String: "loan_request", Valid: true},
			EntityID:   pgtype.UUID{Bytes: loanID, Valid: true},
			Channel:    pgtype.Text{String: "email", Valid: true},
			Metadata:   metaJSON,
		})
	}()

	return nil
}

func (s *DocumentService) RejectLoanRequestAction(ctx context.Context, loanID uuid.UUID, approverID uuid.UUID, reason string) error {
	loan, err := s.repo.GetLoanRequest(ctx, loanID)
	if err != nil {
		return err
	}

	if loan.Status == "pending" || loan.Status == "returned" {
		if err := s.repo.RejectLoanRequest(ctx, repository.RejectLoanRequestParams{
			ID:                loanID,
			L1ApprovedBy:      pgtype.UUID{Bytes: approverID, Valid: true},
			L1RejectionReason: pgtype.Text{String: reason, Valid: true},
		}); err != nil {
			return err
		}
	} else if loan.Status == "l1_approved" {
		if err := s.repo.RejectLoanRequestL2(ctx, repository.RejectLoanRequestL2Params{
			ID:                loanID,
			L2ApprovedBy:      pgtype.UUID{Bytes: approverID, Valid: true},
			L2RejectionReason: pgtype.Text{String: reason, Valid: true},
		}); err != nil {
			return err
		}
		
		// Reset the L1 task so the manager can review it again
		if loan.L1ApprovedBy.Valid {
			_ = s.repo.ResetTaskToPending(ctx, repository.ResetTaskToPendingParams{
				EntityID:        loanID,
				ApproverID:      loan.L1ApprovedBy.Bytes,
				RejectionReason: pgtype.Text{String: reason, Valid: true},
			})
		}
	}

	// Update approval workflow
	_ = s.repo.RejectTask(ctx, repository.RejectTaskParams{
		EntityID:        loanID,
		DecisionNote:    pgtype.Text{String: "Loan request rejected", Valid: true},
		RejectionReason: pgtype.Text{String: reason, Valid: true},
	})

	// Notify requester and/or L1 approver
	go func() {
		bgCtx := context.Background()
		
		meta := map[string]string{"requestNo": loan.RequestNo, "reason": reason}
		metaJSON, _ := json.Marshal(meta)

		if loan.Status == "l1_approved" {
			// Notify L1 approver (manager)
			if loan.L1ApprovedBy.Valid {
				_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
					UserID:     loan.L1ApprovedBy.Bytes,
					Title:      "Peminjaman Dikembalikan (Kepala DC)",
					Body:       pgtype.Text{String: fmt.Sprintf("Permohonan peminjaman %s dikembalikan oleh Kepala DC. Alasan: %s", loan.RequestNo, reason), Valid: true},
					Type:       "loan-rejected",
					EntityType: pgtype.Text{String: "loan_request", Valid: true},
					EntityID:   pgtype.UUID{Bytes: loanID, Valid: true},
					Channel:    pgtype.Text{String: "email", Valid: true},
					Metadata:   metaJSON,
				})
			}
		} else {
			// Notify requester
			_, _ = s.notifSvc.CreateNotification(bgCtx, repository.CreateNotificationParams{
				UserID:     loan.UserID,
				Title:      "Peminjaman Ditolak",
				Body:       pgtype.Text{String: fmt.Sprintf("Permohonan peminjaman %s ditolak. Alasan: %s", loan.RequestNo, reason), Valid: true},
				Type:       "loan-rejected",
				EntityType: pgtype.Text{String: "loan_request", Valid: true},
				EntityID:   pgtype.UUID{Bytes: loanID, Valid: true},
				Channel:    pgtype.Text{String: "email", Valid: true},
				Metadata:   metaJSON,
			})
		}
	}()

	return nil
}

