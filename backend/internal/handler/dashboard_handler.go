package handler

import (
	"log"
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
)

type DashboardHandler struct {
	svc *service.DashboardService
}

func NewDashboardHandler(svc *service.DashboardService) *DashboardHandler {
	return &DashboardHandler{svc: svc}
}

// GetSummary returns dashboard summary stats
// @Summary Get dashboard summary
// @Description Fetch daily stats and counters for the dashboard
// @Tags Dashboard
// @Produce json
// @Success 200 {object} response.APIResponse
// @Router /dashboard/summary [get]
// @Security BearerAuth
func (h *DashboardHandler) GetSummary(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	rawRole := c.Locals("user_role").(string)
	role := ""
	for _, char := range rawRole {
		if char >= 'A' && char <= 'Z' {
			role += string(char + 32)
		} else {
			role += string(char)
		}
	}
	log.Printf("[DashboardHandler] Request from UserID: %v, Role: %v", userID, role)

	// If superadmin or admin, return global summary
	if role == "superadmin" || role == "admin" || role == "admin doc controller" || role == "kepala doc controller" || role == "admin dc" || role == "kepala dc" {
		summary, err := h.svc.GetSummary(c.Context())
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to get dashboard summary", err.Error())
		}

		summaryMap := summary.(map[string]interface{})

		activities, err := h.svc.GetRecentActivities(c.Context(), 10)
		if err != nil {
			activities = []repository.GetRecentActivitiesRow{}
		}

		tasks, err := h.svc.GetRepo().GetUserPriorityTasks(c.Context(), repository.GetUserPriorityTasksParams{
			ApproverID: userID,
			Limit:      10,
		})
		if err != nil {
			tasks = []repository.GetUserPriorityTasksRow{}
		}

		// Filter out loan_requests if role is not kepala doc controller
		if role != "kepala doc controller" && role != "kepala dc" {
			var filteredTasks []repository.GetUserPriorityTasksRow
			for _, t := range tasks {
				if t.EntityType != "loan_request" && t.EntityType != "loan" {
					filteredTasks = append(filteredTasks, t)
				}
			}
			tasks = filteredTasks
		}

		return response.Success(c, fiber.StatusOK, "Dashboard summary retrieved", fiber.Map{
			"stats":        summaryMap["stats"],
			"announcement": summaryMap["announcement"],
			"activities":   activities,
			"tasks":        tasks,
		})
	}

	// For managers, return department/team summary
	if role == "manajer" {
		data, err := h.svc.GetManagerSummary(c.Context(), userID)
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to get manager dashboard summary", err.Error())
		}
		return response.Success(c, fiber.StatusOK, "Manager dashboard summary retrieved", data)
	}

	// For other roles, return user-specific summary
	data, err := h.svc.GetUserSummary(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get user dashboard summary", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User dashboard summary retrieved", data)
}
