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
	"github.com/kreatif/dms-backend/internal/infra"
	"io"
	"encoding/json"
	"net/netip"
	"strings"
)

type MasterService struct {
	repo       *repository.Queries
	ldapSvc    *infra.LDAPService
	aiSvc      *infra.AIService
	searchSvc  *infra.SearchService
	storageSvc *infra.StorageService
	waSvc      *infra.WhatsAppService
	emailSvc   *infra.EmailService
}

func NewMasterService(repo *repository.Queries, ldapSvc *infra.LDAPService, aiSvc *infra.AIService, searchSvc *infra.SearchService, storageSvc *infra.StorageService, waSvc *infra.WhatsAppService, emailSvc *infra.EmailService) *MasterService {
	return &MasterService{repo: repo, ldapSvc: ldapSvc, aiSvc: aiSvc, searchSvc: searchSvc, storageSvc: storageSvc, waSvc: waSvc, emailSvc: emailSvc}
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

func (s *MasterService) GetUser(ctx context.Context, id uuid.UUID) (repository.GetUserByIDRow, error) {
	return s.repo.GetUserByID(ctx, id)
}

func (s *MasterService) RegisterScanner(ctx context.Context, name, serviceType, endpoint string, config map[string]interface{}) (repository.IntegrationNode, error) {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return repository.IntegrationNode{}, err
	}

	// Use Endpoint as a unique key for local scanners (e.g., PC Hostname/ID + Scanner ID)
	// For network scanners, it's the IP.
	
	// Check if exists by endpoint and service type
	existing, err := s.repo.GetIntegrationNodeByEndpoint(ctx, repository.GetIntegrationNodeByEndpointParams{
		Endpoint:    endpoint,
		ServiceType: serviceType,
	})

	if err != nil {
		// Create new
		return s.repo.CreateIntegrationNode(ctx, repository.CreateIntegrationNodeParams{
			Name:        name,
			ServiceType: serviceType,
			Endpoint:    endpoint,
			ConfigJson:  configJSON,
		})
	}

	// Update existing
	return s.repo.UpdateIntegrationNodeConfig(ctx, repository.UpdateIntegrationNodeConfigParams{
		ID:         existing.ID,
		Name:       name,
		Endpoint:   endpoint,
		IsActive:   pgtype.Bool{Bool: true, Valid: true},
		ConfigJson: configJSON,
	})
}

// Company
func (s *MasterService) ListCompanies(ctx context.Context) ([]repository.Company, error) {
	return s.repo.ListCompanies(ctx)
}

func (s *MasterService) CreateCompany(ctx context.Context, name, entityID, npwpStatus, location, status, address, deliveryInstructions string) (repository.Company, error) {
	return s.repo.CreateCompany(ctx, repository.CreateCompanyParams{
		Name:                 name,
		EntityID:             pgtype.Text{String: entityID, Valid: entityID != ""},
		NpwpStatus:           pgtype.Text{String: npwpStatus, Valid: npwpStatus != ""},
		Location:             pgtype.Text{String: location, Valid: location != ""},
		Status:               pgtype.Text{String: status, Valid: status != ""},
		Address:              pgtype.Text{String: address, Valid: address != ""},
		DeliveryInstructions: pgtype.Text{String: deliveryInstructions, Valid: deliveryInstructions != ""},
	})
}

func (s *MasterService) UpdateCompany(ctx context.Context, id uuid.UUID, name, entityID, npwpStatus, location, status, address, deliveryInstructions string) (repository.Company, error) {
	return s.repo.UpdateCompany(ctx, repository.UpdateCompanyParams{
		ID:                   id,
		Name:                 name,
		EntityID:             pgtype.Text{String: entityID, Valid: entityID != ""},
		NpwpStatus:           pgtype.Text{String: npwpStatus, Valid: npwpStatus != ""},
		Location:             pgtype.Text{String: location, Valid: location != ""},
		Status:               pgtype.Text{String: status, Valid: status != ""},
		Address:              pgtype.Text{String: address, Valid: address != ""},
		DeliveryInstructions: pgtype.Text{String: deliveryInstructions, Valid: deliveryInstructions != ""},
	})
}

func (s *MasterService) DeleteCompany(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteCompany(ctx, id)
}

func (s *MasterService) UploadCompanyLogo(ctx context.Context, id uuid.UUID, file io.Reader, fileName string) (string, error) {
	// 1. Check encryption setting
	encryptEnabled := false
	encSetting, err := s.repo.GetSystemSetting(ctx, repository.GetSystemSettingParams{
		Category: "storage",
		Key:      "encryption_enabled",
	})
	if err == nil && encSetting.Value.String == "true" {
		encryptEnabled = true
	}

	// 2. Upload to storage (MinIO)
	objectPath := fmt.Sprintf("companies/%s/logo/%s", id, fileName)
	_, err = s.storageSvc.Upload(ctx, objectPath, file, -1, "image/png", encryptEnabled)
	if err != nil {
		return "", err
	}

	// 3. Update DB
	err = s.repo.UpdateCompanyLogo(ctx, repository.UpdateCompanyLogoParams{
		ID:      id,
		LogoUrl: pgtype.Text{String: objectPath, Valid: true},
	})
	
	return objectPath, err
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

func (s *MasterService) ListFloors(ctx context.Context) ([]repository.Floor, error) {
	return s.repo.ListFloors(ctx)
}

func (s *MasterService) CreateFloor(ctx context.Context, name string, code string) (repository.Floor, error) {
	return s.repo.CreateFloor(ctx, repository.CreateFloorParams{
		Name: name,
		Code: code,
	})
}

func (s *MasterService) UpdateFloor(ctx context.Context, id uuid.UUID, name string, code string) (repository.Floor, error) {
	return s.repo.UpdateFloor(ctx, repository.UpdateFloorParams{
		ID:   id,
		Name: name,
		Code: code,
	})
}

func (s *MasterService) DeleteFloor(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteFloor(ctx, id)
}

func (s *MasterService) CreateRack(ctx context.Context, deptID uuid.UUID, name string, locationDetail string, floorID *uuid.UUID, mapPosX *float64, mapPosY *float64) (repository.Rack, error) {
	var fID pgtype.UUID
	if floorID != nil {
		fID = pgtype.UUID{Bytes: *floorID, Valid: true}
	}
	var px, py pgtype.Numeric
	if mapPosX != nil {
		px.Scan(fmt.Sprintf("%.2f", *mapPosX))
	}
	if mapPosY != nil {
		py.Scan(fmt.Sprintf("%.2f", *mapPosY))
	}

	return s.repo.CreateRack(ctx, repository.CreateRackParams{
		DepartmentID:   deptID,
		Name:           name,
		LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
		FloorID:        fID,
		MapPosX:        px,
		MapPosY:        py,
	})
}

func (s *MasterService) UpdateRack(ctx context.Context, id uuid.UUID, deptID uuid.UUID, name string, locationDetail string, floorID *uuid.UUID, mapPosX *float64, mapPosY *float64) (repository.Rack, error) {
	var fID pgtype.UUID
	if floorID != nil {
		fID = pgtype.UUID{Bytes: *floorID, Valid: true}
	}
	var px, py pgtype.Numeric
	if mapPosX != nil {
		px.Scan(fmt.Sprintf("%.2f", *mapPosX))
	}
	if mapPosY != nil {
		py.Scan(fmt.Sprintf("%.2f", *mapPosY))
	}

	return s.repo.UpdateRack(ctx, repository.UpdateRackParams{
		ID:             id,
		DepartmentID:   deptID,
		Name:           name,
		LocationDetail: pgtype.Text{String: locationDetail, Valid: locationDetail != ""},
		FloorID:        fID,
		MapPosX:        px,
		MapPosY:        py,
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
			FloorID:        pgtype.UUID{Valid: false},
			MapPosX:        pgtype.Numeric{Valid: false},
			MapPosY:        pgtype.Numeric{Valid: false},
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
func (s *MasterService) ListRoles(ctx context.Context) ([]repository.ListRolesRow, error) {
	return s.repo.ListRoles(ctx)
}

func (s *MasterService) ListSystemModules(ctx context.Context) ([]repository.SystemModule, error) {
	return s.repo.ListSystemModules(ctx)
}

func (s *MasterService) GetSystemModule(ctx context.Context, id string) (repository.SystemModule, error) {
	return s.repo.GetSystemModule(ctx, id)
}

func (s *MasterService) CreateSystemModule(ctx context.Context, id, name, category, path, icon string, allowedActions []string, sortOrder int32, parentID *string) (repository.SystemModule, error) {
	params := repository.CreateSystemModuleParams{
		ID:             id,
		Name:           name,
		Category:       category,
		Path:           pgtype.Text{String: path, Valid: path != ""},
		Icon:           pgtype.Text{String: icon, Valid: icon != ""},
		AllowedActions: allowedActions,
		SortOrder:      pgtype.Int4{Int32: sortOrder, Valid: true},
	}
	if parentID != nil {
		params.ParentID = pgtype.Text{String: *parentID, Valid: *parentID != ""}
	}
	return s.repo.CreateSystemModule(ctx, params)
}

func (s *MasterService) UpdateSystemModule(ctx context.Context, id, name, category, path, icon string, allowedActions []string, sortOrder int32, parentID *string) (repository.SystemModule, error) {
	params := repository.UpdateSystemModuleParams{
		ID:             id,
		Name:           name,
		Category:       category,
		Path:           pgtype.Text{String: path, Valid: path != ""},
		Icon:           pgtype.Text{String: icon, Valid: icon != ""},
		AllowedActions: allowedActions,
		SortOrder:      pgtype.Int4{Int32: sortOrder, Valid: true},
	}
	if parentID != nil {
		params.ParentID = pgtype.Text{String: *parentID, Valid: *parentID != ""}
	}
	return s.repo.UpdateSystemModule(ctx, params)
}

func (s *MasterService) DeleteSystemModule(ctx context.Context, id string) error {
	return s.repo.DeleteSystemModule(ctx, id)
}

func (s *MasterService) GetRolePermissions(ctx context.Context, roleID int32) ([]repository.GetRolePermissionsRow, error) {
	return s.repo.GetRolePermissions(ctx, roleID)
}

func (s *MasterService) UpdateRolePermissions(ctx context.Context, roleID int32, permissions []struct {
	ModuleID string `json:"module_id"`
	Action   string `json:"action"`
}) error {
	// Start transaction if possible, but for simplicity we'll just clear and add
	// Actually, repo.Queries usually don't have transaction support out of the box unless we use pgxpool.Begin
	// We'll use the repo's existing interface.
	
	err := s.repo.ClearRolePermissions(ctx, roleID)
	if err != nil {
		return err
	}

	for _, p := range permissions {
		err = s.repo.AddRolePermission(ctx, repository.AddRolePermissionParams{
			RoleID:   roleID,
			ModuleID: p.ModuleID,
			Action:   p.Action,
		})
		if err != nil {
			return err
		}
	}

	return nil
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

// Announcements Management
func (s *MasterService) ListAnnouncements(ctx context.Context) ([]repository.Announcement, error) {
	return s.repo.ListAnnouncements(ctx)
}

func (s *MasterService) CreateAnnouncement(ctx context.Context, title, message, notes string, isActive bool, userID uuid.UUID) (repository.Announcement, error) {
	if isActive {
		_ = s.repo.DeactivateAllAnnouncements(ctx)
	}
	return s.repo.CreateAnnouncement(ctx, repository.CreateAnnouncementParams{
		Title:     title,
		Message:   message,
		Notes:     pgtype.Text{String: notes, Valid: true},
		IsActive:  isActive,
		CreatedBy: pgtype.UUID{Bytes: userID, Valid: true},
	})
}

func (s *MasterService) UpdateAnnouncement(ctx context.Context, id uuid.UUID, title, message, notes string, isActive bool) (repository.Announcement, error) {
	if isActive {
		_ = s.repo.DeactivateAllAnnouncements(ctx)
	}
	return s.repo.UpdateAnnouncement(ctx, repository.UpdateAnnouncementParams{
		ID:       id,
		Title:    title,
		Message:  message,
		Notes:    pgtype.Text{String: notes, Valid: true},
		IsActive: isActive,
	})
}

func (s *MasterService) DeleteAnnouncement(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteAnnouncement(ctx, id)
}

// Warehouse Topology
type TopologyNode struct {
	ID       string         `json:"id"`
	Name     string         `json:"name"`
	Type     string         `json:"type"`
	Children []TopologyNode `json:"children,omitempty"`
}

func (s *MasterService) GetTopology(ctx context.Context, deptID uuid.UUID) ([]TopologyNode, error) {
	rows, err := s.repo.GetWarehouseTopology(ctx)
	if err != nil {
		return nil, err
	}

	return s.buildTopologyTree(rows, deptID), nil
}

func (s *MasterService) buildTopologyTree(rows []repository.GetWarehouseTopologyRow, deptID uuid.UUID) []TopologyNode {
	type ordnerNode struct {
		id   string
		name string
	}
	type boxNode struct {
		id      string
		name    string
		ordners map[string]ordnerNode
	}
	type rackNode struct {
		id    string
		name  string
		boxes map[string]boxNode
	}
	type deptNode struct {
		id    string
		name  string
		racks map[string]rackNode
	}
	type branchNode struct {
		id    string
		name  string
		depts map[string]deptNode
	}
	type companyNode struct {
		id       string
		name     string
		branches map[string]branchNode
	}

	companies := make(map[string]companyNode)

	for _, row := range rows {
		cID := row.CompanyID.String()
		if _, ok := companies[cID]; !ok {
			companies[cID] = companyNode{id: cID, name: row.CompanyName, branches: make(map[string]branchNode)}
		}

		bID := row.BranchID.String()
		if _, ok := companies[cID].branches[bID]; !ok {
			companies[cID].branches[bID] = branchNode{id: bID, name: row.BranchName, depts: make(map[string]deptNode)}
		}

		dID := row.DepartmentID.String()
		// Filter by department if requested
		if deptID != uuid.Nil && dID != deptID.String() {
			continue
		}

		if _, ok := companies[cID].branches[bID].depts[dID]; !ok {
			companies[cID].branches[bID].depts[dID] = deptNode{id: dID, name: row.DepartmentName, racks: make(map[string]rackNode)}
		}

		rID := row.RackID.String()
		if _, ok := companies[cID].branches[bID].depts[dID].racks[rID]; !ok {
			companies[cID].branches[bID].depts[dID].racks[rID] = rackNode{id: rID, name: row.RackName, boxes: make(map[string]boxNode)}
		}

		if row.BoxID.Valid {
			bxID := uuid.UUID(row.BoxID.Bytes).String()
			if _, ok := companies[cID].branches[bID].depts[dID].racks[rID].boxes[bxID]; !ok {
				companies[cID].branches[bID].depts[dID].racks[rID].boxes[bxID] = boxNode{id: bxID, name: row.BoxName.String, ordners: make(map[string]ordnerNode)}
			}

			if row.OrdnerID.Valid {
				oID := uuid.UUID(row.OrdnerID.Bytes).String()
				companies[cID].branches[bID].depts[dID].racks[rID].boxes[bxID].ordners[oID] = ordnerNode{id: oID, name: row.OrdnerName.String}
			}
		}
	}

	// Convert maps to slices
	var tree []TopologyNode
	for _, c := range companies {
		cNode := TopologyNode{ID: c.id, Name: c.name, Type: "company"}
		for _, b := range c.branches {
			bNode := TopologyNode{ID: b.id, Name: b.name, Type: "branch"}
			for _, d := range b.depts {
				dNode := TopologyNode{ID: d.id, Name: d.name, Type: "department"}
				for _, r := range d.racks {
					rNode := TopologyNode{ID: r.id, Name: r.name, Type: "rack"}
					for _, bx := range r.boxes {
						bxNode := TopologyNode{ID: bx.id, Name: bx.name, Type: "box"}
						for _, o := range bx.ordners {
							bxNode.Children = append(bxNode.Children, TopologyNode{ID: o.id, Name: o.name, Type: "ordner"})
						}
						rNode.Children = append(rNode.Children, bxNode)
					}
					dNode.Children = append(dNode.Children, rNode)
				}
				bNode.Children = append(bNode.Children, dNode)
			}
			cNode.Children = append(cNode.Children, bNode)
		}
		tree = append(tree, cNode)
	}

	return tree
}

func (s *MasterService) GetIntegrationStatus(ctx context.Context) ([]repository.IntegrationNode, error) {
	return s.repo.ListIntegrationNodes(ctx)
}

func (s *MasterService) CreateIntegrationNode(ctx context.Context, name, serviceType, driver, endpoint string, isCritical bool, config []byte) (repository.IntegrationNode, error) {
	return s.repo.CreateIntegrationNode(ctx, repository.CreateIntegrationNodeParams{
		Name:        name,
		ServiceType: serviceType,
		Driver:      pgtype.Text{String: driver, Valid: driver != ""},
		Endpoint:    endpoint,
		IsCritical:  pgtype.Bool{Bool: isCritical, Valid: true},
		ConfigJson:  config,
	})
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

func (s *MasterService) DeleteIntegrationNode(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteIntegrationNode(ctx, id)
}

type LDAPTestResult struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Users   int    `json:"users"`
	Groups  int    `json:"groups"`
}

func (s *MasterService) TestLDAPConnection(ctx context.Context, endpoint string, config []byte) (*LDAPTestResult, error) {
	var cfg struct {
		BaseDN       string `json:"base_dn"`
		BindDN       string `json:"bind_dn"`
		BindPass     string `json:"bind_pass"`
		BindPassword string `json:"bind_password"`
		UserFilter   string `json:"user_filter"`
	}
	if err := json.Unmarshal(config, &cfg); err != nil {
		return nil, err
	}

	// Fallback for password field names
	finalPassword := cfg.BindPass
	if finalPassword == "" {
		finalPassword = cfg.BindPassword
	}

	// Implementation: Use ldapSvc to perform real connection test
	users, groups, err := s.ldapSvc.TestConnection(ctx, endpoint, cfg.BaseDN, cfg.BindDN, finalPassword)
	if err != nil {
		// Log Failure to sso_sync_logs
		lastLog, _ := s.repo.CreateSsoSyncLog(ctx, repository.CreateSsoSyncLogParams{
			Provider:  "LDAP",
			Status:    "failed",
			StartedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})
		
		errorDetails, _ := json.Marshal(map[string]string{"error": err.Error()})
		_, _ = s.repo.UpdateSsoSyncLog(ctx, repository.UpdateSsoSyncLogParams{
			ID:           lastLog.ID,
			Status:       "failed",
			ErrorDetails: errorDetails,
			Errors:       pgtype.Int4{Int32: 1, Valid: true},
			CompletedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
		})

		return &LDAPTestResult{
			Success: false,
			Message: err.Error(),
			Users:   0,
			Groups:  0,
		}, err
	}

	res := &LDAPTestResult{
		Success: true,
		Message: "Connection established successfully",
		Users:   users,
		Groups:  groups,
	}

	// Log Success to sso_sync_logs
	lastLog, _ := s.repo.CreateSsoSyncLog(ctx, repository.CreateSsoSyncLogParams{
		Provider:  "LDAP",
		Status:    "success",
		StartedAt: pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})
	
	_, _ = s.repo.UpdateSsoSyncLog(ctx, repository.UpdateSsoSyncLogParams{
		ID:           lastLog.ID,
		Status:       "success",
		UsersSynced:  pgtype.Int4{Int32: int32(users), Valid: true},
		GroupsSynced: pgtype.Int4{Int32: int32(groups), Valid: true},
		Errors:       pgtype.Int4{Int32: 0, Valid: true},
		CompletedAt:  pgtype.Timestamptz{Time: time.Now(), Valid: true},
	})

	return res, nil
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

// Document Categories
func (s *MasterService) ListDocumentCategories(ctx context.Context) ([]repository.DocumentCategory, error) {
	return s.repo.ListDocumentCategories(ctx)
}

func (s *MasterService) CreateDocumentCategory(ctx context.Context, code, name, description string) (repository.DocumentCategory, error) {
	return s.repo.CreateDocumentCategory(ctx, repository.CreateDocumentCategoryParams{
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
}

func (s *MasterService) UpdateDocumentCategory(ctx context.Context, id uuid.UUID, code, name, description string) (repository.DocumentCategory, error) {
	return s.repo.UpdateDocumentCategory(ctx, repository.UpdateDocumentCategoryParams{
		ID:          id,
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	})
}

func (s *MasterService) DeleteDocumentCategory(ctx context.Context, id uuid.UUID) error {
	return s.repo.DeleteDocumentCategory(ctx, id)
}

func (s *MasterService) ExportDocumentCategories(ctx context.Context) ([]byte, string, error) {
	categories, err := s.repo.ListDocumentCategories(ctx)
	if err != nil {
		return nil, "", err
	}

	buf := new(bytes.Buffer)
	writer := csv.NewWriter(buf)
	writer.Write([]string{"Code", "Name", "Description"})

	for _, c := range categories {
		writer.Write([]string{
			c.Code,
			c.Name,
			c.Description.String,
		})
	}

	writer.Flush()
	fileName := fmt.Sprintf("document_categories_%s.csv", time.Now().Format("20060102_150405"))
	return buf.Bytes(), fileName, nil
}

func (s *MasterService) ImportDocumentCategories(ctx context.Context, r io.Reader) (int, error) {
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

		_, err := s.repo.CreateDocumentCategory(ctx, repository.CreateDocumentCategoryParams{
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

type DocumentTypeResponse struct {
	ID           uuid.UUID `json:"id"`
	Code         string    `json:"code"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CategoryID   *uuid.UUID `json:"category_id"`
	CategoryName string    `json:"category_name"`
}

// Document Types
func (s *MasterService) ListDocumentTypes(ctx context.Context) ([]DocumentTypeResponse, error) {
	rows, err := s.repo.ListDocumentTypes(ctx)
	if err != nil {
		return nil, err
	}

	res := make([]DocumentTypeResponse, 0, len(rows))
	for _, r := range rows {
		var catID *uuid.UUID
		if r.CategoryID.Valid {
			id := uuid.UUID(r.CategoryID.Bytes)
			catID = &id
		}

		res = append(res, DocumentTypeResponse{
			ID:           r.ID,
			Code:         r.Code,
			Name:         r.Name,
			Description:  r.Description.String,
			CategoryID:   catID,
			CategoryName: r.CategoryName.String,
		})
	}
	return res, nil
}

func (s *MasterService) CreateDocumentType(ctx context.Context, code, name, description string, categoryID *uuid.UUID) (repository.DocumentType, error) {
	params := repository.CreateDocumentTypeParams{
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	}
	if categoryID != nil {
		params.CategoryID = pgtype.UUID{Bytes: *categoryID, Valid: true}
	}
	return s.repo.CreateDocumentType(ctx, params)
}

func (s *MasterService) UpdateDocumentType(ctx context.Context, id uuid.UUID, code, name, description string, categoryID *uuid.UUID) (repository.DocumentType, error) {
	params := repository.UpdateDocumentTypeParams{
		ID:          id,
		Code:        code,
		Name:        name,
		Description: pgtype.Text{String: description, Valid: description != ""},
	}
	if categoryID != nil {
		params.CategoryID = pgtype.UUID{Bytes: *categoryID, Valid: true}
	}
	return s.repo.UpdateDocumentType(ctx, params)
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

func (s *MasterService) GetZonationLogs(ctx context.Context, limit, offset int32) ([]repository.GetZonationLogsRow, error) {
	return s.repo.GetZonationLogs(ctx, repository.GetZonationLogsParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (s *MasterService) ListSsoSyncLogs(ctx context.Context, limit, offset int32) ([]repository.SsoSyncLog, error) {
	return s.repo.ListSsoSyncLogs(ctx, repository.ListSsoSyncLogsParams{
		Limit:  limit,
		Offset: offset,
	})
}

func (s *MasterService) FetchAIModels(ctx context.Context, driver string, apiKey string) ([]map[string]string, error) {
	if driver == "gemini" {
		return s.aiSvc.FetchGeminiModels(ctx, apiKey)
	} else if driver == "openai" {
		return s.aiSvc.FetchOpenAIModels(ctx, apiKey)
	}
	return nil, fmt.Errorf("unsupported AI driver: %s", driver)
}

func (s *MasterService) GetSyncLogsCount(ctx context.Context) (int64, error) {
	return s.repo.CountSsoSyncLogs(ctx)
}

func (s *MasterService) TestSearchConnection(ctx context.Context, endpoint string, config []byte) (map[string]interface{}, error) {
	var cfg struct {
		Username string `json:"username"`
		Password string `json:"password"`
		APIKey   string `json:"api_key"`
	}
	if err := json.Unmarshal(config, &cfg); err != nil {
		return nil, err
	}

	addresses := []string{endpoint}
	return s.searchSvc.TestConnection(ctx, addresses, cfg.Username, cfg.Password, cfg.APIKey)
}

func (s *MasterService) TestStorageConnection(ctx context.Context, endpoint string, config []byte) error {
	var cfg struct {
		AccessKey string      `json:"access_key"`
		SecretKey string      `json:"secret_key"`
		Bucket    string      `json:"bucket"`
		UseSSL    interface{} `json:"use_ssl"`
	}
	if err := json.Unmarshal(config, &cfg); err != nil {
		return err
	}

	useSSL := false
	switch v := cfg.UseSSL.(type) {
	case bool:
		useSSL = v
	case string:
		useSSL = (v == "true" || v == "1")
	}

	return s.storageSvc.TestConnection(ctx, endpoint, cfg.AccessKey, cfg.SecretKey, cfg.Bucket, useSSL)
}

func (s *MasterService) TestWhatsAppConnection(ctx context.Context, endpoint string, config []byte) (map[string]interface{}, error) {
	var cfg struct {
		Token string `json:"token"`
	}
	if err := json.Unmarshal(config, &cfg); err != nil {
		return nil, err
	}

	return s.waSvc.TestConnection(ctx, endpoint, cfg.Token)
}

func (s *MasterService) TestSMTPConnection(ctx context.Context, endpoint string, config []byte) error {
	var cfg struct {
		Auth bool   `json:"auth"`
		User string `json:"user"`
		Pass string `json:"pass"`
	}
	if err := json.Unmarshal(config, &cfg); err != nil {
		return err
	}

	host := strings.Split(endpoint, ":")[0]
	return s.emailSvc.TestConnection(ctx, endpoint, host, cfg.User, cfg.Pass, cfg.Auth)
}
func (s *MasterService) GetWatermarkSettings(ctx context.Context) (map[string]interface{}, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "WATERMARK")
	if err != nil {
		return map[string]interface{}{
			"type":     "text",
			"text":     "CONFIDENTIAL - {user} - {date}",
			"opacity":  0.3,
			"position": "diagonal",
		}, nil
	}

	var config map[string]interface{}
	if err := json.Unmarshal(node.ConfigJson, &config); err != nil {
		return nil, err
	}

	return config, nil
}

func (s *MasterService) UpdateWatermarkSettings(ctx context.Context, config map[string]interface{}) error {
	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}

	// Try to find existing node
	node, err := s.repo.GetIntegrationNodeByType(ctx, "WATERMARK")
	if err != nil {
		// Create new node
		_, err = s.repo.CreateIntegrationNode(ctx, repository.CreateIntegrationNodeParams{
			Name:        "Watermark Configuration",
			ServiceType: "WATERMARK",
			Endpoint:    "internal",
			ConfigJson:  configJSON,
		})
		return err
	}

	// Update existing node
	_, err = s.repo.UpdateIntegrationNodeConfig(ctx, repository.UpdateIntegrationNodeConfigParams{
		ID:         node.ID,
		Name:       "Watermark Configuration",
		Endpoint:   "internal",
		IsActive:   pgtype.Bool{Bool: true, Valid: true},
		ConfigJson: configJSON,
	})
	return err
}
func (s *MasterService) UpdateDepartmentFloorPlan(ctx context.Context, id uuid.UUID, floorPlanUrl string) (repository.Department, error) {
	urlArg := pgtype.Text{String: floorPlanUrl, Valid: floorPlanUrl != ""}
	return s.repo.UpdateDepartmentFloorPlan(ctx, repository.UpdateDepartmentFloorPlanParams{
		ID:           id,
		FloorPlanUrl: urlArg,
	})
}

func (s *MasterService) UpdateRackMapCoordinates(ctx context.Context, id uuid.UUID, posX, posY float64) (repository.Rack, error) {
	var px, py pgtype.Numeric
	px.Scan(fmt.Sprintf("%.2f", posX))
	py.Scan(fmt.Sprintf("%.2f", posY))

	return s.repo.UpdateRackMapCoordinates(ctx, repository.UpdateRackMapCoordinatesParams{
		ID:      id,
		MapPosX: px,
		MapPosY: py,
	})
}

func (s *MasterService) UpdateRackOverrideStatus(ctx context.Context, id uuid.UUID, isFull bool, reason string) (repository.Rack, error) {
	return s.repo.UpdateRackOverrideStatus(ctx, repository.UpdateRackOverrideStatusParams{
		ID:             id,
		IsFullOverride: pgtype.Bool{Bool: isFull, Valid: true},
		OverrideReason: pgtype.Text{String: reason, Valid: reason != ""},
	})
}
