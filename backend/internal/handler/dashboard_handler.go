package handler

import (
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
	role := c.Locals("user_role").(string)

	// If superadmin or admin, return global summary
	if role == "superadmin" || role == "admin" {
		summary, err := h.svc.GetSummary(c.Context())
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to get dashboard summary", err.Error())
		}

		summaryMap := summary.(map[string]interface{})

		activities, err := h.svc.GetRecentActivities(c.Context(), 10)
		if err != nil {
			activities = []repository.GetRecentActivitiesRow{}
		}

		tasks, err := h.svc.GetPriorityTasks(c.Context(), 10)
		if err != nil {
			tasks = []repository.GetPriorityTasksRow{}
		}

		return response.Success(c, fiber.StatusOK, "Dashboard summary retrieved", fiber.Map{
			"stats":        summaryMap["stats"],
			"announcement": summaryMap["announcement"],
			"activities":   activities,
			"tasks":        tasks,
		})
	}

	// For managers, return department summary
	if role == "manajer" {
		// Get user's department
		user, err := h.svc.GetRepo().GetUserByID(c.Context(), userID)
		if err != nil || !user.DepartmentID.Valid {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to identify manager's department", "")
		}

		data, err := h.svc.GetManagerSummary(c.Context(), userID, user.DepartmentID.Bytes)
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
