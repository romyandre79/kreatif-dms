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
	repo     repository.Querier
	notifSvc *NotificationService
}

func NewIntakeService(repo repository.Querier, notifSvc *NotificationService) *IntakeService {
	return &IntakeService{repo: repo, notifSvc: notifSvc}
}

type ManifestItem struct {
	ID             uuid.UUID `json:"id"`
	DocumentID     uuid.UUID `json:"document_id"`
	Title          string    `json:"title"`
	TypeName       string    `json:"type_name"`
	Date           string    `json:"date"`
	Status         string    `json:"status"`
	PhysicalStatus string    `json:"physical_status"`
	Notes          string    `json:"notes"`
}

type ManifestDetail struct {
	ID             uuid.UUID      `json:"id"`
	ManifestNo     string         `json:"manifest_no"`
	Title          string         `json:"title"`
	TypeName       string         `json:"type_name"`
	OwnerName      string         `json:"owner_name"`
	DepartmentName string         `json:"department_name"`
	Status         string         `json:"status"`
	CreatedAt      string         `json:"created_at"`
	Items          []ManifestItem `json:"items"`
}

func (s *IntakeService) GetDocumentByManifestID(ctx context.Context, shortID string) (*ManifestDetail, error) {
	// 1. Try finding by Manifest No first
	manifest, err := s.repo.GetManifestByNo(ctx, shortID)
	if err == nil {
		items, _ := s.repo.GetManifestItems(ctx, manifest.ID)
		manifestItems := make([]ManifestItem, 0)
		for _, item := range items {
			manifestItems = append(manifestItems, ManifestItem{
				ID:             item.ID,
				DocumentID:     item.DocumentID,
				Title:          item.DocumentTitle,
				TypeName:       item.DocumentType.String,
				Date:           item.DocumentDate.Time.Format("02/01/2006"),
				Status:         item.Status,
				PhysicalStatus: item.CurrentPhysicalStatus.String,
				Notes:          item.Notes.String,
			})
		}

		return &ManifestDetail{
			ID:             manifest.ID,
			ManifestNo:     manifest.ManifestNo,
			Title:          fmt.Sprintf("Manifest %s", manifest.ManifestNo),
			TypeName:       "Manifest Serah Terima",
			OwnerName:      manifest.SenderName,
			DepartmentName: manifest.DepartmentName,
			Status:         manifest.Status,
			CreatedAt:      manifest.CreatedAt.Time.Format("02/01/2006 15:04"),
			Items:          manifestItems,
		}, nil
	}

	// 2. Fallback to Search by prefix (Legacy/Direct Document Scan)
	doc, err := s.repo.GetDocumentByShortID(ctx, pgtype.Text{String: strings.ToLower(shortID), Valid: true})
	if err != nil {
		return nil, fmt.Errorf("manifest or document not found with ID: %s", shortID)
	}

	return &ManifestDetail{
		ID:             uuid.Nil, // Not a saved manifest
		ManifestNo:     shortID,
		Title:          doc.Title,
		TypeName:       doc.TypeName.String,
		OwnerName:      doc.OwnerName,
		DepartmentName: doc.DepartmentName.String,
		Status:         doc.Status,
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
				Body:   notifMsg,
				Type:   "intake",
			})

			// 2. To Department Manager
			dept, err := s.repo.GetDepartment(ctx, manifest.DepartmentID)
			if err == nil && dept.ManagerID.Valid {
				s.notifSvc.CreateNotification(ctx, repository.CreateNotificationParams{
					UserID: dept.ManagerID.Bytes,
					Title:  "Notifikasi Penerimaan Berkas (Departemen)",
					Body:   fmt.Sprintf("Dokumen dari %s telah diterima oleh Central Document.", manifest.SenderName),
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
