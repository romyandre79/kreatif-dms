package service

import (
	"context"
	"log"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
)

type DashboardService struct {
	repo *repository.Queries
}

func NewDashboardService(repo *repository.Queries) *DashboardService {
	return &DashboardService{repo: repo}
}

func (s *DashboardService) GetRepo() *repository.Queries {
	return s.repo
}

func (s *DashboardService) GetAnnouncements(ctx context.Context) (map[string]interface{}, error) {
	ann, err := s.repo.GetActiveAnnouncement(ctx)
	if err != nil {
		return map[string]interface{}{"active": false}, nil
	}

	return map[string]interface{}{
		"id":      ann.ID,
		"title":   ann.Title,
		"message": ann.Message,
		"notes":   ann.Notes.String,
		"active":  ann.IsActive,
	}, nil
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
func (s *DashboardService) GetManagerSummary(ctx context.Context, userID uuid.UUID) (interface{}, error) {
	log.Printf("[DashboardService] Fetching stats for Manager: %v", userID)
	stats, err := s.repo.GetManagerDailyStats(ctx, pgtype.UUID{Bytes: userID, Valid: true})
	if err != nil {
		return nil, err
	}
	log.Printf("[DashboardService] Stats retrieved: daily_received=%d, active_tasks=%d", stats.DailyReceived, stats.ActiveTasks)

	activities, err := s.repo.GetManagerRecentActivities(ctx, repository.GetManagerRecentActivitiesParams{
		HeadID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  10,
	})
	if err != nil {
		activities = []repository.GetManagerRecentActivitiesRow{}
	}

	tasks, err := s.repo.GetUserPriorityTasks(ctx, repository.GetUserPriorityTasksParams{
		ApproverID: userID,
		Limit:      10,
	})
	if err != nil {
		tasks = []repository.GetUserPriorityTasksRow{}
	}

	topSubmitters, err := s.repo.GetManagerTopSubmitters(ctx, repository.GetManagerTopSubmittersParams{
		HeadID: pgtype.UUID{Bytes: userID, Valid: true},
		Limit:  5,
	})
	if err != nil {
		topSubmitters = []repository.GetManagerTopSubmittersRow{}
	}

	announcement, _ := s.GetAnnouncements(ctx)

	return map[string]interface{}{
		"stats":          stats,
		"activities":     activities,
		"tasks":          tasks,
		"top_submitters": topSubmitters,
		"announcement":   announcement,
	}, nil
}
