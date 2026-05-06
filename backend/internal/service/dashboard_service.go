package service

import (
	"context"
	"github.com/kreatif/dms-backend/internal/repository"
)

type DashboardService struct {
	repo *repository.Queries
}

func NewDashboardService(repo *repository.Queries) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetSummary(ctx context.Context) (repository.GetDailyStatsRow, error) {
	return s.repo.GetDailyStats(ctx)
}

func (s *DashboardService) GetRecentActivities(ctx context.Context, limit int32) ([]repository.GetRecentActivitiesRow, error) {
	return s.repo.GetRecentActivities(ctx, limit)
}

func (s *DashboardService) GetPriorityTasks(ctx context.Context, limit int32) ([]repository.GetPriorityTasksRow, error) {
	return s.repo.GetPriorityTasks(ctx, limit)
}
