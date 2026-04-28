package handler

import (
	"encoding/json"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type MasterHandler struct {
	svc *service.MasterService
}

func NewMasterHandler(svc *service.MasterService) *MasterHandler {
	return &MasterHandler{svc: svc}
}

func (h *MasterHandler) ListCompanies(c fiber.Ctx) error {
	companies, err := h.svc.ListCompanies(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list companies", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Companies listed", companies)
}

func (h *MasterHandler) CreateCompany(c fiber.Ctx) error {
	type request struct {
		Name       string `json:"name"`
		EntityID   string `json:"entity_id"`
		NpwpStatus string `json:"npwp_status"`
		Location   string `json:"location"`
		Status     string `json:"status"`
		Address    string `json:"address"`
	}
	req := new(request)
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

func (h *MasterHandler) GetTopology(c fiber.Ctx) error {
	topology, err := h.svc.GetTopology(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get topology", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Warehouse topology", topology)
}

func (h *MasterHandler) ListBranches(c fiber.Ctx) error {
	companyID, err := uuid.Parse(c.Query("company_id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid company_id", err.Error())
	}
	branches, err := h.svc.ListBranches(c.Context(), companyID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list branches", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Branches listed", branches)
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
	type request struct {
		Key       string `json:"key"`
		Value     string `json:"value"`
		Type      string `json:"type"`
		Desc      string `json:"description"`
	}
	category := c.Params("category")
	req := new(request)
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

	type request struct {
		Name       string          `json:"name"`
		Endpoint   string          `json:"endpoint"`
		IsActive   bool            `json:"is_active"`
		IsCritical bool            `json:"is_critical"`
		Config     json.RawMessage `json:"config_json"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	node, err := h.svc.UpdateIntegrationNode(c.Context(), id, req.Name, req.Endpoint, req.IsActive, req.IsCritical, req.Config)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update integration node", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Integration node updated", node)
}
