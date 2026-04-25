package worker

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	"github.com/hibiken/asynq"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
)

type TaskProcessor struct {
	repo    repository.Querier
	storage *infra.StorageService
	ai      *infra.AIService
	search  *infra.SearchService
}

func NewTaskProcessor(repo repository.Querier, storage *infra.StorageService, ai *infra.AIService, search *infra.SearchService) *TaskProcessor {
	return &TaskProcessor{repo: repo, storage: storage, ai: ai, search: search}
}

func (p *TaskProcessor) ProcessDocumentOCR(ctx context.Context, t *asynq.Task) error {
	var payload DocumentOCRPayload
	if err := json.Unmarshal(t.Payload(), &payload); err != nil {
		return fmt.Errorf("json unmarshal failed: %v", err)
	}

	log.Printf("Processing OCR for Document: %s", payload.DocumentID)

	// 1. Get document info
	doc, err := p.repo.GetDocument(ctx, payload.DocumentID)
	if err != nil {
		log.Printf("[TaskProcessor] Error getting document %s: %v", payload.DocumentID, err)
		return err
	}

	// 2. Download from storage
	reader, err := p.storage.Download(ctx, doc.FilePath)
	if err != nil {
		log.Printf("[TaskProcessor] Error downloading file %s: %v", doc.FilePath, err)
		return err
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("[TaskProcessor] Error reading file content %s: %v", doc.FileName, err)
		return err
	}

	// 3. Process OCR
	rawText, err := p.ai.ProcessOCR(ctx, doc.FileName, content)
	if err != nil {
		log.Printf("[TaskProcessor] Error processing OCR for doc %s: %v", doc.ID, err)
		return err
	}

	// 3.1 AI Refinement (Optional based on config)
	text, err := p.ai.RefineOCRText(ctx, rawText)
	if err != nil {
		log.Printf("[TaskProcessor] Warning: AI refinement failed for doc %s: %v", doc.ID, err)
		text = rawText // Fallback
	}

	// 3.2 AI Metadata Extraction (Optional based on config)
	metadata, err := p.ai.ExtractMetadata(ctx, text)
	if err != nil {
		log.Printf("[TaskProcessor] Warning: AI metadata extraction failed for doc %s: %v", doc.ID, err)
		metadata = map[string]interface{}{} // Fallback
	}

	metadataJSON, _ := json.Marshal(metadata)

	// 4. Update DB
	err = p.repo.UpdateDocumentOCR(ctx, repository.UpdateDocumentOCRParams{
		ID:            doc.ID,
		ExtractedText: pgtype.Text{String: text, Valid: text != ""},
		Metadata:      metadataJSON,
	})
	if err != nil {
		log.Printf("[TaskProcessor] Error updating DB with OCR text for doc %s: %v", doc.ID, err)
		return err
	}

	// 5. Index to Elasticsearch
	err = p.search.IndexDocument(ctx, infra.DocumentIndex{
		ID:           doc.ID,
		Title:        doc.Title,
		Content:      text,
		DepartmentID: doc.DepartmentID,
		Tags:         doc.Tags,
		CreatedAt:    doc.CreatedAt.Time.String(),
	})
	if err != nil {
		log.Printf("[TaskProcessor] Warning: Failed to index doc %s to ES: %v", doc.ID, err)
	}

	// 6. Update Batch Progress
	if doc.BatchID.Valid {
		err = p.repo.UpdateBatchProgress(ctx, doc.BatchID.Bytes)
		if err != nil {
			log.Printf("[TaskProcessor] Warning: Failed to update batch progress for %s: %v", doc.BatchID.Bytes, err)
		}
	}

	log.Printf("[TaskProcessor] Completed OCR processing for Document: %s", doc.ID)
	return nil
}
