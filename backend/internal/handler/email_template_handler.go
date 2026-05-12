package handler

import (
	"fmt"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
)

type EmailTemplateHandler struct {
	repo     repository.Querier
	emailSvc *infra.EmailService
}

func NewEmailTemplateHandler(repo repository.Querier, emailSvc *infra.EmailService) *EmailTemplateHandler {
	return &EmailTemplateHandler{repo: repo, emailSvc: emailSvc}
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

func (h *EmailTemplateHandler) TestTemplate(c fiber.Ctx) error {
	slug := c.Params("slug")
	
	var req struct {
		Email string `json:"email"`
	}
	if err := c.Bind().Body(&req); err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Format request tidak valid"})
	}

	if req.Email == "" {
		return c.Status(400).JSON(fiber.Map{"error": "Email tujuan harus diisi"})
	}

	// Mock data for testing
	mockData := map[string]string{
		"fullName":       "Test Approver",
		"docTitle":       "Laporan Keuangan Q1.pdf",
		"departmentName": "Finance & Accounting",
		"ownerName":      "Romy Andre",
		"approveLink":    "http://localhost:3000/approvals/test-uuid",
		"notes":          "Ini adalah pesan notifikasi percobaan.",
		"manifestNo":     "TEST-20240512-001",
		"itemCount":      "5",
		"reason":         "Alasan percobaan untuk persetujuan/penolakan.",
		"resetLink":      "http://localhost:3000/reset-password?token=test-token",
	}

	err := h.emailSvc.SendTemplatedEmail(c.Context(), slug, req.Email, mockData)
	if err != nil {
		return c.Status(500).JSON(fiber.Map{"error": fmt.Sprintf("Gagal mengirim email: %v", err)})
	}

	return c.JSON(fiber.Map{"message": "Email tes berhasil dikirim ke " + req.Email})
}
