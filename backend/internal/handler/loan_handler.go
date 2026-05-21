package handler

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type LoanHandler struct {
	svc *service.DocumentService
}

func NewLoanHandler(svc *service.DocumentService) *LoanHandler {
	return &LoanHandler{svc: svc}
}

// Submit creates a new loan request with its items
func (h *LoanHandler) Submit(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	var req struct {
		DocumentIDs  []string `json:"document_ids"`
		Purpose      string   `json:"purpose"`
		DurationDays int      `json:"duration_days"`
		Notes        string   `json:"notes"`
		PickupDate   string   `json:"pickup_date"`
		TimeSlot     string   `json:"time_slot"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if len(req.DocumentIDs) == 0 {
		return response.Error(c, fiber.StatusBadRequest, "At least one document is required", "")
	}

	if req.Purpose == "" {
		req.Purpose = "Peminjaman Dokumen"
	}
	if req.DurationDays <= 0 {
		req.DurationDays = 7
	}

	// Parse document UUIDs
	var docIDs []uuid.UUID
	for _, idStr := range req.DocumentIDs {
		id, err := uuid.Parse(idStr)
		if err != nil {
			continue // Skip invalid UUIDs
		}
		docIDs = append(docIDs, id)
	}

	if len(docIDs) == 0 {
		return response.Error(c, fiber.StatusBadRequest, "No valid document IDs provided", "")
	}

	log.Printf("[LoanHandler] Submitting loan request from user %s with %d documents", userID, len(docIDs))

	// Create the loan request via service
	result, err := h.svc.CreateLoanRequest(c.Context(), service.CreateLoanRequestParams{
		UserID:       userID,
		DocumentIDs:  docIDs,
		Purpose:      req.Purpose,
		DurationDays: int32(req.DurationDays),
		Notes:        req.Notes,
		PickupDate:   req.PickupDate,
		TimeSlot:     req.TimeSlot,
	})
	if err != nil {
		log.Printf("[LoanHandler] Error creating loan request: %v", err)
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create loan request", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Loan request submitted successfully", result)
}

// ListMyLoans returns loan requests for the current user
func (h *LoanHandler) ListMyLoans(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)

	loans, err := h.svc.ListUserLoanRequests(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch loan requests", err.Error())
	}

	// Format the response for frontend consumption
	formatted := formatLoanListForFrontend(loans)

	return response.Success(c, fiber.StatusOK, "User loan requests retrieved", formatted)
}

// ListAll returns all loan requests (for admin/tracking)
func (h *LoanHandler) ListAll(c fiber.Ctx) error {
	limit, _ := strconv.Atoi(c.Query("limit", "50"))
	
	loans, err := h.svc.ListAllLoanRequests(c.Context(), int32(limit))
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch loan requests", err.Error())
	}

	formatted := formatAllLoanListForFrontend(loans)

	return response.Success(c, fiber.StatusOK, "All loan requests retrieved", formatted)
}

// GetByID returns a single loan request with its items
func (h *LoanHandler) GetByID(c fiber.Ctx) error {
	loanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid loan ID", err.Error())
	}

	loan, err := h.svc.GetLoanRequest(c.Context(), loanID)
	if err != nil {
		return response.Error(c, fiber.StatusNotFound, "Loan request not found", err.Error())
	}

	items, err := h.svc.GetLoanRequestItems(c.Context(), loanID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch loan items", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Loan request retrieved", fiber.Map{
		"loan":  loan,
		"items": items,
	})
}

// Approve approves a loan request (L1)
func (h *LoanHandler) Approve(c fiber.Ctx) error {
	loanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid loan ID", err.Error())
	}

	approverID := c.Locals("user_id").(uuid.UUID)

	if err := h.svc.ApproveLoanRequest(c.Context(), loanID, approverID); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to approve loan", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Loan request approved", nil)
}

// Reject rejects a loan request
func (h *LoanHandler) Reject(c fiber.Ctx) error {
	loanID, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid loan ID", err.Error())
	}

	approverID := c.Locals("user_id").(uuid.UUID)

	var req struct {
		Reason string `json:"reason"`
	}
	c.Bind().JSON(&req)

	if err := h.svc.RejectLoanRequestAction(c.Context(), loanID, approverID, req.Reason); err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to reject loan", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Loan request rejected", nil)
}

// formatLoanListForFrontend converts DB rows to a frontend-friendly format
func formatLoanListForFrontend(loans []repository.ListUserLoanRequestsRow) []fiber.Map {
	result := make([]fiber.Map, 0, len(loans))

	for _, loan := range loans {
		// Calculate remaining days
		remainingText := "Pending"
		remainingColor := "bg-slate-50 text-slate-400 border border-slate-100"
		statusColor := "bg-amber-50 text-amber-600 border border-amber-100"

		if loan.DueDate.Valid && loan.BorrowDate.Valid {
			daysLeft := int(time.Until(loan.DueDate.Time).Hours() / 24)
			if daysLeft > 5 {
				remainingText = fmt.Sprintf("%d Days Left", daysLeft)
				remainingColor = "bg-emerald-50 text-emerald-500 border border-emerald-100"
			} else if daysLeft > 0 {
				remainingText = fmt.Sprintf("%d Days Left", daysLeft)
				remainingColor = "bg-orange-50 text-orange-500 border border-orange-100"
			} else if daysLeft == 0 {
				remainingText = "Due Today"
				remainingColor = "bg-red-50 text-red-500 border border-red-100"
			} else {
				remainingText = fmt.Sprintf("Overdue %d Days", -daysLeft)
				remainingColor = "bg-red-50 text-red-500 border border-red-100"
			}
		}

		switch loan.Status {
		case "pending":
			statusColor = "bg-amber-50 text-amber-600 border border-amber-100"
		case "l1_approved", "l2_approved", "active":
			statusColor = "bg-blue-50 text-blue-600 border border-blue-100"
		case "rejected":
			statusColor = "bg-red-50 text-red-500 border border-red-100"
			remainingText = "Rejected"
			remainingColor = "bg-red-50 text-red-500 border border-red-100"
		case "returned":
			statusColor = "bg-emerald-50 text-emerald-600 border border-emerald-100"
			remainingText = "Returned"
			remainingColor = "bg-emerald-50 text-emerald-500 border border-emerald-100"
		case "overdue":
			statusColor = "bg-red-50 text-red-500 border border-red-100"
		}

		// Format dates
		checkoutDate := ""
		dueDate := ""
		approvedDate := ""

		if loan.BorrowDate.Valid {
			checkoutDate = loan.BorrowDate.Time.Format("Jan 02, 2006")
		}
		if loan.DueDate.Valid {
			dueDate = loan.DueDate.Time.Format("Jan 02, 2006")
		}
		if loan.L1ApprovedAt.Valid {
			approvedDate = loan.L1ApprovedAt.Time.Format("Jan 02, 2006 • 03:04 PM")
		}
		createdAt := ""
		if loan.CreatedAt.Valid {
			createdAt = loan.CreatedAt.Time.Format("Jan 02, 2006 • 03:04 PM")
		}

		displayStatus := loan.Status
		switch loan.Status {
		case "pending":
			displayStatus = "Pending Approval"
		case "l1_approved":
			displayStatus = "Dalam Persiapan"
		case "l2_approved":
			displayStatus = "Siap Diambil"
		case "active":
			displayStatus = "Active"
		case "returned":
			displayStatus = "Returned"
		case "rejected":
			displayStatus = "Rejected"
		case "overdue":
			displayStatus = "Overdue"
		}

		result = append(result, fiber.Map{
			"id":             loan.ID,
			"no":             loan.RequestNo,
			"docsCount":      loan.ItemsCount,
			"checkoutDate":   checkoutDate,
			"dueDate":        dueDate,
			"remainingText":  remainingText,
			"remainingColor": remainingColor,
			"status":         displayStatus,
			"statusColor":    statusColor,
			"approvedDate":   approvedDate,
			"readyDate":      approvedDate,
			"onLoanDate":     checkoutDate,
			"createdAt":      createdAt,
			"userName":       loan.UserName,
			"departmentName": loan.DepartmentName.String,
			"purpose":        loan.Purpose,
			"durationDays":   loan.DurationDays,
			"rawStatus":          loan.Status,
			"l2RejectionReason": loan.L2RejectionReason.String,
		})
	}

	return result
}

// formatAllLoanListForFrontend converts DB rows to a frontend-friendly format
func formatAllLoanListForFrontend(loans []repository.ListAllLoanRequestsRow) []fiber.Map {
	result := make([]fiber.Map, 0, len(loans))

	for _, loan := range loans {
		// Calculate remaining days
		remainingText := "Pending"
		remainingColor := "bg-slate-50 text-slate-400 border border-slate-100"
		statusColor := "bg-amber-50 text-amber-600 border border-amber-100"

		if loan.DueDate.Valid && loan.BorrowDate.Valid {
			daysLeft := int(time.Until(loan.DueDate.Time).Hours() / 24)
			if daysLeft > 5 {
				remainingText = fmt.Sprintf("%d Days Left", daysLeft)
				remainingColor = "bg-emerald-50 text-emerald-500 border border-emerald-100"
			} else if daysLeft > 0 {
				remainingText = fmt.Sprintf("%d Days Left", daysLeft)
				remainingColor = "bg-orange-50 text-orange-500 border border-orange-100"
			} else if daysLeft == 0 {
				remainingText = "Due Today"
				remainingColor = "bg-red-50 text-red-500 border border-red-100"
			} else {
				remainingText = fmt.Sprintf("Overdue %d Days", -daysLeft)
				remainingColor = "bg-red-50 text-red-500 border border-red-100"
			}
		}

		switch loan.Status {
		case "pending":
			statusColor = "bg-amber-50 text-amber-600 border border-amber-100"
		case "l1_approved", "l2_approved", "active":
			statusColor = "bg-blue-50 text-blue-600 border border-blue-100"
		case "rejected":
			statusColor = "bg-red-50 text-red-500 border border-red-100"
			remainingText = "Rejected"
			remainingColor = "bg-red-50 text-red-500 border border-red-100"
		case "returned":
			statusColor = "bg-emerald-50 text-emerald-600 border border-emerald-100"
			remainingText = "Returned"
			remainingColor = "bg-emerald-50 text-emerald-500 border border-emerald-100"
		case "overdue":
			statusColor = "bg-red-50 text-red-500 border border-red-100"
		}

		// Format dates
		checkoutDate := ""
		dueDate := ""
		approvedDate := ""

		if loan.BorrowDate.Valid {
			checkoutDate = loan.BorrowDate.Time.Format("Jan 02, 2006")
		}
		if loan.DueDate.Valid {
			dueDate = loan.DueDate.Time.Format("Jan 02, 2006")
		}
		if loan.L1ApprovedAt.Valid {
			approvedDate = loan.L1ApprovedAt.Time.Format("Jan 02, 2006 • 03:04 PM")
		}
		returnDate := ""
		if loan.ReturnDate.Valid {
			returnDate = loan.ReturnDate.Time.Format("Jan 02, 2006")
		}
		createdAt := ""
		if loan.CreatedAt.Valid {
			createdAt = loan.CreatedAt.Time.Format("Jan 02, 2006 • 03:04 PM")
		}

		displayStatus := loan.Status
		switch loan.Status {
		case "pending":
			displayStatus = "Pending Approval"
		case "l1_approved":
			displayStatus = "Dalam Persiapan"
		case "l2_approved":
			displayStatus = "Siap Diambil"
		case "active":
			displayStatus = "Active"
		case "returned":
			displayStatus = "Returned"
		case "rejected":
			displayStatus = "Rejected"
		case "overdue":
			displayStatus = "Overdue"
		}

		result = append(result, fiber.Map{
			"id":             loan.ID,
			"no":             loan.RequestNo,
			"docsCount":      loan.ItemsCount,
			"checkoutDate":   checkoutDate,
			"dueDate":        dueDate,
			"remainingText":  remainingText,
			"remainingColor": remainingColor,
			"status":         displayStatus,
			"statusColor":    statusColor,
			"approvedDate":   approvedDate,
			"readyDate":      approvedDate,
			"onLoanDate":     checkoutDate,
			"returnDate":     returnDate,
			"createdAt":      createdAt,
			"userName":       loan.UserName,
			"departmentName": loan.DepartmentName.String,
			"purpose":        loan.Purpose,
			"durationDays":   loan.DurationDays,
			"rawStatus":      loan.Status,
			"l2RejectionReason": loan.L2RejectionReason.String,
		})
	}

	return result
}

func (h *LoanHandler) GetPenaltyPolicy(c fiber.Ctx) error {
	setting, err := h.svc.GetSystemSetting(c.Context(), "general", "penalty_policy")
	if err != nil {
		return response.Success(c, fiber.StatusOK, "Penalty policy retrieved", fiber.Map{
			"policy": "Denda keterlambatan: IDR 50.000 / Dokumen / Hari",
		})
	}
	return response.Success(c, fiber.StatusOK, "Penalty policy retrieved", fiber.Map{
		"policy": setting.Value.String,
	})
}

