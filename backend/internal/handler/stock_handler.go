package handler

import (
	"strconv"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/utils"
	"github.com/kreatif/dms-backend/pkg/response"
)

type StockHandler struct {
	svc *service.StockService
}

func NewStockHandler(svc *service.StockService) *StockHandler {
	return &StockHandler{svc: svc}
}

// GetStats returns dashboard stats for stock missions
func (h *StockHandler) GetStats(c fiber.Ctx) error {
	stats, err := h.svc.GetStats(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get stats", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Stats retrieved", stats)
}

// GetMissions returns list of audit missions
func (h *StockHandler) GetMissions(c fiber.Ctx) error {
	tab := c.Query("tab", "all")
	page, _ := strconv.Atoi(c.Query("page", "1"))
	limit, _ := strconv.Atoi(c.Query("limit", "10"))

	missions, total, err := h.svc.GetMissions(c.Context(), tab, page, limit)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get missions", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Missions retrieved", fiber.Map{
		"items": missions,
		"total": total,
	})
}

// CreateMission creates a new manual audit mission
func (h *StockHandler) CreateMission(c fiber.Ctx) error {
	type request struct {
		Title       string `json:"title" validate:"required"`
		Description string `json:"description"`
		AssignedTo  string `json:"assigned_to" validate:"required"`
		TargetArea  string `json:"target_area" validate:"required"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	mission, err := h.svc.CreateMission(c.Context(), req.Title, req.Description, req.AssignedTo, req.TargetArea)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create mission", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Mission created successfully", mission)
}

// StartMission starts the audit mission
func (h *StockHandler) StartMission(c fiber.Ctx) error {
	idStr := c.Params("id")
	missionID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid mission ID", err.Error())
	}

	err = h.svc.StartMission(c.Context(), missionID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to start mission", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Mission started successfully", nil)
}

// GetSessionItems returns scanned & expected items for a session
func (h *StockHandler) GetSessionItems(c fiber.Ctx) error {
	idStr := c.Params("id")
	sessionID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid session ID", err.Error())
	}

	items, err := h.svc.GetSessionItems(c.Context(), sessionID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get session items", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Session items retrieved", items)
}

// ScanItem submits a physical scan result
func (h *StockHandler) ScanItem(c fiber.Ctx) error {
	type request struct {
		SessionID   string `json:"session_id" validate:"required"`
		SkuCode     string `json:"sku_code" validate:"required"`
		CurrentRack string `json:"current_rack"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	sessID, err := uuid.Parse(req.SessionID)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid session ID", err.Error())
	}

	result, err := h.svc.ScanItem(c.Context(), sessID, req.SkuCode, req.CurrentRack)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to scan item", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Item scanned", result)
}

// ResolveItem handles individual discrepancy resolution
func (h *StockHandler) ResolveItem(c fiber.Ctx) error {
	idStr := c.Params("id")
	itemID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid item ID", err.Error())
	}

	type request struct {
		Resolution string `json:"resolution" validate:"required"`
		Note       string `json:"note"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err = h.svc.ResolveItem(c.Context(), itemID, req.Resolution, req.Note)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to resolve item", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Item resolved successfully", nil)
}

// ApproveSession handles final supervisor/head sign-off
func (h *StockHandler) ApproveSession(c fiber.Ctx) error {
	idStr := c.Params("id")
	sessionID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid session ID", err.Error())
	}

	type request struct {
		PIN            string `json:"pin" validate:"required"`
		ResolutionNote string `json:"resolution_note"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Get authenticated user ID
	userID := c.Locals("user_id").(uuid.UUID)

	err = h.svc.ApproveSession(c.Context(), sessionID, userID, req.PIN, req.ResolutionNote)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Approval failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Session approved and inventory updated", nil)
}
