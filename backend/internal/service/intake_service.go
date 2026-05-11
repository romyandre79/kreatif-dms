package service

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
)

type IntakeService struct {
	repo repository.Querier
}

func NewIntakeService(repo repository.Querier) *IntakeService {
	return &IntakeService{repo: repo}
}

type ManifestDetail struct {
	DocumentID     uuid.UUID `json:"document_id"`
	Title          string    `json:"title"`
	TypeName       string    `json:"type_name"`
	OwnerName      string    `json:"owner_name"`
	DepartmentName string    `json:"department_name"`
	Status         string    `json:"status"`
}

func (s *IntakeService) GetDocumentByManifestID(ctx context.Context, shortID string) (*ManifestDetail, error) {
	// Search by prefix (first 8 chars)
	doc, err := s.repo.GetDocumentByShortID(ctx, pgtype.Text{String: strings.ToLower(shortID), Valid: true})
	if err != nil {
		return nil, fmt.Errorf("document not found with ID prefix: %s", shortID)
	}

	return &ManifestDetail{
		DocumentID:     doc.ID,
		Title:          doc.Title,
		TypeName:       doc.TypeName.String,
		OwnerName:      doc.OwnerName,
		DepartmentName: doc.DepartmentName.String,
		Status:         doc.Status,
	}, nil
}

func (s *IntakeService) ReceiveDocument(ctx context.Context, docID uuid.UUID, receivedBy uuid.UUID, notes string) error {
	// 1. Get doc to get sender info
	doc, err := s.repo.GetDocument(ctx, docID)
	if err != nil {
		return err
	}

	manifest, err := s.repo.CreatePhysicalManifest(ctx, repository.CreatePhysicalManifestParams{
		ManifestNo:   fmt.Sprintf("INB-%s", doc.ID.String()[:8]),
		SenderID:     doc.OwnerID,
		DepartmentID: doc.DepartmentID,
		TotalItems:   1,
		Status:       "received",
		Notes:        pgtype.Text{String: notes, Valid: notes != ""},
	})
	if err != nil {
		return err
	}

	// 3. Update Manifest status (finalize)
	err = s.repo.UpdateManifestStatus(ctx, repository.UpdateManifestStatusParams{
		ID:         manifest.ID,
		Status:     "received",
		ReceivedBy: pgtype.UUID{Bytes: receivedBy, Valid: true},
	})
	if err != nil {
		return err
	}

	// 4. Link Document to Manifest
	_, err = s.repo.AddManifestItem(ctx, repository.AddManifestItemParams{
		ManifestID: manifest.ID,
		DocumentID: docID,
		Status:     "received",
		Notes:      pgtype.Text{String: notes, Valid: notes != ""},
	})
	if err != nil {
		return err
	}

	// 5. Update document status
	log.Printf("[IntakeService] Document %s received by %s", docID, receivedBy)
	return nil
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
