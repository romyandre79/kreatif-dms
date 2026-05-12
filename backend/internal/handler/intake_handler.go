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

func (h *IntakeHandler) ListPending(c fiber.Ctx) error {
	manifests, err := h.svc.ListPendingManifests(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch pending manifests", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Pending manifests retrieved", manifests)
}

func (h *IntakeHandler) Receive(c fiber.Ctx) error {
	var req struct {
		ManifestID uuid.UUID                    `json:"manifest_id"`
		Notes      string                       `json:"notes"`
		Items      []service.ReceiveRequestItem `json:"items"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	receivedBy := c.Locals("user_id").(uuid.UUID)
	err := h.svc.ReceiveDocument(c.Context(), req.ManifestID, receivedBy, req.Notes, req.Items)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to process intake", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Manifest/Documents received successfully", nil)
}

func (h *IntakeHandler) Reject(c fiber.Ctx) error {
	idStr := c.Params("id")
	manifestID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid manifest ID", err.Error())
	}

	var req struct {
		Reason string `json:"reason"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	rejectedBy := c.Locals("user_id").(uuid.UUID)
	err = h.svc.RejectManifest(c.Context(), manifestID, rejectedBy, req.Reason)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to reject manifest", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Manifest rejected successfully", nil)
}

func (h *IntakeHandler) ListStaging(c fiber.Ctx) error {
	manifests, err := h.svc.ListStagingManifests(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch staging manifests", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Staging manifests retrieved", manifests)
}

func (h *IntakeHandler) StagingStats(c fiber.Ctx) error {
	stats, err := h.svc.GetStagingStats(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch staging stats", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Staging stats retrieved", stats)
}

func (h *IntakeHandler) Index(c fiber.Ctx) error {
	var req service.IndexingRequest
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	// If ID is in params (PUT /indexing/:id), override the one in body
	idParam := c.Params("id")
	if idParam != "" {
		docID, err := uuid.Parse(idParam)
		if err == nil {
			req.DocumentID = docID
		}
	}

	err := h.svc.IndexDocument(c.Context(), req)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to index document", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document indexed successfully", nil)
}

func (h *IntakeHandler) GetLabelingDocuments(c fiber.Ctx) error {
	status := c.Query("status", "digitized")
	docs, err := h.svc.ListLabelingDocuments(c.Context(), status)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list documents", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Documents retrieved successfully", docs)
}

func (h *IntakeHandler) MarkAsLabeled(c fiber.Ctx) error {
	var req struct {
		DocumentIDs []uuid.UUID `json:"document_ids"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request", err.Error())
	}

	err := h.svc.MarkAsLabeled(c.Context(), req.DocumentIDs)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to mark documents as labeled", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Documents marked as labeled successfully", nil)
}

func (h *IntakeHandler) GetLabelingStats(c fiber.Ctx) error {
	stats, err := h.svc.GetLabelingStats(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get stats", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Stats retrieved successfully", stats)
}

