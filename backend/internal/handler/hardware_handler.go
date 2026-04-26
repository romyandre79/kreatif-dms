package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type HardwareHandler struct {
	svc *service.HardwareService
}

func NewHardwareHandler(svc *service.HardwareService) *HardwareHandler {
	return &HardwareHandler{svc: svc}
}

func (h *HardwareHandler) AssignRFID(c fiber.Ctx) error {
	type request struct {
		TagID      string    `json:"tag_id"`
		DocumentID uuid.UUID `json:"document_id"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	tag, err := h.svc.AssignRFID(c.Context(), req.TagID, req.DocumentID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to assign RFID", err.Error())
	}
	return response.Success(c, fiber.StatusCreated, "RFID assigned successfully", tag)
}

func (h *HardwareHandler) GenerateLabel(c fiber.Ctx) error {
	entityType := c.Query("type", "document")
	idStr := c.Query("id")
	entityID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid entity ID", err.Error())
	}

	label, err := h.svc.GenerateLabel(c.Context(), entityType, entityID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate label", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Label generated", label)
}

func (h *HardwareHandler) ListRFID(c fiber.Ctx) error {
	tags, err := h.svc.ListRFID(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list RFID tags", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "RFID tags retrieved", tags)
}
