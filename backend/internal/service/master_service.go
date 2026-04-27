package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
)

type MasterService struct {
	repo *repository.Queries
}

func NewMasterService(repo *repository.Queries) *MasterService {
	return &MasterService{repo: repo}
}

// Company
func (s *MasterService) ListCompanies(ctx context.Context) ([]repository.Company, error) {
	return s.repo.ListCompanies(ctx)
}

func (s *MasterService) CreateCompany(ctx context.Context, name, entityID, npwpStatus, location, status, address string) (repository.Company, error) {
	return s.repo.CreateCompany(ctx, repository.CreateCompanyParams{
		Name:       name,
		EntityID:   pgtype.Text{String: entityID, Valid: entityID != ""},
		NpwpStatus: pgtype.Text{String: npwpStatus, Valid: npwpStatus != ""},
		Location:   pgtype.Text{String: location, Valid: location != ""},
		Status:     pgtype.Text{String: status, Valid: status != ""},
		Address:    pgtype.Text{String: address, Valid: address != ""},
	})
}

func (s *MasterService) UpdateCompany(ctx context.Context, id uuid.UUID, name, entityID, npwpStatus, location, status, address string) (repository.Company, error) {
	return s.repo.UpdateCompany(ctx, repository.UpdateCompanyParams{
		ID:         id,
		Name:       name,
		EntityID:   pgtype.Text{String: entityID, Valid: entityID != ""},
		NpwpStatus: pgtype.Text{String: npwpStatus, Valid: npwpStatus != ""},
		Location:   pgtype.Text{String: location, Valid: location != ""},
		Status:     pgtype.Text{String: status, Valid: status != ""},
		Address:    pgtype.Text{String: address, Valid: address != ""},
	})
}

func (s *MasterService) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCompany(ctx, id)
}

// Branch
func (s *MasterService) ListBranches(ctx context.Context, companyID uuid.UUID) ([]repository.Branch, error) {
	return s.repo.ListBranches(ctx, companyID)
}

// Department
func (s *MasterService) ListDepartments(ctx context.Context, branchID uuid.UUID) ([]repository.Department, error) {
	return s.repo.ListDepartments(ctx, branchID)
}

// Roles
func (s *MasterService) ListRoles(ctx context.Context) ([]repository.Role, error) {
	return s.repo.ListRoles(ctx)
}

// Retention Policies
func (s *MasterService) ListRetentionPolicies(ctx context.Context) ([]repository.RetentionPolicy, error) {
	return s.repo.ListRetentionPolicies(ctx)
}

// System Settings
func (s *MasterService) GetSettings(ctx context.Context, category string) ([]repository.SystemSetting, error) {
	return s.repo.GetSystemSettingsByCategory(ctx, category)
}

func (s *MasterService) UpdateSetting(ctx context.Context, category, key, value, valueType, desc string, userID uuid.UUID) (repository.SystemSetting, error) {
	return s.repo.UpsertSystemSetting(ctx, repository.UpsertSystemSettingParams{
		Category:    category,
		Key:         key,
		Value:       pgtype.Text{String: value, Valid: true},
		ValueType:   pgtype.Text{String: valueType, Valid: true},
		Description: pgtype.Text{String: desc, Valid: true},
		UpdatedBy:   pgtype.UUID{Bytes: userID, Valid: true},
	})
}

// Warehouse Topology
type TopologyNode struct {
	Name     string         `json:"name"`
	Type     string         `json:"type"`
	Children []TopologyNode `json:"children,omitempty"`
}

func (s *MasterService) GetTopology(ctx context.Context) ([]TopologyNode, error) {
	rows, err := s.repo.GetWarehouseTopology(ctx)
	if err != nil {
		return nil, err
	}

	// Build tree from flat rows
	// For simplicity in this example, we just return a nested structure
	// In a real app, you'd use maps to build the hierarchy
	return s.buildTopologyTree(rows), nil
}

func (s *MasterService) buildTopologyTree(rows []repository.GetWarehouseTopologyRow) []TopologyNode {
	// Simple grouping logic for simulation
	// In a real scenario, this would be a recursive map-based builder
	var tree []TopologyNode
	companies := make(map[string]*TopologyNode)

	for _, row := range rows {
		if _, ok := companies[row.CompanyName]; !ok {
			companies[row.CompanyName] = &TopologyNode{Name: row.CompanyName, Type: "company"}
		}
		// ... logic to nest branches, depts, racks ...
	}
	
	for _, c := range companies {
		tree = append(tree, *c)
	}
	return tree
}
