package worker

import (
	"encoding/json"

	"github.com/google/uuid"
	"github.com/hibiken/asynq"
)

const (
	TypeDocumentOCR = "doc:ocr"
	TypeDocumentIndex = "doc:index"
)

type DocumentOCRPayload struct {
	DocumentID uuid.UUID `json:"document_id"`
}

func NewDocumentOCRTask(docID uuid.UUID) (*asynq.Task, error) {
	payload, err := json.Marshal(DocumentOCRPayload{DocumentID: docID})
	if err != nil {
		return nil, err
	}
	return asynq.NewTask(TypeDocumentOCR, payload), nil
}
