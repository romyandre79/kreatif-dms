package service

import (
	"context"
	"encoding/json"
	"log"

	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/internal/infra"
)

type NotificationService struct {
	repo     repository.Querier
	waSvc    *infra.WhatsAppService
	emailSvc *infra.EmailService
}

func NewNotificationService(repo repository.Querier, waSvc *infra.WhatsAppService, emailSvc *infra.EmailService) *NotificationService {
	return &NotificationService{repo: repo, waSvc: waSvc, emailSvc: emailSvc}
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

	// Trigger Email if type matches a template slug
	go func() {
		user, err := s.repo.GetUserByID(context.Background(), notif.UserID)
		if err != nil {
			return
		}

		data := map[string]string{
			"fullName": user.FullName,
			"body":     notif.Body.String,
			"title":    notif.Title,
		}

		// Merge metadata if present
		if len(notif.Metadata) > 0 {
			var meta map[string]string
			if err := json.Unmarshal(notif.Metadata, &meta); err == nil {
				for k, v := range meta {
					data[k] = v
				}
			}
		}

		// Try to send using the specific slug if provided in notification type
		_ = s.emailSvc.SendTemplatedEmail(context.Background(), notif.Type, user.Email, data)
	}()

	return notif, nil
}
