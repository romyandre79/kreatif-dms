package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type IntakeHandler struct {
	svc *service.IntakeService
}

func NewIntakeHandler(svc *service.IntakeService) *IntakeHandler {
	return &IntakeHandler{svc: svc}
}

func (h *IntakeHandler) GetManifest(c fiber.Ctx) error {
	shortID := c.Params("id")
	if shortID == "" {
		return response.Error(c, fiber.StatusBadRequest, "Manifest ID is required", "")
	}

	detail, err := h.svc.GetDocumentByManifestID(c.Context(), shortID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, err.Error(), "")
	}

	return response.Success(c, fiber.StatusOK, "Manifest found", detail)
}

func (h *IntakeHandler) GetStats(c fiber.Ctx) error {
	stats, err := h.svc.GetStats(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch stats", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Intake stats retrieved", stats)
}

func (h *IntakeHandler) Receive(c fiber.Ctx) error {
	var req struct {
		DocumentID uuid.UUID `json:"document_id"`
		Notes      string    `json:"notes"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	receivedBy := c.Locals("user_id").(uuid.UUID)
	err := h.svc.ReceiveDocument(c.Context(), req.DocumentID, receivedBy, req.Notes)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to process intake", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document received successfully", nil)
}
