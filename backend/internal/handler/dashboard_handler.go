package handler

import (
	"github.com/gofiber/fiber/v3"
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
	stats, err := h.svc.GetSummary(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get dashboard summary", err.Error())
	}

	activities, err := h.svc.GetRecentActivities(c.Context(), 10)
	if err != nil {
		activities = []repository.GetRecentActivitiesRow{}
	}

	tasks, err := h.svc.GetPriorityTasks(c.Context(), 10)
	if err != nil {
		tasks = []repository.GetPriorityTasksRow{}
	}

	return response.Success(c, fiber.StatusOK, "Dashboard summary retrieved", fiber.Map{
		"stats":      stats,
		"activities": activities,
		"tasks":      tasks,
	})
}
