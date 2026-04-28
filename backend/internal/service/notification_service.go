package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/infra"
)

type NotificationService struct {
	repo    repository.Querier
	waSvc   *infra.WhatsAppService
}

func NewNotificationService(repo repository.Querier, waSvc *infra.WhatsAppService) *NotificationService {
	return &NotificationService{repo: repo, waSvc: waSvc}
}

func (s *NotificationService) GetUserNotifications(ctx context.Context, userID uuid.UUID, limit, offset int32) ([]repository.Notification, error) {
	return s.repo.ListNotifications(ctx, repository.ListNotificationsParams{
		UserID: userID,
		Limit:  limit,
		Offset: offset,
	})
}

func (s *NotificationService) GetUnreadCount(ctx context.Context, userID uuid.UUID) (int64, error) {
	return s.repo.GetUnreadCount(ctx, userID)
}

func (s *NotificationService) MarkAsRead(ctx context.Context, notificationID, userID uuid.UUID) error {
	return s.repo.MarkAsRead(ctx, repository.MarkAsReadParams{
		ID:     notificationID,
		UserID: userID,
	})
}

func (s *NotificationService) MarkAllAsRead(ctx context.Context, userID uuid.UUID) error {
	return s.repo.MarkAllAsRead(ctx, userID)
}

func (s *NotificationService) CreateNotification(ctx context.Context, arg repository.CreateNotificationParams) (repository.Notification, error) {
	notif, err := s.repo.CreateNotification(ctx, arg)
	if err != nil {
		log.Printf("[NotificationService] Failed to create notification: %v", err)
		return repository.Notification{}, err
	}
	return notif, nil
}
