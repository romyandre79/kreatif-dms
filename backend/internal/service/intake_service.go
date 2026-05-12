package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
	"errors"
)

type IntakeService struct {
	repo     repository.Querier
	notifSvc *NotificationService
}

func NewIntakeService(repo repository.Querier, notifSvc *NotificationService) *IntakeService {
	return &IntakeService{repo: repo, notifSvc: notifSvc}
}

type ManifestItem struct {
	ID             uuid.UUID              `json:"id"`
	DocumentID     uuid.UUID              `json:"document_id"`
	Title          string                 `json:"title"`
	TypeID         *uuid.UUID             `json:"type_id"`
	TypeName       string                 `json:"type_name"`
	CategoryID     *uuid.UUID             `json:"category_id"`
	Date           string                 `json:"date"`
	Status         string                 `json:"status"`
	PhysicalStatus string                 `json:"physical_status"`
	PreviewPath    string                 `json:"preview_path"`
	Metadata       map[string]interface{} `json:"metadata"`
	ExtractedText  string                 `json:"extracted_text"`
	Notes          string                 `json:"notes"`
}

type TimelineItem struct {
	Status    string `json:"status"`
	Timestamp string `json:"timestamp"`
	IsActive  bool   `json:"is_active"`
}

type ManifestDetail struct {
	ID             uuid.UUID      `json:"id"`
	ManifestNo     string         `json:"manifest_no"`
	Title          string         `json:"title"`
	TypeName       string         `json:"type_name"`
	OwnerName      string         `json:"owner_name"`
	DepartmentName string         `json:"department_name"`
	Status         string         `json:"status"`
	TotalItems     int32          `json:"total_items"`
	Priority       string         `json:"priority"`
	Estimate       string         `json:"estimate"`
	CreatedAt      string         `json:"created_at"`
	ReceivedAt     string         `json:"received_at,omitempty"`
	Source         string         `json:"source"` // 'INB' or 'MIG'
	Items          []ManifestItem `json:"items"`
	Timeline       []TimelineItem `json:"timeline"`
}

type LabelingDocument struct {
	DocumentID     uuid.UUID `json:"document_id"`
	Title          string    `json:"title"`
	DepartmentName string    `json:"department_name"`
	ManifestNo     string    `json:"manifest_no"`
	Status         string    `json:"status"`
}

type LabelingStats struct {
	WaitingCount int64 `json:"waiting_count"`
	PrintedCount int64 `json:"printed_count"`
	StoredCount  int64 `json:"stored_count"`
}

type BoxInfo struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	RackName       string    `json:"rack_name"`
	DepartmentName string    `json:"department_name"`
	CompanyName    string    `json:"company_name"`
	BranchName     string    `json:"branch_name"`
	LocationDetail string    `json:"location_detail"`
}

func (s *IntakeService) GetDocumentByManifestID(ctx context.Context, idStr string) (*ManifestDetail, error) {
	// 1. Try parsing as UUID first (for direct links from Batch Area)
	if manifestID, err := uuid.Parse(idStr); err == nil {
		manifest, err := s.repo.GetManifest(ctx, manifestID)
		if err == nil {
			items, _ := s.repo.GetManifestItems(ctx, manifest.ID)
			manifestItems := make([]ManifestItem, 0)
			for _, item := range items {
				var typeID *uuid.UUID
				if item.DocumentTypeID.Valid {
					id := uuid.UUID(item.DocumentTypeID.Bytes)
					typeID = &id
				}
				var catID *uuid.UUID
				if item.DocumentCategoryID.Valid {
					id := uuid.UUID(item.DocumentCategoryID.Bytes)
					catID = &id
				}
				mItem := ManifestItem{
					ID:             item.ID,
					DocumentID:     item.DocumentID,
					Title:          item.DocumentTitle,
					TypeID:         typeID,
					TypeName:       item.DocumentType.String,
					CategoryID:     catID,
					Date:           item.DocumentDate.Time.Format("02/01/2006"),
					Status:         item.Status,
					PhysicalStatus: item.CurrentPhysicalStatus.String,
					PreviewPath:    item.PreviewPath.String,
					Metadata:       map[string]interface{}{},
					ExtractedText:  item.ExtractedText.String,
					Notes:          item.Notes.String,
				}
				if len(item.DocumentMetadata) > 0 {
					json.Unmarshal(item.DocumentMetadata, &mItem.Metadata)
				}
				manifestItems = append(manifestItems, mItem)
			}

			receivedAt := "-"
			if manifest.ReceivedAt.Valid {
				receivedAt = manifest.ReceivedAt.Time.Format("02/01/2006 15:04")
			}

			return &ManifestDetail{
				ID:             manifest.ID,
				ManifestNo:     manifest.ManifestNo,
				Title:          fmt.Sprintf("Manifest %s", manifest.ManifestNo),
				TypeName:       "Manifest Serah Terima",
				OwnerName:      manifest.SenderName,
				DepartmentName: manifest.DepartmentName,
				Status:         manifest.Status,
				TotalItems:     manifest.TotalItems,
				Priority:       "Tinggi",
				Estimate:       fmt.Sprintf("%d Dokumen", manifest.TotalItems),
				CreatedAt:      manifest.CreatedAt.Time.Format("02/01/2006 15:04"),
				ReceivedAt:     receivedAt,
				Source:         "INB",
				Items:          manifestItems,
				Timeline: []TimelineItem{
					{Status: "Manifest Diajukan", Timestamp: manifest.CreatedAt.Time.Format("02/01/2006 15:04"), IsActive: true},
					{Status: "Diterima di Staging Area", Timestamp: receivedAt, IsActive: true},
					{Status: "Sedang Diproses Digitalisasi", Timestamp: time.Now().Format("02/01/2006 15:04"), IsActive: true},
				},
			}, nil
		}
	}

	// 2. Try finding by Manifest No
	manifest, err := s.repo.GetManifestByNo(ctx, idStr)
	if err == nil {
		items, _ := s.repo.GetManifestItems(ctx, manifest.ID)
		manifestItems := make([]ManifestItem, 0)
		for _, item := range items {
			var typeID *uuid.UUID
			if item.DocumentTypeID.Valid {
				id := uuid.UUID(item.DocumentTypeID.Bytes)
				typeID = &id
			}
			var catID *uuid.UUID
			if item.DocumentCategoryID.Valid {
				id := uuid.UUID(item.DocumentCategoryID.Bytes)
				catID = &id
			}
			mItem := ManifestItem{
				ID:             item.ID,
				DocumentID:     item.DocumentID,
				Title:          item.DocumentTitle,
				TypeID:         typeID,
				TypeName:       item.DocumentType.String,
				CategoryID:     catID,
				Date:           item.DocumentDate.Time.Format("02/01/2006"),
				Status:         item.Status,
				PhysicalStatus: item.CurrentPhysicalStatus.String,
				PreviewPath:    item.PreviewPath.String,
				Metadata:       map[string]interface{}{},
				ExtractedText:  item.ExtractedText.String,
				Notes:          item.Notes.String,
			}
			if len(item.DocumentMetadata) > 0 {
				json.Unmarshal(item.DocumentMetadata, &mItem.Metadata)
			}
			manifestItems = append(manifestItems, mItem)
		}

		receivedAt := "-"
		if manifest.ReceivedAt.Valid {
			receivedAt = manifest.ReceivedAt.Time.Format("02/01/2006 15:04")
		}

		return &ManifestDetail{
			ID:             manifest.ID,
			ManifestNo:     manifest.ManifestNo,
			Title:          fmt.Sprintf("Manifest %s", manifest.ManifestNo),
			TypeName:       "Manifest Serah Terima",
			OwnerName:      manifest.SenderName,
			DepartmentName: manifest.DepartmentName,
			Status:         manifest.Status,
			TotalItems:     manifest.TotalItems,
			Priority:       "Tinggi",
			Estimate:       fmt.Sprintf("%d Dokumen", manifest.TotalItems),
			CreatedAt:      manifest.CreatedAt.Time.Format("02/01/2006 15:04"),
			ReceivedAt:     receivedAt,
			Source:         "INB",
			Items:          manifestItems,
			Timeline: []TimelineItem{
				{Status: "Manifest Diajukan", Timestamp: manifest.CreatedAt.Time.Format("02/01/2006 15:04"), IsActive: true},
				{Status: "Diterima di Staging Area", Timestamp: receivedAt, IsActive: true},
				{Status: "Sedang Diproses Digitalisasi", Timestamp: time.Now().Format("02/01/2006 15:04"), IsActive: true},
			},
		}, nil
	}

	// 3. Fallback to Search by prefix (Legacy/Direct Document Scan)
	doc, err := s.repo.GetDocumentByShortID(ctx, pgtype.Text{String: strings.ToLower(idStr), Valid: true})
	if err != nil {
		return nil, fmt.Errorf("manifest or document not found with ID: %s", idStr)
	}

	return &ManifestDetail{
		ID:             uuid.Nil, // Not a saved manifest
		ManifestNo:     idStr,
		Title:          doc.Title,
		TypeName:       doc.TypeName.String,
		OwnerName:      doc.OwnerName,
		DepartmentName: doc.DepartmentName.String,
		Status:         doc.Status,
		TotalItems:     1,
		CreatedAt:      doc.CreatedAt.Time.Format("02/01/2006 15:04"),
		Items: []ManifestItem{
			{
				ID:             uuid.Nil,
				DocumentID:     doc.ID,
				Title:          doc.Title,
				TypeName:       doc.TypeName.String,
				Date:           doc.CreatedAt.Time.Format("02/01/2006"),
				Status:         doc.Status,
				PhysicalStatus: doc.PhysicalStatus.String,
			},
		},
	}, nil
}

type ReceiveRequestItem struct {
	DocumentID     uuid.UUID `json:"document_id"`
	PhysicalStatus string    `json:"physical_status"`
	Notes          string    `json:"notes"`
}

type IndexingRequest struct {
	ManifestID uuid.UUID              `json:"manifest_id"`
	DocumentID uuid.UUID              `json:"document_id"`
	CategoryID uuid.UUID              `json:"category_id"`
	TypeID     uuid.UUID              `json:"type_id"`
	Metadata   map[string]interface{} `json:"metadata"`
}

func (s *IntakeService) IndexDocument(ctx context.Context, req IndexingRequest) error {
	metaJson, _ := json.Marshal(req.Metadata)
	
	// Determine Title from metadata if possible
	title := "Indexed Document"
	if val, ok := req.Metadata["document_number"].(string); ok && val != "" {
		title = fmt.Sprintf("Doc: %s", val)
	}

	err := s.repo.UpdateDocumentIndexing(ctx, repository.UpdateDocumentIndexingParams{
		ID:             req.DocumentID,
		Title:          title,
		TypeID:         pgtype.UUID{Bytes: req.TypeID, Valid: req.TypeID != uuid.Nil},
		Metadata:       metaJson,
		PhysicalStatus: pgtype.Text{String: "digitized", Valid: true},
	})
	if err != nil {
		return err
	}

	err = s.repo.UpdateManifestItemStatusByDocID(ctx, repository.UpdateManifestItemStatusByDocIDParams{
		ManifestID: req.ManifestID,
		DocumentID: req.DocumentID,
		Status:     "digitized",
	})
	if err != nil {
		return err
	}

	progress, err := s.repo.GetManifestProgress(ctx, req.ManifestID)
	if err == nil && progress.TotalItems == progress.DigitizedCount {
		_ = s.repo.UpdateManifestStatus(ctx, repository.UpdateManifestStatusParams{
			ID:     req.ManifestID,
			Status: "completed",
		})
	}

	return nil
}

func (s *IntakeService) ReceiveDocument(ctx context.Context, manifestID uuid.UUID, receivedBy uuid.UUID, mainNotes string, items []ReceiveRequestItem) error {
	// 1. If ManifestID is provided, update the manifest and its items
	if manifestID != uuid.Nil {
		err := s.repo.UpdateManifestStatus(ctx, repository.UpdateManifestStatusParams{
			ID:         manifestID,
			Status:     "received",
			ReceivedBy: pgtype.UUID{Bytes: receivedBy, Valid: true},
		})
		if err != nil {
			return err
		}

		// Update each item
		for _, item := range items {
			// Update document global physical status
			err = s.repo.UpdateDocumentPhysicalStatus(ctx, repository.UpdateDocumentPhysicalStatusParams{
				ID:                item.DocumentID,
				PhysicalStatus:    pgtype.Text{String: item.PhysicalStatus, Valid: true},
				CurrentManifestID: pgtype.UUID{Bytes: manifestID, Valid: true},
			})
			if err != nil {
				log.Printf("Failed to update doc %s physical status: %v", item.DocumentID, err)
			}

			// ACTIVATE DOCUMENT: Once physically received, it becomes 'active'
			err = s.repo.UpdateDocumentStatus(ctx, repository.UpdateDocumentStatusParams{
				ID:     item.DocumentID,
				Status: "active",
			})
			if err != nil {
				log.Printf("Failed to activate doc %s: %v", item.DocumentID, err)
			}
		}

		// SEND NOTIFICATIONS
		manifest, _ := s.repo.GetManifest(ctx, manifestID)
		if manifest.ID != uuid.Nil {
			count := len(items)
			notifMsg := fmt.Sprintf("Berkas fisik manifest %s (%d dokumen) telah diterima oleh Document Controller.", manifest.ManifestNo, count)
			
			// 1. To Sender/Owner
			s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
				UserID: manifest.SenderID,
				Title:  "Berkas Fisik Diterima",
				Body:   pgtype.Text{String: notifMsg, Valid: true},
				Type:   "intake",
			})

			// 2. To Department Manager (Head)
			dept, err := s.repo.GetDepartment(ctx, manifest.DepartmentID)
			if err == nil && dept.HeadID.Valid {
				s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
					UserID: dept.HeadID.Bytes,
					Title:  "Notifikasi Penerimaan Berkas (Departemen)",
					Body:   pgtype.Text{String: fmt.Sprintf("Dokumen dari %s telah diterima oleh Central Document.", manifest.SenderName), Valid: true},
					Type:   "intake",
				})
			}
		}

		return nil
	}

	// 2. Legacy/Single Document Mode (If no ManifestID provided, fallback to creating one)
	if len(items) == 0 {
		return fmt.Errorf("no items to receive")
	}

	firstItem := items[0]
	doc, err := s.repo.GetDocument(ctx, firstItem.DocumentID)
	if err != nil {
		return err
	}

	manifest, err := s.repo.CreatePhysicalManifest(ctx, repository.CreatePhysicalManifestParams{
		ManifestNo:   fmt.Sprintf("INB-%s", doc.ID.String()[:8]),
		SenderID:     doc.OwnerID,
		DepartmentID: doc.DepartmentID,
		TotalItems:   int32(len(items)),
		Status:       "received",
		Notes:        pgtype.Text{String: mainNotes, Valid: mainNotes != ""},
	})
	if err != nil {
		return err
	}

	// Finalize manifest
	s.repo.UpdateManifestStatus(ctx, repository.UpdateManifestStatusParams{
		ID:         manifest.ID,
		Status:     "received",
		ReceivedBy: pgtype.UUID{Bytes: receivedBy, Valid: true},
	})

	for _, item := range items {
		s.repo.AddManifestItem(ctx, repository.AddManifestItemParams{
			ManifestID: manifest.ID,
			DocumentID: item.DocumentID,
			Status:     item.PhysicalStatus,
			Notes:      pgtype.Text{String: item.Notes, Valid: item.Notes != ""},
		})

		s.repo.UpdateDocumentPhysicalStatus(ctx, repository.UpdateDocumentPhysicalStatusParams{
			ID:                item.DocumentID,
			PhysicalStatus:    pgtype.Text{String: item.PhysicalStatus, Valid: true},
			CurrentManifestID: pgtype.UUID{Bytes: manifest.ID, Valid: true},
		})

		// ACTIVATE DOCUMENT
		s.repo.UpdateDocumentStatus(ctx, repository.UpdateDocumentStatusParams{
			ID:     item.DocumentID,
			Status: "active",
		})
	}

	return nil
}

func (s *IntakeService) RejectManifest(ctx context.Context, manifestID uuid.UUID, rejectedBy uuid.UUID, reason string) error {
	// Update manifest status to rejected
	err := s.repo.UpdateManifestStatus(ctx, repository.UpdateManifestStatusParams{
		ID:         manifestID,
		Status:     "rejected",
		ReceivedBy: pgtype.UUID{Bytes: rejectedBy, Valid: true},
	})
	if err != nil {
		return err
	}

	// Optionally update items/docs status to reflect rejection
	items, _ := s.repo.GetManifestItems(ctx, manifestID)
	for _, item := range items {
		_ = s.repo.UpdateDocumentPhysicalStatus(ctx, repository.UpdateDocumentPhysicalStatusParams{
			ID:                item.DocumentID,
			PhysicalStatus:    pgtype.Text{String: "rejected", Valid: true},
			CurrentManifestID: pgtype.UUID{Bytes: manifestID, Valid: true},
		})
	}

	return nil
}

func (s *IntakeService) ListPendingManifests(ctx context.Context) ([]ManifestDetail, error) {
	// 1. Get real manifests
	manifests, err := s.repo.ListPendingManifests(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]ManifestDetail, 0)
	for _, m := range manifests {
		result = append(result, ManifestDetail{
			ID:             m.ID,
			ManifestNo:     m.ManifestNo,
			Title:          fmt.Sprintf("Manifest %s", m.ManifestNo),
			TypeName:       "Manifest Serah Terima",
			OwnerName:      m.SenderName,
			DepartmentName: m.DepartmentName,
			Status:         m.Status,
			TotalItems:     m.TotalItems,
			CreatedAt:      m.CreatedAt.Time.Format("02/01/2006 15:04"),
		})
	}

	// 2. Get individual documents that are pending intake but not in a manifest
	docs, err := s.repo.ListPendingDocumentsWithoutManifest(ctx)
	if err == nil {
		for _, d := range docs {
			result = append(result, ManifestDetail{
				ID:             uuid.Nil, // Special marker for direct document
				ManifestNo:     d.ID.String()[:8],
				Title:          d.Title,
				TypeName:       d.TypeName.String,
				OwnerName:      d.OwnerName,
				DepartmentName: d.DepartmentName,
				Status:         "pending",
				CreatedAt:      d.CreatedAt.Time.Format("02/01/2006 15:04"),
			})
		}
	}

	return result, nil
}

func (s *IntakeService) GetStats(ctx context.Context) (map[string]int64, error) {
	stats, err := s.repo.GetIntakeStats(ctx)
	if err != nil {
		return nil, err
	}
	return map[string]int64{
		"received": stats.ReceivedToday,
		"pending":  stats.PendingCount,
		"rejected": stats.RejectedToday,
	}, nil
}

func (s *IntakeService) ListStagingManifests(ctx context.Context) ([]ManifestDetail, error) {
	manifests, err := s.repo.ListStagingManifests(ctx)
	if err != nil {
		return nil, err
	}

	result := make([]ManifestDetail, 0)
	for _, m := range manifests {
		source := "INB"
		if strings.HasPrefix(m.ManifestNo, "MIG") {
			source = "MIG"
		}

		receivedAt := ""
		if m.ReceivedAt.Valid {
			receivedAt = m.ReceivedAt.Time.Format("02/01/2006 15:04")
		}

		result = append(result, ManifestDetail{
			ID:             m.ID,
			ManifestNo:     m.ManifestNo,
			Title:          fmt.Sprintf("Manifest %s", m.ManifestNo),
			TypeName:       "Manifest Serah Terima",
			OwnerName:      m.SenderName,
			DepartmentName: m.DepartmentName,
			Status:         m.Status,
			TotalItems:     m.TotalItems,
			CreatedAt:      m.CreatedAt.Time.Format("02/01/2006 15:04"),
			ReceivedAt:     receivedAt,
			Source:         source,
		})
	}

	return result, nil
}

func (s *IntakeService) GetStagingStats(ctx context.Context) (map[string]interface{}, error) {
	stats, err := s.repo.GetStagingStats(ctx)
	if err != nil {
		return nil, err
	}

	dcs, _ := s.repo.GetActiveDocControllers(ctx)
	dcList := []map[string]string{}
	for _, dc := range dcs {
		avatar := ""
		if dc.AvatarUrl.Valid {
			avatar = dc.AvatarUrl.String
		}
		dcList = append(dcList, map[string]string{
			"name":   dc.FullName,
			"avatar": avatar,
		})
	}

	return map[string]interface{}{
		"active_count":    stats.TotalStaging,
		"overdue_count":   stats.OverdueCount,
		"completed_today": stats.CompletedToday,
		"avg_wait_time":   "1h 24m",
		"active_dc_count": stats.ActiveDcCount,
		"active_dc_list":  dcList,
	}, nil
}

func (s *IntakeService) ListLabelingDocuments(ctx context.Context, status string) ([]LabelingDocument, error) {
	items, err := s.repo.ListDocumentsForLabeling(ctx, status)
	if err != nil {
		return nil, err
	}

	result := make([]LabelingDocument, 0, len(items))
	for _, item := range items {
		result = append(result, LabelingDocument{
			DocumentID:     item.DocumentID,
			Title:          item.Title,
			DepartmentName: item.DepartmentName.String,
			ManifestNo:     item.ManifestNo,
			Status:         item.ManifestItemStatus,
		})
	}
	return result, nil
}

func (s *IntakeService) MarkAsLabeled(ctx context.Context, docIDs []uuid.UUID) error {
	// 1. Update manifest items status to 'labeled'
	err := s.repo.UpdateManifestItemStatusBulk(ctx, repository.UpdateManifestItemStatusBulkParams{
		Column1: docIDs,
		Status:  "labeled",
	})
	if err != nil {
		return err
	}

	// 2. Update documents physical status
	return s.repo.UpdateDocumentPhysicalStatusBulk(ctx, repository.UpdateDocumentPhysicalStatusBulkParams{
		Column1:        docIDs,
		PhysicalStatus: pgtype.Text{String: "received", Valid: true},
	})
}

func (s *IntakeService) GetLabelingStats(ctx context.Context) (*LabelingStats, error) {
	stats, err := s.repo.GetLabelingStats(ctx)
	if err != nil {
		return nil, err
	}
	return &LabelingStats{
		WaitingCount: stats.WaitingCount,
		PrintedCount: stats.PrintedCount,
		StoredCount:  stats.StoredCount,
	}, nil
}

func (s *IntakeService) SearchBoxes(ctx context.Context, query string) ([]BoxInfo, error) {
	items, err := s.repo.SearchBoxes(ctx, pgtype.Text{String: query, Valid: true})
	if err != nil {
		return nil, err
	}

	result := make([]BoxInfo, 0, len(items))
	for _, item := range items {
		result = append(result, BoxInfo{
			ID:             item.ID,
			Name:           item.Name,
			RackName:       item.RackName,
			DepartmentName: item.DepartmentName,
			CompanyName:    item.CompanyName,
			BranchName:     item.BranchName,
			LocationDetail: item.LocationDetail.String,
		})
	}
	return result, nil
}

func (s *IntakeService) MarkAsArchivedBulk(ctx context.Context, docIDs []uuid.UUID, boxID *uuid.UUID) error {
	// 1. Update manifest items status to 'archived'
	err := s.repo.UpdateManifestItemStatusBulk(ctx, repository.UpdateManifestItemStatusBulkParams{
		Column1: docIDs,
		Status:  "archived",
	})
	if err != nil {
		return err
	}

	// 2. Update documents physical status and box_id
	boxIDStr := ""
	if boxID != nil {
		boxIDStr = boxID.String()
	}

	return s.repo.UpdateDocumentPhysicalStatusBulk(ctx, repository.UpdateDocumentPhysicalStatusBulkParams{
		Column1:        docIDs,
		PhysicalStatus: pgtype.Text{String: "archived", Valid: true},
		BoxID:          boxIDStr,
	})
}

func (s *IntakeService) GetSmartRecommendation(ctx context.Context, docIDs []uuid.UUID) (*BoxInfo, error) {
	if len(docIDs) == 0 {
		return nil, errors.New("no documents provided")
	}

	// 1. Get first document metadata to determine zonation requirements
	doc, err := s.repo.GetDocument(ctx, docIDs[0])
	if err != nil {
		return nil, err
	}

	// 2. Find eligible boxes based on department, category zonation, and capacity
	var categoryID uuid.UUID
	if doc.CategoryID.Valid {
		categoryID = doc.CategoryID.Bytes
	}

	boxes, err := s.repo.FindEligibleBoxes(ctx, repository.FindEligibleBoxesParams{
		DepartmentID: doc.DepartmentID,
		CategoryID:   categoryID,
	})
	if err != nil || len(boxes) == 0 {
		return nil, nil // No recommendation found
	}

	// 3. Return the best one (highest occupancy first to fill partially empty boxes)
	best := boxes[0]
	return &BoxInfo{
		ID:             best.ID,
		Name:           best.Name,
		RackName:       best.RackName,
		DepartmentName: best.DepartmentName,
		LocationDetail: best.LocationDetail.String,
	}, nil
}



