package service

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
)

type HardwareService struct {
	repo *repository.Queries
}

func NewHardwareService(repo *repository.Queries) *HardwareService {
	return &HardwareService{repo: repo}
}

// RFID Tags
func (s *HardwareService) AssignRFID(ctx context.Context, tagID string, docID uuid.UUID) (repository.RfidTag, error) {
	return s.repo.CreateRfidTag(ctx, repository.CreateRfidTagParams{
		TagID:      tagID,
		DocumentID: docID,
		Status:     "active",
	})
}

func (s *HardwareService) ListRFID(ctx context.Context) ([]repository.RfidTag, error) {
	return s.repo.ListRfidTags(ctx)
}

// Label Generation
type LabelInfo struct {
	Code      string `json:"code"`
	Type      string `json:"type"`
	QRCode    string `json:"qr_code"` // Base64 or URL
	AssetName string `json:"asset_name"`
}

func (s *HardwareService) GenerateLabel(ctx context.Context, entityType string, entityID uuid.UUID) (*LabelInfo, error) {
	// Simulation: logic to generate a unique code and a mock QR
	code := fmt.Sprintf("DMS-%s-%s", entityType, entityID.String()[:8])
	
	return &LabelInfo{
		Code:      code,
		Type:      entityType,
		QRCode:    fmt.Sprintf("https://api.qrserver.com/v1/create-qr-code/?size=150x150&data=%s", code),
		AssetName: "Mock Physical Asset",
	}, nil
}
