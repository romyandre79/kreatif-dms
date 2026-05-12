package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
)

type EmailTemplateHandler struct {
	repo repository.Querier
}

func NewEmailTemplateHandler(repo repository.Querier) *EmailTemplateHandler {
	return &EmailTemplateHandler{repo: repo}
}

func (h *EmailTemplateHandler) ListTemplates(c fiber.Ctx) error {
	templates, err := h.repo.ListEmailTemplates(c.Context())
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(templates)
}

func (h *EmailTemplateHandler) GetTemplate(c fiber.Ctx) error {
	slug := c.Params("slug")
	template, err := h.repo.GetEmailTemplateBySlug(c.Context(), slug)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"error": "Template not found"})
	}
	return c.JSON(template)
}

func (h *EmailTemplateHandler) UpdateTemplate(c fiber.Ctx) error {
	idStr := c.Params("id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid ID"})
	}

	var req struct {
		Subject  string `json:"subject"`
		BodyHtml string `json:"body_html"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": err.Error()})
	}

	template, err := h.repo.UpdateEmailTemplate(c.Context(), repository.UpdateEmailTemplateParams{
		ID:       id,
		Subject:  req.Subject,
		BodyHtml: req.BodyHtml,
	})
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(template)
}
