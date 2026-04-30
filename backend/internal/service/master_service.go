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
	"io"
	"encoding/json"
	"net/netip"
)

type MasterService struct {
	repo *repository.Queries
}

func NewMasterService(repo *repository.Queries) *MasterService {
	return &MasterService{repo: repo}
}

func (s *MasterService) LogActivity(ctx context.Context, userID uuid.UUID, action, entityType string, entityID *uuid.UUID, details interface{}, ipAddress string) error {
	detailsJSON, _ := json.Marshal(details)
	
	params := repository.CreateActivityLogParams{
		UserID:     userID,
		Action:     action,
		EntityType: entityType,
		Details:    detailsJSON,
	}

	if entityID != nil {
		params.EntityID = pgtype.UUID{Bytes: *entityID, Valid: true}
	}

	if ipAddress != "" {
		if ip, err := netip.ParseAddr(ipAddress); err == nil {
			params.IpAddress = &ip
		}
	}

	_, err := s.repo.CreateActivityLog(ctx, params)
	return err
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

func (s *MasterService) ExportCompanies(ctx context.Context) ([]byte, string, error) {
	companies, err := s.repo.ListCompanies(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	// Header
	writer.Write([]string{"Entity ID", "Name", "NPWP Status", "Location", "Status", "Address"})

	for _, c := range companies {
		writer.Write([]string{
			c.EntityID.String,
			c.Name,
			c.NpwpStatus.String,
			c.Location.String,
			c.Status.String,
			c.Address.String,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("companies_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportCompanies(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 { // Skip header and invalid rows
			continue
		}

		entityID := record[0]
		name := record[1]
		npwpStatus := "PENDING"
		if len(record) > 2 {
			npwpStatus = record[2]
		}
		location := ""
		if len(record) > 3 {
			location = record[3]
		}
		status := "Aktif"
		if len(record) > 4 {
			status = record[4]
		}
		address := ""
		if len(record) > 5 {
			address = record[5]
		}

		if name == "" || entityID == "" {
			continue
		}

		_, err := s.repo.CreateCompany(ctx, repository.CreateCompanyParams{
			EntityID:   pgtype.Text{String: entityID, Valid: true},
			Name:       name,
			NpwpStatus: pgtype.Text{String: npwpStatus, Valid: true},
			Location:   pgtype.Text{String: location, Valid: true},
			Status:     pgtype.Text{String: status, Valid: true},
			Address:    pgtype.Text{String: address, Valid: true},
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

func (s *MasterService) UpdateBranch(ctx context.Context, id uuid.UUID, companyID uuid.UUID, name string, location string, headID *uuid.UUID) (repository.Branch, error) {
	params := repository.UpdateBranchParams{
		ID:        id,
		CompanyID: companyID,
		Name:      name,
		Location:  pgtype.Text{String: location, Valid: location != ""},
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.UpdateBranch(ctx, params)
}

func (s *MasterService) DeleteBranch(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBranch(ctx, id)
}

func (s *MasterService) ExportBranches(ctx context.Context) ([]byte, string, error) {
	branches, err := s.repo.ListAllBranchesGlobal(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	// Header
	writer.Write([]string{"Name", "Location", "Company Name", "Head Name"})

	for _, b := range branches {
		headName := "-"
		if b.HeadName.Valid {
			headName = b.HeadName.String
		}
		writer.Write([]string{
			b.Name,
			b.Location.String,
			b.CompanyName,
			headName,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("branches_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportBranches(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	// Get all companies to map name to ID
	companies, err := s.repo.ListCompanies(ctx)
	if err != nil {
		return 0, err
	}
	companyMap := make(map[string]uuid.UUID)
	for _, c := range companies {
		companyMap[c.Name] = c.ID
	}

	count := 0
	for i, record := range records {
		if i == 0 { // Skip header
			continue
		}

		if len(record) < 3 {
			continue
		}

		name := record[0]
		location := record[1]
		companyName := record[2]

		companyID, ok := companyMap[companyName]
		if !ok {
			// Skip or handle error
			continue
		}

		_, err := s.repo.CreateBranch(ctx, repository.CreateBranchParams{
			CompanyID: companyID,
			Name:      name,
			Location:  pgtype.Text{String: location, Valid: location != ""},
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

func (s *MasterService) UpdateDepartment(ctx context.Context, id uuid.UUID, branchID uuid.UUID, name string, headID *uuid.UUID) (repository.Department, error) {
	params := repository.UpdateDepartmentParams{
		ID:       id,
		BranchID: branchID,
		Name:     name,
	}
	if headID != nil {
		params.HeadID = pgtype.UUID{Bytes: *headID, Valid: true}
	}
	return s.repo.UpdateDepartment(ctx, params)
}

func (s *MasterService) DeleteDepartment(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDepartment(ctx, id)
}

func (s *MasterService) ExportDepartments(ctx context.Context) ([]byte, string, error) {
	depts, err := s.repo.ListAllDepartments(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)

	// Header
	writer.Write([]string{"Name", "Branch Name", "Head Name"})

	for _, d := range depts {
		headName := "-"
		if d.HeadName.Valid {
			headName = d.HeadName.String
		}
		writer.Write([]string{
			d.Name,
			d.BranchName,
			headName,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("departments_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportDepartments(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	branches, err := s.repo.ListAllBranchesGlobal(ctx)
	if err != nil {
		return 0, err
	}
	branchMap := make(map[string]uuid.UUID)
	for _, b := range branches {
		branchMap[b.Name] = b.ID
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}

		name := record[0]
		branchName := record[1]

		branchID, ok := branchMap[branchName]
		if !ok {
			continue
		}

		_, err := s.repo.CreateDepartment(ctx, repository.CreateDepartmentParams{
			BranchID: branchID,
			Name:     name,
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

func (s *MasterService) UpdateRack(ctx context.Context, id uuid.UUID, deptID uuid.UUID, name string, locationDetail string) (repository.Rack, error) {
	return s.repo.UpdateRack(ctx, repository.UpdateRackParams{
		ID:             id,
		DepartmentID:   deptID,
		Name:           name,
		LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
	})
}

func (s *MasterService) DeleteRack(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteRack(ctx, id)
}

func (s *MasterService) ExportRacks(ctx context.Context) ([]byte, string, error) {
	racks, err := s.repo.ListAllRacksGlobal(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"Name", "Department Name", "Branch Name", "Location Detail"})

	for _, r := range racks {
		writer.Write([]string{
			r.Name,
			r.DepartmentName,
			r.BranchName,
			r.LocationDetail.String,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("racks_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportRacks(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	depts, err := s.repo.ListAllDepartments(ctx)
	if err != nil {
		return 0, err
	}
	deptMap := make(map[string]uuid.UUID)
	for _, d := range depts {
		deptMap[d.Name] = d.ID
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}

		name := record[0]
		deptName := record[1]
		locationDetail := ""
		if len(record) > 3 {
			locationDetail = record[3]
		}

		deptID, ok := deptMap[deptName]
		if !ok {
			continue
		}

		_, err := s.repo.CreateRack(ctx, repository.CreateRackParams{
			DepartmentID:   deptID,
			Name:           name,
			LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

func (s *MasterService) UpdateBox(ctx context.Context, id uuid.UUID, rackID uuid.UUID, name string) (repository.Box, error) {
	return s.repo.UpdateBox(ctx, repository.UpdateBoxParams{
		ID:     id,
		RackID: rackID,
		Name:   name,
	})
}

func (s *MasterService) DeleteBox(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteBox(ctx, id)
}

func (s *MasterService) ExportBoxes(ctx context.Context) ([]byte, string, error) {
	boxes, err := s.repo.ListAllBoxesGlobal(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"Name", "Rack Name", "Department Name"})

	for _, b := range boxes {
		writer.Write([]string{
			b.Name,
			b.RackName,
			b.DepartmentName,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("boxes_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportBoxes(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	racks, err := s.repo.ListAllRacksGlobal(ctx)
	if err != nil {
		return 0, err
	}
	rackMap := make(map[string]uuid.UUID)
	for _, r := range racks {
		rackMap[r.Name] = r.ID
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}

		name := record[0]
		rackName := record[1]

		rackID, ok := rackMap[rackName]
		if !ok {
			continue
		}

		_, err := s.repo.CreateBox(ctx, repository.CreateBoxParams{
			RackID: rackID,
			Name:   name,
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

func (s *MasterService) UpdateOrdner(ctx context.Context, id uuid.UUID, boxID uuid.UUID, name string) (repository.Ordner, error) {
	return s.repo.UpdateOrdner(ctx, repository.UpdateOrdnerParams{
		ID:    id,
		BoxID: boxID,
		Name:  name,
	})
}

func (s *MasterService) DeleteOrdner(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteOrdner(ctx, id)
}

func (s *MasterService) ExportOrdners(ctx context.Context) ([]byte, string, error) {
	ordners, err := s.repo.ListAllOrdnersGlobal(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"Name", "Box Name", "Rack Name"})

	for _, o := range ordners {
		writer.Write([]string{
			o.Name,
			o.BoxName,
			o.RackName,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("ordners_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportOrdners(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	boxes, err := s.repo.ListAllBoxesGlobal(ctx)
	if err != nil {
		return 0, err
	}
	boxMap := make(map[string]uuid.UUID)
	for _, b := range boxes {
		boxMap[b.Name] = b.ID
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}

		name := record[0]
		boxName := record[1]

		boxID, ok := boxMap[boxName]
		if !ok {
			continue
		}

		_, err := s.repo.CreateOrdner(ctx, repository.CreateOrdnerParams{
			BoxID: boxID,
			Name:  name,
		})
		if err == nil {
			count++
		}
	}

	return count, nil
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

// Document Types
func (s *MasterService) ListDocumentTypes(ctx context.Context) ([]repository.DocumentType, error) {
	return s.repo.ListDocumentTypes(ctx)
}

func (s *MasterService) CreateDocumentType(ctx context.Context, code, name, description string) (repository.DocumentType, error) {
	return s.repo.CreateDocumentType(ctx, repository.CreateDocumentTypeParams{
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
}

func (s *MasterService) UpdateDocumentType(ctx context.Context, id uuid.UUID, code, name, description string) (repository.DocumentType, error) {
	return s.repo.UpdateDocumentType(ctx, repository.UpdateDocumentTypeParams{
		ID:          id,
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
}

func (s *MasterService) DeleteDocumentType(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDocumentType(ctx, id)
}

func (s *MasterService) ExportDocumentTypes(ctx context.Context) ([]byte, string, error) {
	types, err := s.repo.ListDocumentTypes(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"Code", "Name", "Description"})

	for _, t := range types {
		writer.Write([]string{
			t.Code,
			t.Name,
			t.Description.String,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("document_types_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportDocumentTypes(ctx context.Context, r io.Reader) (int, error) {
	reader := csv.NewReader(r)
	records, err := reader.ReadAll()
	if err != nil {
		return 0, err
	}

	if len(records) <= 1 {
		return 0, nil
	}

	count := 0
	for i, record := range records {
		if i == 0 || len(record) < 2 {
			continue
		}

		code := record[0]
		name := record[1]
		description := ""
		if len(record) > 2 {
			description = record[2]
		}

		_, err := s.repo.CreateDocumentType(ctx, repository.CreateDocumentTypeParams{
			Code:        code,
			Name:        name,
			Description: pgtype.Text{String: description, Valid: description != ""},
		})
		if err == nil {
			count++
		}
	}

	return count, nil
}

func (s *MasterService) ListActivityLogs(ctx context.Context, limit, offset int32) ([]repository.ListActivityLogsRow, error) {
	return s.repo.ListActivityLogs(ctx, repository.ListActivityLogsParams{
		Limit:  limit,
		Offset: offset,
	})
}
