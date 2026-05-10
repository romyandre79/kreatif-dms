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
	"time"
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
	log.Printf("!!! WORKER RECEIVED TASK: %s", t.Type())
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

	// Send "Processing" notification
	_, _ = p.repo.CreateNotification(ctx, repository.CreateNotificationParams{
		UserID:     doc.OwnerID,
		Title:      "Dokumen Sedang Diproses",
		Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' sedang dalam proses analisis OCR.", doc.Title), Valid: true},
		Type:       "INFO",
		EntityType: pgtype.Text{String: "document", Valid: true},
		EntityID:   pgtype.UUID{Bytes: doc.ID, Valid: true},
		Channel:    pgtype.Text{String: "app", Valid: true},
	})

	// 2. Download from storage
	reader, err := p.storage.Download(ctx, doc.FilePath.String)
	if err != nil {
		log.Printf("[TaskProcessor] Error downloading file %s: %v", doc.FilePath.String, err)
		return err
	}
	defer reader.Close()

	content, err := io.ReadAll(reader)
	if err != nil {
		log.Printf("[TaskProcessor] Error reading file content %s: %v", doc.FileName.String, err)
		return err
	}

	// 3. Process OCR
	ocrStart := time.Now()
	ocrRes, err := p.ai.ProcessOCR(ctx, doc.FileName.String, content)
	if err != nil {
		log.Printf("[TaskProcessor] Error processing OCR for doc %s: %v", doc.ID, err)
		return err
	}
	ocrDuration := time.Since(ocrStart)
	rawText := ocrRes.FullText

	// 3.1 Save to ocr_jobs table (Detailed results)
	wordsJSON, _ := json.Marshal(ocrRes.Words)
	
	// Calculate avg confidence
	var totalConf float64
	if len(ocrRes.Words) > 0 {
		for _, w := range ocrRes.Words {
			totalConf += w.Confidence
		}
		totalConf = totalConf / float64(len(ocrRes.Words)) * 100
	}

	var confNumeric pgtype.Numeric
	confNumeric.Scan(fmt.Sprintf("%.2f", totalConf))

	// Save preview paths
	if ocrRes.PreviewPath == "" && len(ocrRes.PreviewPaths) > 0 {
		ocrRes.PreviewPath = ocrRes.PreviewPaths[0]
	}

	previewPathsJSON, _ := json.Marshal(ocrRes.PreviewPaths)
	if len(ocrRes.PreviewPaths) == 0 && ocrRes.PreviewPath != "" {
		previewPathsJSON, _ = json.Marshal([]string{ocrRes.PreviewPath})
	}

	_, err = p.repo.CreateOCRJob(ctx, repository.CreateOCRJobParams{
		EntityType:       "document",
		EntityID:         doc.ID,
		SourceFilePath:   doc.FilePath,
		RawText:          pgtype.Text{String: rawText, Valid: true},
		WordCount:        pgtype.Int4{Int32: int32(len(ocrRes.Words)), Valid: true},
		ConfidenceAvg:    confNumeric,
		WordsJson:        wordsJSON,
		Status:           "Success",
		ProcessingTimeMs: pgtype.Int4{Int32: int32(ocrDuration.Milliseconds()), Valid: true},
		PreviewPath:      pgtype.Text{String: ocrRes.PreviewPath, Valid: ocrRes.PreviewPath != ""},
		PreviewPaths:     previewPathsJSON,
	})
	if err != nil {
		log.Printf("[TaskProcessor] Warning: Failed to save to ocr_jobs: %v", err)
	}

	// 3.1 Use AI Insights from OCR Service if available
	metadata := map[string]interface{}{}
	text := rawText

	if ocrRes.Insight != nil {
		log.Printf("[TaskProcessor] Using AI Insights for Document: %s (Type: %s)", doc.ID, ocrRes.Insight.DocType)
		metadata = ocrRes.Insight.Entities
		// Include basic info in metadata
		metadata["ai_doc_type"] = ocrRes.Insight.DocType
		metadata["ai_summary"] = ocrRes.Insight.Summary
		
		if ocrRes.Insight.CleanedText != "" {
			text = ocrRes.Insight.CleanedText
		}
	} else {
		// Fallback to legacy refinement if ocr-service didn't provide insight
		log.Printf("[TaskProcessor] Warning: No AI Insights from OCR Service, falling back to legacy refinement for doc %s", doc.ID)
		refined, err := p.ai.RefineOCRText(ctx, rawText)
		if err == nil {
			text = refined
		}
		
		extractedMeta, err := p.ai.ExtractMetadata(ctx, text)
		if err == nil {
			metadata = extractedMeta
		}
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
		Metadata:     metadata,
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

	// Send "Completed" notification
	_, _ = p.repo.CreateNotification(ctx, repository.CreateNotificationParams{
		UserID:     doc.OwnerID,
		Title:      "Dokumen Selesai Diproses",
		Body:       pgtype.Text{String: fmt.Sprintf("Dokumen '%s' telah berhasil dianalisis.", doc.Title), Valid: true},
		Type:       "SUCCESS",
		EntityType: pgtype.Text{String: "document", Valid: true},
		EntityID:   pgtype.UUID{Bytes: doc.ID, Valid: true},
		Channel:    pgtype.Text{String: "app", Valid: true},
	})

	return nil
}
