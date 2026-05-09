package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
)

type DashboardService struct {
	repo *repository.Queries
}

func NewDashboardService(repo *repository.Queries) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetAnnouncements(ctx context.Context) (map[string]interface{}, error) {
	settings, err := s.repo.GetSystemSettingsByCategory(ctx, "dashboard")
	if err != nil {
		return nil, err
	}

	announcement := map[string]interface{}{
		"active": false,
	}

	for _, setting := range settings {
		switch setting.Key {
		case "announcement_title":
			announcement["title"] = setting.Value.String
		case "announcement_message":
			announcement["message"] = setting.Value.String
		case "announcement_notes":
			announcement["notes"] = setting.Value.String
		case "announcement_active":
			announcement["active"] = setting.Value.String == "true"
		}
	}

	return announcement, nil
}

func (s *DashboardService) GetSummary(ctx context.Context) (interface{}, error) {
	stats, err := s.repo.GetDailyStats(ctx)
	if err != nil {
		return nil, err
	}

	announcement, _ := s.GetAnnouncements(ctx)

	return map[string]interface{}{
		"stats":        stats,
		"announcement": announcement,
	}, nil
}

func (s *DashboardService) GetUserSummary(ctx context.Context, userID uuid.UUID) (interface{}, error) {
	stats, err := s.repo.GetUserDailyStats(ctx, userID)
	if err != nil {
		return nil, err
	}

	activities, err := s.repo.GetUserRecentActivities(ctx, repository.GetUserRecentActivitiesParams{
		UserID: userID,
		Limit:  10,
	})
	if err != nil {
		activities = []repository.GetUserRecentActivitiesRow{}
	}

	tasks, err := s.repo.GetUserPriorityTasks(ctx, repository.GetUserPriorityTasksParams{
		ApproverID: userID,
		Limit:      10,
	})
	if err != nil {
		tasks = []repository.GetUserPriorityTasksRow{}
	}

	loans, err := s.repo.GetUserLoanHistory(ctx, repository.GetUserLoanHistoryParams{
		UserID: userID,
		Limit:  10,
	})
	if err != nil {
		loans = []repository.GetUserLoanHistoryRow{}
	}

	announcement, _ := s.GetAnnouncements(ctx)

	return map[string]interface{}{
		"stats":        stats,
		"activities":   activities,
		"tasks":        tasks,
		"loans":        loans,
		"announcement": announcement,
	}, nil
}

func (s *DashboardService) GetRecentActivities(ctx context.Context, limit int32) ([]repository.GetRecentActivitiesRow, error) {
	return s.repo.GetRecentActivities(ctx, limit)
}

func (s *DashboardService) GetPriorityTasks(ctx context.Context, limit int32) ([]repository.GetPriorityTasksRow, error) {
	return s.repo.GetPriorityTasks(ctx, limit)
}
