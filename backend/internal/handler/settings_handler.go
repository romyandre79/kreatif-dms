package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/pkg/response"
	"github.com/google/uuid"
	"context"
)

type SettingsHandler struct {
	repo *repository.Queries
}

func NewSettingsHandler(repo *repository.Queries) *SettingsHandler {
	return &SettingsHandler{repo: repo}
}

// GetDashboardSettings returns settings for dashboard category
func (h *SettingsHandler) GetDashboardSettings(c fiber.Ctx) error {
	settings, err := h.repo.GetSystemSettingsByCategory(c.Context(), "dashboard")
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to fetch dashboard settings", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Dashboard settings retrieved", settings)
}

// UpdateDashboardSetting updates a specific setting
func (h *SettingsHandler) UpdateDashboardSetting(c fiber.Ctx) error {
	var req struct {
		Key   string `json:"key"`
		Value string `json:"value"`
	}

	if err := c.Bind().JSON(&req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	userID := c.Locals("user_id").(uuid.UUID)

	err := h.repo.UpdateSystemSetting(c.Context(), repository.UpdateSystemSettingParams{
		Category:  "dashboard",
		Key:       req.Key,
		Value:     repository.ToNullString(req.Value),
		UpdatedBy: uuid.NullUUID{UUID: userID, Valid: true},
	})

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update dashboard setting", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Dashboard setting updated", nil)
}
