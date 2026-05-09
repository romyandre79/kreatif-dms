package handler

import (
	"encoding/json"
	"fmt"
	"net"
	"net/http"
	"strconv"
	"strings"
	"time"
	"crypto/tls"
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "company", &company.ID, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "company", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "company", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Company deleted", nil)
}

func (h *MasterHandler) ExportCompanies(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportCompanies(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export companies", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportCompanies(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportCompanies(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import companies", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d companies", count), nil)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "rack", &rack.ID, req, c.IP())

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
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}
	req := new(UpdateRackRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	deptID, err := uuid.Parse(req.DepartmentID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid department ID", err.Error())
	}
	rack, err := h.svc.UpdateRack(c.Context(), id, deptID, req.Name, req.LocationDetail)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update rack", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "rack", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "rack", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Rack deleted", nil)
}

func (h *MasterHandler) ExportRacks(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportRacks(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export racks", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportRacks(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportRacks(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import racks", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d racks", count), nil)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "box", &box.ID, req, c.IP())

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
	rackID, err := uuid.Parse(req.RackID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid rack ID", err.Error())
	}
	box, err := h.svc.UpdateBox(c.Context(), id, rackID, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update box", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "box", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "box", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Box deleted", nil)
}

func (h *MasterHandler) ExportBoxes(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportBoxes(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export boxes", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportBoxes(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportBoxes(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import boxes", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d boxes", count), nil)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "ordner", &ordner.ID, req, c.IP())

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
	boxID, err := uuid.Parse(req.BoxID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid box ID", err.Error())
	}
	ordner, err := h.svc.UpdateOrdner(c.Context(), id, boxID, req.Name)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update ordner", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "ordner", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "ordner", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Ordner deleted", nil)
}

func (h *MasterHandler) ExportOrdners(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportOrdners(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export ordners", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportOrdners(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportOrdners(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import ordners", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d ordners", count), nil)
}

func (h *MasterHandler) GetTopology(c fiber.Ctx) error {
	var deptID uuid.UUID
	
	role := c.Locals("user_role").(string)
	if role != "superadmin" && role != "manajer" && !strings.Contains(role, "doc controller") {
		userID := c.Locals("user_id").(uuid.UUID)
		user, err := h.svc.GetUser(c.Context(), userID)
		if err == nil && user.DepartmentID.Valid {
			deptID = user.DepartmentID.Bytes
		}
	}

	topology, err := h.svc.GetTopology(c.Context(), deptID)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "branch", &branch.ID, req, c.IP())

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

	companyID, err := uuid.Parse(req.CompanyID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid company ID", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	branch, err := h.svc.UpdateBranch(c.Context(), id, companyID, req.Name, req.Location, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update branch", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "branch", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "branch", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Branch deleted", nil)
}

func (h *MasterHandler) ExportBranches(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportBranches(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export branches", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportBranches(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportBranches(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import branches", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d branches", count), nil)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "department", &dept.ID, req, c.IP())

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

	branchID, err := uuid.Parse(req.BranchID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid branch ID", err.Error())
	}

	var headID *uuid.UUID
	if req.HeadID != nil && *req.HeadID != "" {
		uid, _ := uuid.Parse(*req.HeadID)
		headID = &uid
	}

	dept, err := h.svc.UpdateDepartment(c.Context(), id, branchID, req.Name, headID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update department", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "department", &id, req, c.IP())

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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "department", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Department deleted", nil)
}

func (h *MasterHandler) ExportDepartments(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportDepartments(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export departments", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportDepartments(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportDepartments(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import departments", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d departments", count), nil)
}

func (h *MasterHandler) ListRoles(c fiber.Ctx) error {
	roles, err := h.svc.ListRoles(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list roles", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Roles listed", roles)
}

func (h *MasterHandler) ListSystemModules(c fiber.Ctx) error {
	modules, err := h.svc.ListSystemModules(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list modules", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Modules listed", modules)
}

func (h *MasterHandler) GetSystemModule(c fiber.Ctx) error {
	id := c.Params("id")
	module, err := h.svc.GetSystemModule(c.Context(), id)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Module not found", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Module retrieved", module)
}

func (h *MasterHandler) CreateSystemModule(c fiber.Ctx) error {
	type request struct {
		ID             string   `json:"id"`
		Name           string   `json:"name"`
		Category       string   `json:"category"`
		Path           string   `json:"path"`
		Icon           string   `json:"icon"`
		AllowedActions []string `json:"allowed_actions"`
		SortOrder      int32    `json:"sort_order"`
		ParentID       *string  `json:"parent_id"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	module, err := h.svc.CreateSystemModule(c.Context(), req.ID, req.Name, req.Category, req.Path, req.Icon, req.AllowedActions, req.SortOrder, req.ParentID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create module", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "system_module", nil, req, c.IP())

	return response.Success(c, fiber.StatusCreated, "Module created", module)
}

func (h *MasterHandler) UpdateSystemModule(c fiber.Ctx) error {
	type request struct {
		Name           string   `json:"name"`
		Category       string   `json:"category"`
		Path           string   `json:"path"`
		Icon           string   `json:"icon"`
		AllowedActions []string `json:"allowed_actions"`
		SortOrder      int32    `json:"sort_order"`
		ParentID       *string  `json:"parent_id"`
	}
	id := c.Params("id")
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	module, err := h.svc.UpdateSystemModule(c.Context(), id, req.Name, req.Category, req.Path, req.Icon, req.AllowedActions, req.SortOrder, req.ParentID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update module", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "system_module", nil, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Module updated", module)
}

func (h *MasterHandler) DeleteSystemModule(c fiber.Ctx) error {
	id := c.Params("id")
	if err := h.svc.DeleteSystemModule(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete module", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "system_module", nil, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Module deleted", nil)
}

func (h *MasterHandler) GetRolePermissions(c fiber.Ctx) error {
	roleID, err := strconv.Atoi(c.Params("role_id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}

	permissions, err := h.svc.GetRolePermissions(c.Context(), int32(roleID))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get role permissions", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Role permissions retrieved", permissions)
}

func (h *MasterHandler) UpdateRolePermissions(c fiber.Ctx) error {
	roleID, err := strconv.Atoi(c.Params("role_id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid role ID", err.Error())
	}

	var req []struct {
		ModuleID string `json:"module_id"`
		Action   string `json:"action"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err = h.svc.UpdateRolePermissions(c.Context(), int32(roleID), req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update role permissions", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Role permissions updated", nil)
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

// Announcements Management
func (h *MasterHandler) ListAnnouncements(c fiber.Ctx) error {
	announcements, err := h.svc.ListAnnouncements(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list announcements", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Announcements listed", announcements)
}

func (h *MasterHandler) CreateAnnouncement(c fiber.Ctx) error {
	var req struct {
		Title    string `json:"title"`
		Message  string `json:"message"`
		Notes    string `json:"notes"`
		IsActive bool   `json:"is_active"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	userID := c.Locals("user_id").(uuid.UUID)
	ann, err := h.svc.CreateAnnouncement(c.Context(), req.Title, req.Message, req.Notes, req.IsActive, userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create announcement", err.Error())
	}

	// Log Activity
	h.svc.LogActivity(c.Context(), userID, "CREATE", "announcement", &ann.ID, req, c.IP())

	return response.Success(c, fiber.StatusCreated, "Announcement created", ann)
}

func (h *MasterHandler) UpdateAnnouncement(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid announcement ID", err.Error())
	}

	var req struct {
		Title    string `json:"title"`
		Message  string `json:"message"`
		Notes    string `json:"notes"`
		IsActive bool   `json:"is_active"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	ann, err := h.svc.UpdateAnnouncement(c.Context(), id, req.Title, req.Message, req.Notes, req.IsActive)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update announcement", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "announcement", &id, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Announcement updated", ann)
}

func (h *MasterHandler) DeleteAnnouncement(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid announcement ID", err.Error())
	}

	if err := h.svc.DeleteAnnouncement(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete announcement", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "announcement", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Announcement deleted", nil)
}

func (h *MasterHandler) GetIntegrationStatus(c fiber.Ctx) error {
	nodes, err := h.svc.GetIntegrationStatus(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get integration status", err.Error())
	}

	syncCount, _ := h.svc.GetSyncLogsCount(c.Context())

	return response.Success(c, fiber.StatusOK, "Integration status retrieved", fiber.Map{
		"nodes":      nodes,
		"sync_count": syncCount,
	})
}

func (h *MasterHandler) CreateIntegrationNode(c fiber.Ctx) error {
	req := new(CreateIntegrationNodeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	node, err := h.svc.CreateIntegrationNode(c.Context(), req.Name, req.ServiceType, req.Driver, req.Endpoint, req.IsCritical, req.Config)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create integration node", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "integration_node", &node.ID, req, c.IP())

	return response.Success(c, fiber.StatusCreated, "Integration node created", node)
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

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "integration_node", &id, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Integration node updated", node)
}

func (h *MasterHandler) TestIntegrationNode(c fiber.Ctx) error {
	req := new(TestIntegrationNodeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if req.ServiceType == "LDAP" {
		result, err := h.svc.TestLDAPConnection(c.Context(), req.Endpoint, req.Config)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "LDAP Connection Failed", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", result)
	}

	if req.ServiceType == "SEARCH" {
		result, err := h.svc.TestSearchConnection(c.Context(), req.Endpoint, req.Config)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Search Engine Connection Failed", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", result)
	}

	if req.ServiceType == "S3" || req.ServiceType == "STORAGE" {
		err := h.svc.TestStorageConnection(c.Context(), req.Endpoint, req.Config)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Storage Connection Failed", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", nil)
	}

	if req.ServiceType == "WHATSAPP" {
		result, err := h.svc.TestWhatsAppConnection(c.Context(), req.Endpoint, req.Config)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "WhatsApp Gateway Connection Failed", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", result)
	}

	if req.ServiceType == "SMTP" {
		err := h.svc.TestSMTPConnection(c.Context(), req.Endpoint, req.Config)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "SMTP Connection Failed", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", nil)
	}

	if req.ServiceType == "SCANNER_LOCAL" || req.ServiceType == "SCANNER_NETWORK" {
		var cfg struct {
			ConnectionString string `json:"connection_string"`
		}
		json.Unmarshal(req.Config, &cfg)
		
		target := cfg.ConnectionString
		if target == "" {
			target = req.Endpoint
		}
		
		if strings.HasPrefix(target, "http") {
			if !strings.HasSuffix(target, "/") {
				target += "/"
			}
			// Use a simple HTTP check
			tr := &http.Transport{
				TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			}
			client := &http.Client{
				Timeout:   5 * time.Second,
				Transport: tr,
			}
			req, err := http.NewRequest("GET", target+"health", nil)
			if err != nil {
				return response.Error(c, fiber.StatusInternalServerError, "Request Creation Failed", err.Error())
			}

			// Bypass ngrok warning
			req.Header.Set("ngrok-skip-browser-warning", "true")

			resp, err := client.Do(req)
			if err != nil {
				fmt.Printf("[DEBUG] Scanner Bridge Unreachable at %s: %v\n", target+"health", err)
				return response.Error(c, fiber.StatusInternalServerError, "Scanner Bridge Unreachable", err.Error())
			}
			defer resp.Body.Close()
			if resp.StatusCode >= 400 {
				return response.Error(c, fiber.StatusInternalServerError, "Scanner Bridge Error", fmt.Sprintf("Status: %d", resp.StatusCode))
			}
		} else {
			// TCP check
			conn, err := net.DialTimeout("tcp", target, 5*time.Second)
			if err != nil {
				return response.Error(c, fiber.StatusInternalServerError, "Scanner Node Unreachable", err.Error())
			}
			conn.Close()
		}
		return response.Success(c, fiber.StatusOK, "Connection Successful", nil)
	}

	return response.Error(c, fiber.StatusBadRequest, "Service type not supported for testing", "")
}

func (h *MasterHandler) DeleteIntegrationNode(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid node ID", err.Error())
	}

	if err := h.svc.DeleteIntegrationNode(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete integration node", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "integration_node", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Integration node deleted", nil)
}

type TestIntegrationNodeRequest struct {
	ServiceType string          `json:"service_type"`
	Endpoint    string          `json:"endpoint"`
	Config      json.RawMessage `json:"config_json"`
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
	CompanyID string  `json:"company_id"`
	Name      string  `json:"name"`
	Location  string  `json:"location"`
	HeadID    *string `json:"head_id"`
}

type CreateDepartmentRequest struct {
	BranchID string  `json:"branch_id"`
	Name     string  `json:"name"`
	HeadID   *string `json:"head_id"`
}

type UpdateDepartmentRequest struct {
	BranchID string  `json:"branch_id"`
	Name     string  `json:"name"`
	HeadID   *string `json:"head_id"`
}

type CreateRackRequest struct {
	DepartmentID   string `json:"department_id"`
	Name           string `json:"name"`
	LocationDetail string `json:"location_detail"`
}

type UpdateRackRequest struct {
	DepartmentID   string `json:"department_id"`
	Name           string `json:"name"`
	LocationDetail string `json:"location_detail"`
}

type CreateBoxRequest struct {
	RackID string `json:"rack_id"`
	Name   string `json:"name"`
}

type UpdateBoxRequest struct {
	RackID string `json:"rack_id"`
	Name   string `json:"name"`
}

type CreateOrdnerRequest struct {
	BoxID string `json:"box_id"`
	Name  string `json:"name"`
}

type UpdateOrdnerRequest struct {
	BoxID string `json:"box_id"`
	Name  string `json:"name"`
}

type UpdateSettingRequest struct {
	Key   string `json:"key"`
	Value string `json:"value"`
	Type  string `json:"type"`
	Desc  string `json:"description"`
}

type CreateIntegrationNodeRequest struct {
	Name        string          `json:"name"`
	ServiceType string          `json:"service_type"`
	Driver      string          `json:"driver"`
	Endpoint    string          `json:"endpoint"`
	IsCritical  bool            `json:"is_critical"`
	Config      json.RawMessage `json:"config_json"`
}

type UpdateIntegrationNodeRequest struct {
	Name       string          `json:"name"`
	Endpoint   string          `json:"endpoint"`
	IsActive   bool            `json:"is_active"`
	IsCritical bool            `json:"is_critical"`
	Config     json.RawMessage `json:"config_json"`
}

// Document Types
func (h *MasterHandler) ListDocumentTypes(c fiber.Ctx) error {
	types, err := h.svc.ListDocumentTypes(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list document types", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Document types listed", types)
}

func (h *MasterHandler) CreateDocumentType(c fiber.Ctx) error {
	req := new(CreateDocumentTypeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	docType, err := h.svc.CreateDocumentType(c.Context(), req.Code, req.Name, req.Description)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create document type", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "CREATE", "document_type", &docType.ID, req, c.IP())

	return response.Success(c, fiber.StatusCreated, "Document type created", docType)
}

func (h *MasterHandler) UpdateDocumentType(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document type ID", err.Error())
	}
	req := new(UpdateDocumentTypeRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	docType, err := h.svc.UpdateDocumentType(c.Context(), id, req.Code, req.Name, req.Description)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update document type", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "document_type", &id, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Document type updated", docType)
}

func (h *MasterHandler) DeleteDocumentType(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document type ID", err.Error())
	}

	if err := h.svc.DeleteDocumentType(c.Context(), id); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete document type", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "DELETE", "document_type", &id, nil, c.IP())

	return response.Success(c, fiber.StatusOK, "Document type deleted", nil)
}

func (h *MasterHandler) ExportDocumentTypes(c fiber.Ctx) error {
	data, fileName, err := h.svc.ExportDocumentTypes(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to export document types", err.Error())
	}

	c.Set("Content-Type", "text/csv")
	c.Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", fileName))
	return c.Send(data)
}

func (h *MasterHandler) ImportDocumentTypes(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Failed to get file from request", err.Error())
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	count, err := h.svc.ImportDocumentTypes(c.Context(), f)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to import document types", err.Error())
	}

	return response.Success(c, fiber.StatusOK, fmt.Sprintf("Successfully imported %d document types", count), nil)
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

func (h *MasterHandler) ListActivityLogs(c fiber.Ctx) error {
	limit := int32(50)
	offset := int32(0)
	
	// Optional: Parse limit/offset from query params
	
	logs, err := h.svc.ListActivityLogs(c.Context(), limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list activity logs", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Activity logs listed", logs)
}

func (h *MasterHandler) ListSsoSyncLogs(c fiber.Ctx) error {
	limit := int32(50)
	offset := int32(0)
	
	logs, err := h.svc.ListSsoSyncLogs(c.Context(), limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list sync logs", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Sync logs listed", logs)
}


type CreateDocumentTypeRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type UpdateDocumentTypeRequest struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
func (h *MasterHandler) FetchAIModels(c fiber.Ctx) error {
	driver := c.Query("driver")
	apiKey := c.Query("api_key")

	if driver == "" || apiKey == "" {
		return response.Error(c, fiber.StatusBadRequest, "Driver and API Key are required", "")
	}

	models, err := h.svc.FetchAIModels(c.Context(), driver, apiKey)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch AI models", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "AI models fetched", models)
}
func (h *MasterHandler) GetWatermarkSettings(c fiber.Ctx) error {
	settings, err := h.svc.GetWatermarkSettings(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get watermark settings", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Watermark settings retrieved", settings)
}

func (h *MasterHandler) UpdateWatermarkSettings(c fiber.Ctx) error {
	var req map[string]interface{}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.svc.UpdateWatermarkSettings(c.Context(), req); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update watermark settings", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "UPDATE", "watermark_settings", nil, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Watermark settings updated", nil)
}

func (h *MasterHandler) RegisterScanner(c fiber.Ctx) error {
	var req struct {
		Name        string                 `json:"name"`
		ServiceType string                 `json:"service_type"` // SCANNER_LOCAL or SCANNER_NETWORK
		Endpoint    string                 `json:"endpoint"`
		Config      map[string]interface{} `json:"config"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if req.ServiceType == "" {
		req.ServiceType = "SCANNER_LOCAL"
	}

	node, err := h.svc.RegisterScanner(c.Context(), req.Name, req.ServiceType, req.Endpoint, req.Config)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to register scanner", err.Error())
	}

	// Log Activity
	userID := c.Locals("user_id").(uuid.UUID)
	h.svc.LogActivity(c.Context(), userID, "REGISTER", "scanner", &node.ID, req, c.IP())

	return response.Success(c, fiber.StatusOK, "Scanner registered successfully", node)
}
