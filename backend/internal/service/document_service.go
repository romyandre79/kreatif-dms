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
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/worker"
	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
)

type DocumentService struct {
	repo    repository.Querier
	storage *infra.StorageService
	asynq   *asynq.Client
}

func NewDocumentService(repo repository.Querier, storage *infra.StorageService, asynqClient *asynq.Client) *DocumentService {
	return &DocumentService{repo: repo, storage: storage, asynq: asynqClient}
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
}

func (s *DocumentService) UploadDocument(ctx context.Context, p UploadDocumentParams) (repository.Document, error) {
	log.Printf("[DocumentService] Uploading document: %s (Size: %d, Owner: %s)", p.FileName, p.FileSize, p.OwnerID)

	// 1. Generate unique file path
	objectName := fmt.Sprintf("%s/%s", p.DepartmentID, p.FileName)

	// 2. Upload to MinIO (Encryption disabled for local dev/KMS not configured)
	_, err := s.storage.Upload(ctx, objectName, p.Content, p.FileSize, p.MimeType, false)
	if err != nil {
		log.Printf("[DocumentService] Error uploading to storage: %v", err)
		return repository.Document{}, err
	}

	log.Printf("[DocumentService] Storage upload successful, creating DB record for %s", p.FileName)

	// 3. Save to DB
	doc, err := s.repo.CreateDocument(ctx, repository.CreateDocumentParams{
		Title:        p.Title,
		FileName:     p.FileName,
		FilePath:     objectName,
		FileSize:     p.FileSize,
		MimeType:     p.MimeType,
		OwnerID:      p.OwnerID,
		CompanyID:    p.CompanyID,
		BranchID:     p.BranchID,
		DepartmentID: p.DepartmentID,
		Status:       "processing",
		BatchID:      pgtype.UUID{Bytes: p.BatchID, Valid: p.BatchID != uuid.Nil},
	})
	if err != nil {
		log.Printf("[DocumentService] Error creating document record in DB: %v", err)
		return repository.Document{}, err
	}

	log.Printf("[DocumentService] DB record created successfully: %s", doc.ID)

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

	return doc, nil
}


func (s *DocumentService) GetWatermarkedPDF(ctx context.Context, docID uuid.UUID, userFullName string, position string) ([]byte, error) {
	log.Printf("[DocumentService] Accessing watermarked PDF: %s (Requested by: %s)", docID, userFullName)

	// 1. Get Doc from DB
	doc, err := s.repo.GetDocument(ctx, docID)
	if err != nil {
		log.Printf("[DocumentService] Error getting document from DB: %v", err)
		return nil, err
	}

	// 2. Download from MinIO
	reader, err := s.storage.Download(ctx, doc.FilePath)
	if err != nil {
		log.Printf("[DocumentService] Error downloading from storage: %v", err)
		return nil, err
	}
	defer reader.Close()

	// 3. Read content
	content, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("[DocumentService] Error reading content: %v", err)
		return nil, err
	}

	// 4. Apply Watermark using pdfcpu
	watermarkText := fmt.Sprintf("CONFIDENTIAL - %s - %s", userFullName, time.Now().Format("2006-01-02"))
	
	// Default configuration for watermark
	wm, err := api.TextWatermark(watermarkText, "font:Helvetica, points:24, scale:0.5, op:0.3, rot:45", true, false, types.POINTS)
	if err != nil {
		log.Printf("[DocumentService] Error creating watermark: %v", err)
		return nil, err
	}

	// Output buffer
	var out bytes.Buffer
	err = api.AddWatermarks(bytes.NewReader(content), &out, nil, wm, nil)
	if err != nil {
		log.Printf("[DocumentService] Error applying watermark: %v", err)
		return nil, err
	}

	log.Printf("[DocumentService] Watermarked PDF generated: %s", docID)
	return out.Bytes(), nil
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

func (s *DocumentService) ListRecentDocuments(ctx context.Context, limit int) ([]repository.ListRecentDocumentsRow, error) {
	return s.repo.ListRecentDocuments(ctx, repository.ListRecentDocumentsParams{
		Limit:  int32(limit),
		Offset: 0,
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
