package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type DocumentHandler struct {
	svc    *service.DocumentService
	search *infra.SearchService
}

func NewDocumentHandler(svc *service.DocumentService, search *infra.SearchService) *DocumentHandler {
	return &DocumentHandler{svc: svc, search: search}
}

func (h *DocumentHandler) Search(c fiber.Ctx) error {
	query := c.Query("q")
	if query == "" {
		return response.Error(c, fiber.StatusBadRequest, "Query is required", "")
	}

	deptIDStr := c.Query("department_id")
	var deptID uuid.UUID
	if deptIDStr != "" {
		id, err := uuid.Parse(deptIDStr)
		if err == nil {
			deptID = id
		}
	}

	results, err := h.search.Search(c.Context(), query, deptID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to search documents", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Search results", results)
}

func (h *DocumentHandler) Upload(c fiber.Ctx) error {
	file, err := c.FormFile("file")
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "File is required", err.Error())
	}

	// Read form values
	title := c.FormValue("title")
	companyID, _ := uuid.Parse(c.FormValue("company_id"))
	branchID, _ := uuid.Parse(c.FormValue("branch_id"))
	departmentID, _ := uuid.Parse(c.FormValue("department_id"))
	batchID, _ := uuid.Parse(c.FormValue("batch_id"))
	
	ownerID := c.Locals("user_id").(uuid.UUID)

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	doc, err := h.svc.UploadDocument(c.Context(), service.UploadDocumentParams{
		Title:        title,
		FileName:     file.Filename,
		FileSize:     file.Size,
		MimeType:     file.Header.Get("Content-Type"),
		Content:      f,
		OwnerID:      ownerID,
		CompanyID:    companyID,
		BranchID:     branchID,
		DepartmentID: departmentID,
		BatchID:      batchID,
	})

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to upload document", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Document uploaded successfully", doc)
}

func (h *DocumentHandler) Preview(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	userName := "Authorized User" // Should be taken from context if available
	
	pdfData, err := h.svc.GetWatermarkedPDF(c.Context(), docID, userName, "center")
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate preview", err.Error())
	}

	c.Set("Content-Type", "application/pdf")
	c.Set("Content-Disposition", "inline")
	return c.Send(pdfData)
}
