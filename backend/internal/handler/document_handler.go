package handler

import (
	"fmt"
	"io"
	"log"
	"strconv"
	"strings"
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
	} else {
		// Default to user's department if not elevated role
		role := c.Locals("user_role").(string)
		if role != "superadmin" && role != "manajer" && !strings.Contains(role, "doc controller") {
			userID := c.Locals("user_id").(uuid.UUID)
			user, err := h.svc.GetUser(c.Context(), userID)
			if err == nil && user.DepartmentID.Valid {
				deptID = user.DepartmentID.Bytes
			}
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
	description := c.FormValue("description")
	if description == "" {
		description = c.FormValue("notes")
	}
	
	typeID, _ := uuid.Parse(c.FormValue("type_id"))
	companyID, _ := uuid.Parse(c.FormValue("company_id"))
	branchID, _ := uuid.Parse(c.FormValue("branch_id"))
	departmentID, _ := uuid.Parse(c.FormValue("department_id"))
	batchID, _ := uuid.Parse(c.FormValue("batch_id"))
	
	sensitivity := c.FormValue("sensitivity")
	if sensitivity == "" {
		sensitivity = c.FormValue("category") // fallback to category from UI
	}
	
	urgency := c.FormValue("urgency")
	documentDateStr := c.FormValue("document_date")
	pageCount, _ := strconv.Atoi(c.FormValue("page_count"))
	
	ownerID := c.Locals("user_id").(uuid.UUID)

	// Auto-detect Department, Branch, and Company from User
	user, err := h.svc.GetUser(c.Context(), ownerID)
	if err == nil {
		if departmentID == uuid.Nil && user.DepartmentID.Valid {
			departmentID = user.DepartmentID.Bytes
		}
		
		// If we have a department, trace back to Branch and Company
		if departmentID != uuid.Nil {
			dept, err := h.svc.GetDepartment(c.Context(), departmentID)
			if err == nil {
				if branchID == uuid.Nil {
					branchID = dept.BranchID
				}
				
				branch, err := h.svc.GetBranch(c.Context(), dept.BranchID)
				if err == nil && companyID == uuid.Nil {
					companyID = branch.CompanyID
				}
			}
		}
	}

	if companyID == uuid.Nil || branchID == uuid.Nil || departmentID == uuid.Nil {
		return response.Error(c, fiber.StatusBadRequest, "Missing required location data", "Please ensure your profile is complete or select Company, Branch, and Department manually.")
	}

	f, err := file.Open()
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
	}
	defer f.Close()

	doc, err := h.svc.UploadDocument(c.Context(), service.UploadDocumentParams{
		Title:        title,
		Description:  description,
		FileName:     file.Filename,
		FileSize:     file.Size,
		MimeType:     file.Header.Get("Content-Type"),
		Content:      f,
		OwnerID:      ownerID,
		CompanyID:    companyID,
		BranchID:     branchID,
		DepartmentID: departmentID,
		BatchID:      batchID,
		TypeID:       typeID,
		Sensitivity:  sensitivity,
		Urgency:      urgency,
		DocumentDate: documentDateStr,
		PageCount:    pageCount,
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

	userName := "Authorized User"
	
	// Try to get user from token in query if context is missing (for external links)
	token := c.Query("token")
	if token != "" {
		// In a real app, we would validate the token here if not already done by middleware
		// For now, let's assume middleware handled it or we use it to get the user
	}
	
	if val := c.Locals("user_name"); val != nil {
		userName = val.(string)
	}

	log.Printf("[DocumentHandler] Generating preview for doc: %s (User: %s)", docID, userName)
	pdfData, mimeType, err := h.svc.GetWatermarkedPDF(c.Context(), docID, userName, "center")
	if err != nil {
		log.Printf("[DocumentHandler] Error generating preview: %v", err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate preview", err.Error())
	}

	if len(pdfData) == 0 {
		log.Printf("[DocumentHandler] Warning: Preview data is empty for doc: %s", docID)
	}

	log.Printf("[DocumentHandler] Preview generated successfully for doc: %s (Type: %s, Size: %d)", docID, mimeType, len(pdfData))

	if mimeType == "" {
		mimeType = "application/pdf"
	}

	c.Set("Content-Type", mimeType)
	
	disposition := "inline"
	if c.Query("download") == "true" {
		doc, err := h.svc.GetDocument(c.Context(), docID)
		if err == nil {
			disposition = fmt.Sprintf("attachment; filename=\"%s\"", doc.FileName)
		} else {
			disposition = "attachment; filename=\"document.pdf\""
		}
	}
	c.Set("Content-Disposition", disposition)
	
	// Allow framing from any origin for previews (security is handled by the token)
	c.Response().Header.Del("X-Frame-Options")
	c.Set("Content-Security-Policy", "frame-ancestors *")
	
	return c.Send(pdfData)
}
func (h *DocumentHandler) List(c fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mine := c.Query("mine") == "true"
	
	var docs interface{}
	var err error
	
	if mine {
		userID := c.Locals("user_id").(uuid.UUID)
		docs, err = h.svc.ListRecentDocumentsByOwner(c.Context(), userID, limit)
	} else {
		docs, err = h.svc.ListRecentDocuments(c.Context(), limit)
	}
	
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list documents", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Documents retrieved successfully", docs)
}

func (h *DocumentHandler) ListOCRHistory(c fiber.Ctx) error {
	page, _ := strconv.Atoi(c.Query("page", "1"))
	pageSize, _ := strconv.Atoi(c.Query("page_size", "10"))
	
	if page < 1 { page = 1 }
	if pageSize < 1 { pageSize = 10 }

	history, total, err := h.svc.ListOCRHistory(c.Context(), page, pageSize)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list history", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "History retrieved successfully", fiber.Map{
		"history": history,
		"pagination": fiber.Map{
			"total_items": total,
			"page":        page,
			"page_size":   pageSize,
			"total_pages": (total + int64(pageSize) - 1) / int64(pageSize),
		},
	})
}
func (h *DocumentHandler) GetByID(c fiber.Ctx) error {
	idStr := c.Params("id")
	docID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	log.Printf("[DocumentHandler] Fetching document: %s", docID)
	doc, err := h.svc.GetDocument(c.Context(), docID)
	if err != nil {
		log.Printf("[DocumentHandler] Error fetching document %s: %v", docID, err)
		return response.Error(c, fiber.StatusNotFound, "Document not found", err.Error())
	}

	log.Printf("[DocumentHandler] Document %s found: %s (Mime: %s)", docID, doc.Title, doc.MimeType)
	return response.Success(c, fiber.StatusOK, "Document found", doc)
}

func (h *DocumentHandler) GetOCRData(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	job, err := h.svc.GetOCRJob(c.Context(), docID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "OCR data not found", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "OCR data retrieved", fiber.Map{
		"id":            job.EntityID,
		"filename":      job.SourceFilePath.String,
		"status":        job.Status,
		"words_json":    string(job.WordsJson),
		"ai_analysis":   string(job.AiMetadata),
		"preview_path":  job.PreviewPath.String,
		"preview_paths": string(job.PreviewPaths),
	})
}
func (h *DocumentHandler) GetLoans(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	loans, err := h.svc.GetLoanHistory(c.Context(), docID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch loan history", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Loan history retrieved", loans)
}
func (h *DocumentHandler) Approve(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	var req struct {
		Notes string `json:"notes"`
	}
	c.Bind().JSON(&req)

	if err := h.svc.ApproveDocument(c.Context(), docID, req.Notes); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to approve document", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document approved", nil)
}

func (h *DocumentHandler) Reject(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	var req struct {
		Reason string `json:"reason"`
		Notes  string `json:"notes"`
	}
	c.Bind().JSON(&req)

	if err := h.svc.RejectDocument(c.Context(), docID, req.Reason, req.Notes); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to reject document", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document rejected", nil)
}
func (h *DocumentHandler) BulkApprove(c fiber.Ctx) error {
	var req struct {
		IDs []uuid.UUID `json:"ids"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.svc.BulkApproveDocuments(c.Context(), req.IDs); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to bulk approve", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Documents approved", nil)
}

func (h *DocumentHandler) BulkReject(c fiber.Ctx) error {
	var req struct {
		IDs []uuid.UUID `json:"ids"`
	}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if err := h.svc.BulkRejectDocuments(c.Context(), req.IDs); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to bulk reject", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Documents rejected", nil)
}

func (h *DocumentHandler) Update(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	title := c.FormValue("title")
	description := c.FormValue("description")
	typeID, _ := uuid.Parse(c.FormValue("type_id"))
	sensitivity := c.FormValue("sensitivity")
	urgency := c.FormValue("urgency")
	documentDate := c.FormValue("document_date")
	
	// Handle optional file replacement
	file, _ := c.FormFile("file")
	var fileContent io.Reader
	var fileName string
	var fileSize int64
	var mimeType string

	if file != nil {
		f, err := file.Open()
		if err == nil {
			defer f.Close()
			fileContent = f
			fileName = file.Filename
			fileSize = file.Size
			mimeType = file.Header.Get("Content-Type")
		}
	}

	doc, err := h.svc.UpdateDocument(c.Context(), service.UpdateDocumentParams{
		ID:           docID,
		Title:        title,
		Description:  description,
		TypeID:       typeID,
		Sensitivity:  sensitivity,
		Urgency:      urgency,
		DocumentDate: documentDate,
		FileContent:  fileContent,
		FileName:     fileName,
		FileSize:     fileSize,
		MimeType:     mimeType,
	})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update document", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document updated successfully", doc)
}
