package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	_ "github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type MasterHandler struct {
	svc *service.MasterService
}

func NewMasterHandler(svc *service.MasterService) *MasterHandler {
	return &MasterHandler{svc: svc}
}

// ListCompanies lists all companies
// @Summary List all companies
// @Description Fetch all companies from the database
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.Company}
// @Router /master/companies [get]
// @Security BearerAuth
func (h *MasterHandler) ListCompanies(c fiber.Ctx) error {
	companies, err := h.svc.ListCompanies(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list companies", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Companies listed", companies)
}

// CreateCompany creates a new company
// @Summary Create a company
// @Description Create a new company record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateCompanyRequest true "Company details"
// @Success 201 {object} response.APIResponse{data=repository.Company}
// @Router /master/companies [post]
// @Security BearerAuth
func (h *MasterHandler) CreateCompany(c fiber.Ctx) error {
	req := new(CreateCompanyRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	company, err := h.svc.CreateCompany(c.Context(), req.Name, req.EntityID, req.NpwpStatus, req.Location, req.Status, req.Address)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create company", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Company created", company)
}

func (h *MasterHandler) UpdateCompany(c fiber.Ctx) error {
	type request struct {
		Name       string `json:"name"`
		EntityID   string `json:"entity_id"`
		NpwpStatus string `json:"npwp_status"`
		Location   string `json:"location"`
		Status     string `json:"status"`
		Address    string `json:"address"`
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid company ID", err.Error())
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	company, err := h.svc.UpdateCompany(c.Context(), id, req.Name, req.EntityID, req.NpwpStatus, req.Location, req.Status, req.Address)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update company", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Company updated", company)
}

func (h *MasterHandler) DeleteCompany(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid company ID", err.Error())
	}

	if err := h.svc.DeleteCompany(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete company", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Company deleted", nil)
}

// Racks
// ListAllRacks lists all racks
// @Summary List all racks
// @Description Fetch all racks globally
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.ListAllRacksGlobalRow}
// @Router /master/racks [get]
// @Security BearerAuth
func (h *MasterHandler) ListAllRacks(c fiber.Ctx) error {
	racks, err := h.svc.ListAllRacksGlobal(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list racks", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Racks listed", racks)
}

// CreateRack creates a new rack
// @Summary Create a rack
// @Description Create a new rack record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateRackRequest true "Rack details"
// @Success 201 {object} response.APIResponse{data=repository.Rack}
// @Router /master/racks [post]
// @Security BearerAuth
func (h *MasterHandler) CreateRack(c fiber.Ctx) error {
	req := new(CreateRackRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	deptID, err := uuid.Parse(req.DepartmentID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid department ID", err.Error())
	}
	rack, err := h.svc.CreateRack(c.Context(), deptID, req.Name, req.LocationDetail)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create rack", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Rack created", rack)
}

// UpdateRack updates a rack
// @Summary Update a rack
// @Description Update an existing rack record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Rack ID"
// @Param request body UpdateRackRequest true "Rack updates"
// @Success 200 {object} response.APIResponse{data=repository.Rack}
// @Router /master/racks/{id} [put]
// @Security BearerAuth
func (h *MasterHandler) UpdateRack(c fiber.Ctx) error {
	type request struct {
		Name           string `json:"name"`
		LocationDetail string `json:"location_detail"`
	}
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}
	req := new(UpdateRackRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	rack, err := h.svc.UpdateRack(c.Context(), id, req.Name, req.LocationDetail)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update rack", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Rack updated", rack)
}

// DeleteRack deletes a rack
// @Summary Delete a rack
// @Description Delete a rack record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Rack ID"
// @Success 200 {object} response.APIResponse
// @Router /master/racks/{id} [delete]
// @Security BearerAuth
func (h *MasterHandler) DeleteRack(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}
	if err := h.svc.DeleteRack(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete rack", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Rack deleted", nil)
}

// Boxes
// ListAllBoxes lists all boxes
// @Summary List all boxes
// @Description Fetch all boxes globally
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.ListAllBoxesGlobalRow}
// @Router /master/boxes [get]
// @Security BearerAuth
func (h *MasterHandler) ListAllBoxes(c fiber.Ctx) error {
	boxes, err := h.svc.ListAllBoxesGlobal(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list boxes", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Boxes listed", boxes)
}

// CreateBox creates a new box
// @Summary Create a box
// @Description Create a new box record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateBoxRequest true "Box details"
// @Success 201 {object} response.APIResponse{data=repository.Box}
// @Router /master/boxes [post]
// @Security BearerAuth
func (h *MasterHandler) CreateBox(c fiber.Ctx) error {
	req := new(CreateBoxRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	rackID, err := uuid.Parse(req.RackID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}
	box, err := h.svc.CreateBox(c.Context(), rackID, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create box", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Box created", box)
}

// UpdateBox updates a box
// @Summary Update a box
// @Description Update an existing box record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Box ID"
// @Param request body UpdateBoxRequest true "Box updates"
// @Success 200 {object} response.APIResponse{data=repository.Box}
// @Router /master/boxes/{id} [put]
// @Security BearerAuth
func (h *MasterHandler) UpdateBox(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid box ID", err.Error())
	}
	req := new(UpdateBoxRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	box, err := h.svc.UpdateBox(c.Context(), id, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update box", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Box updated", box)
}

// DeleteBox deletes a box
// @Summary Delete a box
// @Description Delete a box record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Box ID"
// @Success 200 {object} response.APIResponse
// @Router /master/boxes/{id} [delete]
// @Security BearerAuth
func (h *MasterHandler) DeleteBox(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid box ID", err.Error())
	}
	if err := h.svc.DeleteBox(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete box", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Box deleted", nil)
}

// Ordners
// ListAllOrdners lists all ordners
// @Summary List all ordners
// @Description Fetch all ordners globally
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.ListAllOrdnersGlobalRow}
// @Router /master/ordners [get]
// @Security BearerAuth
func (h *MasterHandler) ListAllOrdners(c fiber.Ctx) error {
	ordners, err := h.svc.ListAllOrdnersGlobal(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list ordners", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Ordners listed", ordners)
}

// CreateOrdner creates a new ordner
// @Summary Create an ordner
// @Description Create a new ordner record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateOrdnerRequest true "Ordner details"
// @Success 201 {object} response.APIResponse{data=repository.Ordner}
// @Router /master/ordners [post]
// @Security BearerAuth
func (h *MasterHandler) CreateOrdner(c fiber.Ctx) error {
	req := new(CreateOrdnerRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	boxID, err := uuid.Parse(req.BoxID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid box ID", err.Error())
	}
	ordner, err := h.svc.CreateOrdner(c.Context(), boxID, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create ordner", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Ordner created", ordner)
}

// UpdateOrdner updates an ordner
// @Summary Update an ordner
// @Description Update an existing ordner record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Ordner ID"
// @Param request body UpdateOrdnerRequest true "Ordner updates"
// @Success 200 {object} response.APIResponse{data=repository.Ordner}
// @Router /master/ordners/{id} [put]
// @Security BearerAuth
func (h *MasterHandler) UpdateOrdner(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid ordner ID", err.Error())
	}
	req := new(UpdateOrdnerRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	ordner, err := h.svc.UpdateOrdner(c.Context(), id, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update ordner", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Ordner updated", ordner)
}

// DeleteOrdner deletes an ordner
// @Summary Delete an ordner
// @Description Delete an ordner record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Ordner ID"
// @Success 200 {object} response.APIResponse
// @Router /master/ordners/{id} [delete]
// @Security BearerAuth
func (h *MasterHandler) DeleteOrdner(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid ordner ID", err.Error())
	}
	if err := h.svc.DeleteOrdner(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete ordner", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Ordner deleted", nil)
}

func (h *MasterHandler) GetTopology(c fiber.Ctx) error {
	topology, err := h.svc.GetTopology(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get topology", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Warehouse topology", topology)
}

// ListBranches lists branches
// @Summary List branches
// @Description Fetch branches, optionally filtered by company_id
// @Tags Master
// @Accept json
// @Produce json
// @Param company_id query string false "Filter by Company ID"
// @Success 200 {object} response.APIResponse{data=[]repository.Branch}
// @Router /master/branches [get]
// @Security BearerAuth
func (h *MasterHandler) ListBranches(c fiber.Ctx) error {
	companyIDStr := c.Query("company_id")
	if companyIDStr != "" {
		companyID, err := uuid.Parse(companyIDStr)
		if err != nil {
			return response.Error(c, fiber.StatusBadRequest, "Invalid company_id", err.Error())
		}
		branches, err := h.svc.ListBranches(c.Context(), companyID)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to list branches", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Branches listed", branches)
	}

	// If no company_id, list all branches globally
	branches, err := h.svc.ListAllBranchesGlobal(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list all branches", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "All branches listed", branches)
}

// ListAllBranchesGlobal lists all branches globally
// @Summary List all branches
// @Description Fetch all branches with company names
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.ListAllBranchesGlobalRow}
// @Router /master/branches-all [get]
// @Security BearerAuth
func (h *MasterHandler) ListAllBranchesGlobal(c fiber.Ctx) error {
	branches, err := h.svc.ListAllBranchesGlobal(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list all branches", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "All branches listed", branches)
}

// CreateBranch creates a new branch
// @Summary Create a branch
// @Description Create a new branch record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateBranchRequest true "Branch details"
// @Success 201 {object} response.APIResponse{data=repository.Branch}
// @Router /master/branches [post]
// @Security BearerAuth
func (h *MasterHandler) CreateBranch(c fiber.Ctx) error {
	req := new(CreateBranchRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	companyID, err := uuid.Parse(req.CompanyID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid company ID", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	branch, err := h.svc.CreateBranch(c.Context(), companyID, req.Name, req.Location, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create branch", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Branch created", branch)
}

// UpdateBranch updates a branch
// @Summary Update a branch
// @Description Update an existing branch record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Branch ID"
// @Param request body UpdateBranchRequest true "Branch updates"
// @Success 200 {object} response.APIResponse{data=repository.Branch}
// @Router /master/branches/{id} [put]
// @Security BearerAuth
func (h *MasterHandler) UpdateBranch(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid branch ID", err.Error())
	}
	req := new(UpdateBranchRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	branch, err := h.svc.UpdateBranch(c.Context(), id, req.Name, req.Location, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update branch", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Branch updated", branch)
}

// DeleteBranch deletes a branch
// @Summary Delete a branch
// @Description Delete a branch record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Branch ID"
// @Success 200 {object} response.APIResponse
// @Router /master/branches/{id} [delete]
// @Security BearerAuth
func (h *MasterHandler) DeleteBranch(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid branch ID", err.Error())
	}

	if err := h.svc.DeleteBranch(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete branch", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Branch deleted", nil)
}

// Departments
// ListAllDepartments lists all departments
// @Summary List all departments
// @Description Fetch all departments from the database
// @Tags Master
// @Accept json
// @Produce json
// @Success 200 {object} response.APIResponse{data=[]repository.ListAllDepartmentsRow}
// @Router /master/departments [get]
// @Security BearerAuth
func (h *MasterHandler) ListAllDepartments(c fiber.Ctx) error {
	depts, err := h.svc.ListAllDepartments(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list departments", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Departments listed", depts)
}

// CreateDepartment creates a new department
// @Summary Create a department
// @Description Create a new department record
// @Tags Master
// @Accept json
// @Produce json
// @Param request body CreateDepartmentRequest true "Department details"
// @Success 201 {object} response.APIResponse{data=repository.Department}
// @Router /master/departments [post]
// @Security BearerAuth
func (h *MasterHandler) CreateDepartment(c fiber.Ctx) error {
	req := new(CreateDepartmentRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	branchID, err := uuid.Parse(req.BranchID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid branch ID", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	dept, err := h.svc.CreateDepartment(c.Context(), branchID, req.Name, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create department", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "Department created", dept)
}

// UpdateDepartment updates a department
// @Summary Update a department
// @Description Update an existing department record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Param request body UpdateDepartmentRequest true "Department updates"
// @Success 200 {object} response.APIResponse{data=repository.Department}
// @Router /master/departments/{id} [put]
// @Security BearerAuth
func (h *MasterHandler) UpdateDepartment(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid department ID", err.Error())
	}
	req := new(UpdateDepartmentRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	dept, err := h.svc.UpdateDepartment(c.Context(), id, req.Name, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update department", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Department updated", dept)
}

// DeleteDepartment deletes a department
// @Summary Delete a department
// @Description Delete a department record
// @Tags Master
// @Accept json
// @Produce json
// @Param id path string true "Department ID"
// @Success 200 {object} response.APIResponse
// @Router /master/departments/{id} [delete]
// @Security BearerAuth
func (h *MasterHandler) DeleteDepartment(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid department ID", err.Error())
	}

	if err := h.svc.DeleteDepartment(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete department", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Department deleted", nil)
}

func (h *MasterHandler) ListRoles(c fiber.Ctx) error {
	roles, err := h.svc.ListRoles(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list roles", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Roles listed", roles)
}

func (h *MasterHandler) ListRetentionPolicies(c fiber.Ctx) error {
	policies, err := h.svc.ListRetentionPolicies(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list retention policies", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Retention policies listed", policies)
}

func (h *MasterHandler) GetSettings(c fiber.Ctx) error {
	category := c.Params("category")
	settings, err := h.svc.GetSettings(c.Context(), category)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get settings", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Settings retrieved", settings)
}

func (h *MasterHandler) UpdateSetting(c fiber.Ctx) error {
	category := c.Params("category")
	req := new(UpdateSettingRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	userID := c.Locals("user_id").(uuid.UUID)
	setting, err := h.svc.UpdateSetting(c.Context(), category, req.Key, req.Value, req.Type, req.Desc, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update setting", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Setting updated", setting)
}

func (h *MasterHandler) GetIntegrationStatus(c fiber.Ctx) error {
	nodes, err := h.svc.GetIntegrationStatus(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get integration status", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Integration status retrieved", nodes)
}

func (h *MasterHandler) UpdateIntegrationNode(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid node ID", err.Error())
	}

	req := new(UpdateIntegrationNodeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	node, err := h.svc.UpdateIntegrationNode(c.Context(), id, req.Name, req.Endpoint, req.IsActive, req.IsCritical, req.Config)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update integration node", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Integration node updated", node)
}

type CreateCompanyRequest struct {
	Name       string `json:"name"`
	EntityID   string `json:"entity_id"`
	NpwpStatus string `json:"npwp_status"`
	Location   string `json:"location"`
	Status     string `json:"status"`
	Address    string `json:"address"`
}

type UpdateCompanyRequest struct {
	Name       string `json:"name"`
	EntityID   string `json:"entity_id"`
	NpwpStatus string `json:"npwp_status"`
	Location   string `json:"location"`
	Status     string `json:"status"`
	Address    string `json:"address"`
}

type CreateBranchRequest struct {
	CompanyID string  `json:"company_id"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	HeadID    *string `json:"head_id"`
}

type UpdateBranchRequest struct {
	Name     string  `json:"name"`
	Location string  `json:"location"`
	HeadID   *string `json:"head_id"`
}

type CreateDepartmentRequest struct {
	BranchID string  `json:"branch_id"`
	Name     string  `json:"name"`
	HeadID   *string `json:"head_id"`
}

type UpdateDepartmentRequest struct {
	Name   string  `json:"name"`
	HeadID *string `json:"head_id"`
}

type CreateRackRequest struct {
	DepartmentID   string `json:"department_id"`
	Name           string `json:"name"`
	LocationDetail string `json:"location_detail"`
}

type UpdateRackRequest struct {
	Name           string `json:"name"`
	LocationDetail string `json:"location_detail"`
}

type CreateBoxRequest struct {
	RackID string `json:"rack_id"`
	Name   string `json:"name"`
}

type UpdateBoxRequest struct {
	Name string `json:"name"`
}

type CreateOrdnerRequest struct {
	BoxID string `json:"box_id"`
	Name  string `json:"name"`
}

type UpdateOrdnerRequest struct {
	Name string `json:"name"`
}

type UpdateSettingRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
	Desc  string `json:"description"`
}

type UpdateIntegrationNodeRequest struct {
	Name       string          `json:"name"`
	Endpoint   string          `json:"endpoint"`
	IsActive   bool            `json:"is_active"`
	IsCritical bool            `json:"is_critical"`
	Config     json.RawMessage `json:"config_json"`
}

func (h *MasterHandler) DownloadIntegrationReport(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportIntegrationReport(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate report", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}
