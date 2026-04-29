package service

import (
	"context"
	"encoding/csv"
	"fmt"
	"time"
	"bytes"
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

func (s *MasterService) ListAllBranchesGlobal(ctx context.Context) ([]repository.ListAllBranchesGlobalRow, error) {
	return s.repo.ListAllBranchesGlobal(ctx)
}

func (s *MasterService) CreateBranch(ctx context.Context, companyID uuid.UUID, name string, location string, headID *uuid.UUID) (repository.Branch, error) {
	params := repository.CreateBranchParams{
		CompanyID: companyID,
		Name:      name,
		Location:  pgtype.Text{String: location, Valid: location != ""},
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.CreateBranch(ctx, params)
}

func (s *MasterService) UpdateBranch(ctx context.Context, id uuid.UUID, name string, location string, headID *uuid.UUID) (repository.Branch, error) {
	params := repository.UpdateBranchParams{
		ID:       id,
		Name:     name,
		Location: pgtype.Text{String: location, Valid: location != ""},
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.UpdateBranch(ctx, params)
}

func (s *MasterService) DeleteBranch(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBranch(ctx, id)
}

// Department
func (s *MasterService) ListDepartments(ctx context.Context, branchID uuid.UUID) ([]repository.Department, error) {
	return s.repo.ListDepartments(ctx, branchID)
}

func (s *MasterService) ListAllDepartments(ctx context.Context) ([]repository.ListAllDepartmentsRow, error) {
	return s.repo.ListAllDepartments(ctx)
}

func (s *MasterService) CreateDepartment(ctx context.Context, branchID uuid.UUID, name string, headID *uuid.UUID) (repository.Department, error) {
	params := repository.CreateDepartmentParams{
		BranchID: branchID,
		Name:     name,
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.CreateDepartment(ctx, params)
}

func (s *MasterService) UpdateDepartment(ctx context.Context, id uuid.UUID, name string, headID *uuid.UUID) (repository.Department, error) {
	params := repository.UpdateDepartmentParams{
		ID:   id,
		Name: name,
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.UpdateDepartment(ctx, params)
}

func (s *MasterService) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDepartment(ctx, id)
}

// Racks
func (s *MasterService) ListAllRacksGlobal(ctx context.Context) ([]repository.ListAllRacksGlobalRow, error) {
	return s.repo.ListAllRacksGlobal(ctx)
}

func (s *MasterService) CreateRack(ctx context.Context, deptID uuid.UUID, name string, locationDetail string) (repository.Rack, error) {
	return s.repo.CreateRack(ctx, repository.CreateRackParams{
		DepartmentID:   deptID,
		Name:           name,
		LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
	})
}

func (s *MasterService) UpdateRack(ctx context.Context, id uuid.UUID, name string, locationDetail string) (repository.Rack, error) {
	return s.repo.UpdateRack(ctx, repository.UpdateRackParams{
		ID:             id,
		Name:           name,
		LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
	})
}

func (s *MasterService) DeleteRack(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteRack(ctx, id)
}

// Boxes
func (s *MasterService) ListAllBoxesGlobal(ctx context.Context) ([]repository.ListAllBoxesGlobalRow, error) {
	return s.repo.ListAllBoxesGlobal(ctx)
}

func (s *MasterService) CreateBox(ctx context.Context, rackID uuid.UUID, name string) (repository.Box, error) {
	return s.repo.CreateBox(ctx, repository.CreateBoxParams{
		RackID: rackID,
		Name:   name,
	})
}

func (s *MasterService) UpdateBox(ctx context.Context, id uuid.UUID, name string) (repository.Box, error) {
	return s.repo.UpdateBox(ctx, repository.UpdateBoxParams{
		ID:   id,
		Name: name,
	})
}

func (s *MasterService) DeleteBox(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBox(ctx, id)
}

// Ordners
func (s *MasterService) ListAllOrdnersGlobal(ctx context.Context) ([]repository.ListAllOrdnersGlobalRow, error) {
	return s.repo.ListAllOrdnersGlobal(ctx)
}

func (s *MasterService) CreateOrdner(ctx context.Context, boxID uuid.UUID, name string) (repository.Ordner, error) {
	return s.repo.CreateOrdner(ctx, repository.CreateOrdnerParams{
		BoxID: boxID,
		Name:  name,
	})
}

func (s *MasterService) UpdateOrdner(ctx context.Context, id uuid.UUID, name string) (repository.Ordner, error) {
	return s.repo.UpdateOrdner(ctx, repository.UpdateOrdnerParams{
		ID:   id,
		Name: name,
	})
}

func (s *MasterService) DeleteOrdner(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteOrdner(ctx, id)
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

func (s *MasterService) GetIntegrationStatus(ctx context.Context) ([]repository.IntegrationNode, error) {
	return s.repo.ListIntegrationNodes(ctx)
}

func (s *MasterService) UpdateIntegrationNode(ctx context.Context, id uuid.UUID, name, endpoint string, isActive, isCritical bool, config []byte) (repository.IntegrationNode, error) {
	return s.repo.UpdateIntegrationNodeConfig(ctx, repository.UpdateIntegrationNodeConfigParams{
		ID:         id,
		Name:       name,
		Endpoint:   endpoint,
		IsActive:   pgtype.Bool{Bool: isActive, Valid: true},
		IsCritical: pgtype.Bool{Bool: isCritical, Valid: true},
		ConfigJson: config,
	})
}

func (s *MasterService) ExportIntegrationReport(ctx context.Context) ([]byte, string, error) {
	nodes, err := s.repo.ListIntegrationNodes(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	// Header
	writer.Write([]string{"Service Name", "Type", "Endpoint", "Status", "Last Latency (ms)", "Last Check", "Critical"})

	for _, node := range nodes {
		lastCheck := "Never"
		if node.LastCheckAt.Valid {
			lastCheck = node.LastCheckAt.Time.Format("2006-01-02 15:04:05")
		}

		writer.Write([]string{
			node.Name,
			node.ServiceType,
			node.Endpoint,
			node.Status.String,
			fmt.Sprintf("%d", node.LastLatency.Int32),
			lastCheck,
			fmt.Sprintf("%t", node.IsCritical.Bool),
		})
	}

	writer.Flush()
	
	fileName := fmt.Sprintf("integration_report_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}
