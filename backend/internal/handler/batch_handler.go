package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/service"
)

type BatchHandler struct {
	docSvc *service.DocumentService
}

func NewBatchHandler(docSvc *service.DocumentService) *BatchHandler {
	return &BatchHandler{docSvc: docSvc}
}

// Create starts a new processing batch
// @Summary Create Batch
// @Description Start a new batch for multiple uploads
// @Tags Batches
// @Param total_files body int true "Total files in batch"
// @Success 201 {object} repository.ProcessingBatch
// @Router /batches [post]
func (h *BatchHandler) Create(c fiber.Ctx) error {
	var req struct {
		TotalFiles int `json:"total_files"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid request body"})
	}

	userID := c.Locals("user_id").(uuid.UUID)
	batch, err := h.docSvc.CreateBatch(c.Context(), userID, req.TotalFiles)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(201).JSON(batch)
}

// GetStatus returns the details of a processing batch
// @Summary Get Batch Status
// @Description Get progress and individual file status for a batch
// @Tags Batches
// @Param id path string true "Batch ID"
// @Success 200 {object} service.BatchDetails
// @Router /batches/{id} [get]
func (h *BatchHandler) GetStatus(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "invalid batch id"})
	}

	details, err := h.docSvc.GetBatchDetails(c.Context(), id)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(details)
}
