package handler

import (
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"strconv"
	"strings"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/infra"
	"github.com/kreatif/dms-backend/internal/repository"
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
	status := c.FormValue("status")

	// Parse multipart form to support multiple files
	form, err := c.MultipartForm()
	if err != nil && status != "draft" {
		return response.Error(c, fiber.StatusBadRequest, "Invalid multipart form", err.Error())
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
		sensitivity = c.FormValue("category")
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

	// Collect all uploaded files — supports field name "files" (multi) and "file" (single, backward compat)
	var allFiles []*multipart.FileHeader
	if form != nil {
		if mf, ok := form.File["files"]; ok {
			allFiles = append(allFiles, mf...)
		}
		if mf, ok := form.File["file"]; ok && len(allFiles) == 0 {
			allFiles = append(allFiles, mf...)
		}
	}

	if len(allFiles) == 0 && status != "draft" {
		return response.Error(c, fiber.StatusBadRequest, "At least one file is required", "")
	}

	// Primary file = first in list
	var fileContent io.Reader
	var fileName, mimeType string
	var fileSize int64

	if len(allFiles) > 0 {
		primary := allFiles[0]
		f, err := primary.Open()
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to open file", err.Error())
		}
		defer f.Close()
		fileContent = f
		fileName = primary.Filename
		fileSize = primary.Size
		mimeType = primary.Header.Get("Content-Type")
	}

	// Extra files = rest of the list (index 1+)
	var extraFiles []service.ExtraFileParam
	for i := 1; i < len(allFiles); i++ {
		fh := allFiles[i]
		f, err := fh.Open()
		if err != nil {
			log.Printf("[DocumentHandler] Warning: failed to open extra file %s: %v", fh.Filename, err)
			continue
		}
		defer f.Close()
		extraFiles = append(extraFiles, service.ExtraFileParam{
			FileName: fh.Filename,
			FileSize: fh.Size,
			MimeType: fh.Header.Get("Content-Type"),
			Content:  f,
		})
	}

	doc, err := h.svc.UploadDocument(c.Context(), service.UploadDocumentParams{
		Title:        title,
		Description:  description,
		FileName:     fileName,
		FileSize:     fileSize,
		MimeType:     mimeType,
		Content:      fileContent,
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
		Status:       status,
		ExtraFiles:   extraFiles,
	})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to upload document", err.Error())
	}

	// Fetch the saved extra files to include in response
	docFiles, _ := h.svc.GetDocumentFiles(c.Context(), doc.ID)

	return response.Success(c, fiber.StatusCreated, "Document uploaded successfully", fiber.Map{
		"document": doc,
		"files":    formatDocumentFiles(doc.ID, doc.FileName, doc.FileSize, doc.MimeType, docFiles),
	})
}

func (h *DocumentHandler) ListFiles(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	doc, err := h.svc.GetDocumentBasic(c.Context(), docID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Document not found", err.Error())
	}

	extraFiles, err := h.svc.GetDocumentFiles(c.Context(), docID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get document files", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document files retrieved", formatDocumentFiles(doc.ID, doc.FileName, doc.FileSize, doc.MimeType, extraFiles))
}

func formatDocumentFiles(docID uuid.UUID, primaryFileName pgtype.Text, primaryFileSize pgtype.Int8, primaryMimeType pgtype.Text, extras []repository.DocumentFile) []fiber.Map {
	var result []fiber.Map

	// Primary file from documents table
	if primaryFileName.Valid && primaryFileName.String != "" {
		result = append(result, fiber.Map{
			"id":         docID,
			"file_name":  primaryFileName.String,
			"file_size":  primaryFileSize.Int64,
			"mime_type":  primaryMimeType.String,
			"is_primary": true,
		})
	}

	// Extra files from document_files table
	for _, f := range extras {
		result = append(result, fiber.Map{
			"id":         f.ID,
			"file_name":  f.FileName,
			"file_size":  f.FileSize,
			"mime_type":  f.MimeType,
			"is_primary": false,
		})
	}

	return result
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
			disposition = fmt.Sprintf("attachment; filename=\"%s\"", doc.FileName.String)
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
// PreviewFile serves a watermarked preview of an extra file (from document_files table).
// @Summary Preview an extra file attached to a document
// @Tags Documents
// @Produce application/pdf
// @Param id path string true "Document ID"
// @Param fileId path string true "File ID"
// @Success 200
// @Router /documents/{id}/files/{fileId}/preview [get]
// @Security BearerAuth
func (h *DocumentHandler) PreviewFile(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}
	fileID, err := uuid.Parse(c.Params("fileId"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid file ID", err.Error())
	}

	userName := "Authorized User"
	if val := c.Locals("user_name"); val != nil {
		userName = val.(string)
	}

	data, mimeType, err := h.svc.GetWatermarkedFilePreview(c.Context(), docID, fileID, userName)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate preview", err.Error())
	}

	if mimeType == "" {
		mimeType = "application/pdf"
	}

	c.Set("Content-Type", mimeType)
	c.Set("Content-Disposition", "inline")
	c.Response().Header.Del("X-Frame-Options")
	c.Set("Content-Security-Policy", "frame-ancestors *")
	return c.Send(data)
}

func (h *DocumentHandler) List(c fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "10"))
	mine := c.Query("mine") == "true"
	
	// Filtering params
	filters := repository.SearchDocumentsParams{
		Limit:  int32(limit),
		Offset: 0, // for now
	}

	if val, err := uuid.Parse(c.Query("company_id")); err == nil { filters.CompanyID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("branch_id")); err == nil { filters.BranchID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("department_id")); err == nil { filters.DepartmentID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("rack_id")); err == nil { filters.RackID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("box_id")); err == nil { filters.BoxID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("ordner_id")); err == nil { filters.OrdnerID = pgtype.UUID{Bytes: val, Valid: true} }
	if val, err := uuid.Parse(c.Query("type_id")); err == nil { filters.TypeID = pgtype.UUID{Bytes: val, Valid: true} }
	if val := c.Query("year"); val != "" {
		if yearInt, err := strconv.Atoi(val); err == nil {
			filters.Year = pgtype.Int4{Int32: int32(yearInt), Valid: true}
		}
	}

	var docs interface{}
	var err error
	
	if mine {
		userID := c.Locals("user_id").(uuid.UUID)
		docs, err = h.svc.ListRecentDocumentsByOwner(c.Context(), userID, limit)
	} else if filters.CompanyID.Valid || filters.BranchID.Valid || filters.DepartmentID.Valid || filters.RackID.Valid || filters.BoxID.Valid || filters.OrdnerID.Valid || filters.TypeID.Valid || filters.Year.Valid {
		docs, err = h.svc.FilterDocuments(c.Context(), filters)
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

	log.Printf("[DocumentHandler] Document %s found: %s (Mime: %s)", docID, doc.Title, doc.MimeType.String)
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
		Status:       c.FormValue("status"),
	})
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update document", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Document updated successfully", doc)
}

func (h *DocumentHandler) GetImage(c fiber.Ctx) error {
	docID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid document ID", err.Error())
	}

	// 1. Try to get OCR preview path first
	job, err := h.svc.GetOCRJob(c.Context(), docID)
	path := ""
	if err == nil && job.PreviewPath.Valid {
		path = job.PreviewPath.String
	} else {
		// 2. Fallback to original file if it's an image
		doc, err := h.svc.GetDocument(c.Context(), docID)
		if err == nil {
			mime := strings.ToLower(doc.MimeType.String)
			if strings.HasPrefix(mime, "image/") {
				path = doc.FilePath.String
			}
		}
	}

	if path == "" {
		// No image preview available, return 404 or a placeholder if you prefer
		return response.Error(c, fiber.StatusNotFound, "No image preview available for this document", "")
	}

	data, mime, err := h.svc.GetRawFile(c.Context(), path)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to retrieve image", err.Error())
	}

	c.Set("Content-Type", mime)
	return c.Send(data)
}

func (h *DocumentHandler) GetExplorerTree(c fiber.Ctx) error {
	tree, err := h.svc.GetExplorerTree(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to build explorer tree", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Explorer tree retrieved", tree)
}

func (h *DocumentHandler) GetExplorerConfig(c fiber.Ctx) error {
	settings, err := h.svc.GetSystemSettingsByCategory(c.Context(), "explorer")
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get explorer config", err.Error())
	}
	
	config := make(map[string]interface{})
	for _, s := range settings {
		if s.ValueType.String == "boolean" {
			config[s.Key] = s.Value.String == "true"
		} else {
			config[s.Key] = s.Value.String
		}
	}
	
	return response.Success(c, fiber.StatusOK, "Explorer config retrieved", config)
}

func (h *DocumentHandler) UpdateExplorerConfig(c fiber.Ctx) error {
	var req map[string]interface{}
	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}
	
	userID := c.Locals("user_id").(uuid.UUID)
	
	for k, v := range req {
		valStr := fmt.Sprintf("%v", v)
		valType := "string"
		if _, ok := v.(bool); ok {
			valType = "boolean"
		}
		
		_, err := h.svc.UpsertSystemSetting(c.Context(), repository.UpsertSystemSettingParams{
			Category:    "explorer",
			Key:         k,
			Value:       pgtype.Text{String: valStr, Valid: true},
			ValueType:   pgtype.Text{String: valType, Valid: true},
			UpdatedBy:   pgtype.UUID{Bytes: userID, Valid: true},
		})
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update explorer config", err.Error())
		}
	}
	
	return response.Success(c, fiber.StatusOK, "Explorer config updated", nil)
}

// GetMySubmissions returns all documents submitted by the authenticated user with approval status.
// @Summary Get my submissions
// @Tags Documents
// @Produce json
// @Param status query string false "Filter by status"
// @Param q query string false "Search by title"
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 20)"
// @Success 200 {object} response.APIResponse
// @Router /documents/my-submissions [get]
// @Security BearerAuth
func (h *DocumentHandler) GetMySubmissions(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	statusFilter := strings.TrimSpace(c.Query("status"))
	searchQuery := strings.TrimSpace(c.Query("q"))

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 20 }
	offset := (page - 1) * limit

	items, total, err := h.svc.GetMySubmissions(c.Context(), userID, statusFilter, searchQuery, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get submissions", err.Error())
	}

	totalPages := int64(1)
	if total > 0 {
		totalPages = (total + int64(limit) - 1) / int64(limit)
	}

	return response.Success(c, fiber.StatusOK, "My submissions retrieved", fiber.Map{
		"items": items,
		"pagination": fiber.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
		},
	})
}

// GetDeptSubmissions returns all document submissions in the authenticated user's department.
// @Summary Get department submission status
// @Tags Documents
// @Produce json
// @Param status query string false "Filter by status (pending, approved, rejected, active, draft)"
// @Param q query string false "Search by title or submitter name"
// @Param page query int false "Page number (default 1)"
// @Param limit query int false "Items per page (default 20)"
// @Success 200 {object} response.APIResponse
// @Router /documents/submissions [get]
// @Security BearerAuth
func (h *DocumentHandler) GetDeptSubmissions(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	user, err := h.svc.GetUser(c.Context(), userID)
	if err != nil || !user.DepartmentID.Valid {
		return response.Error(c, fiber.StatusBadRequest, "Department not found for current user", "")
	}
	deptID := user.DepartmentID.Bytes

	statusFilter := strings.TrimSpace(c.Query("status"))
	searchQuery := strings.TrimSpace(c.Query("q"))

	limit, _ := strconv.Atoi(c.Query("limit", "20"))
	page, _ := strconv.Atoi(c.Query("page", "1"))
	if page < 1 { page = 1 }
	if limit < 1 || limit > 100 { limit = 20 }
	offset := (page - 1) * limit

	items, total, err := h.svc.GetDeptSubmissions(c.Context(), deptID, statusFilter, searchQuery, limit, offset)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get submissions", err.Error())
	}

	totalPages := int64(1)
	if total > 0 {
		totalPages = (total + int64(limit) - 1) / int64(limit)
	}

	return response.Success(c, fiber.StatusOK, "Submissions retrieved", fiber.Map{
		"items": items,
		"pagination": fiber.Map{
			"total":       total,
			"page":        page,
			"limit":       limit,
			"total_pages": totalPages,
		},
	})
}

