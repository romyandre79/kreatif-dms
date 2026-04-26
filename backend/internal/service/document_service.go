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

	// 2. Upload to MinIO with Encryption
	_, err := s.storage.Upload(ctx, objectName, p.Content, p.FileSize, p.MimeType, true)
	if err != nil {
		log.Printf("[DocumentService] Error uploading to storage: %v", err)
		return repository.Document{}, err
	}

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
		log.Printf("[DocumentService] Error creating document record: %v", err)
		return repository.Document{}, err
	}

	// 4. Enqueue OCR Task
	if s.asynq != nil {
		task, err := worker.NewDocumentOCRTask(doc.ID)
		if err == nil {
			_, err = s.asynq.Enqueue(task)
			if err != nil {
				log.Printf("[DocumentService] Error enqueuing OCR task: %v", err)
			} else {
				log.Printf("[DocumentService] OCR task enqueued for doc: %s", doc.ID)
			}
		}
	} else {
		log.Printf("[DocumentService] WARNING: Skipping OCR task for doc %s because Redis/Asynq is unavailable", doc.ID)
	}

	log.Printf("[DocumentService] Document uploaded successfully: %s", doc.ID)
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
	wm, err := api.TextWatermark(watermarkText, "font:Roboto, points:24, scale:0.5, op:0.3, rot:45", true, false, types.POINTS)
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
