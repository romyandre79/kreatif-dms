package infra

import (
	"fmt"
	"log"
	"net/smtp"

	"github.com/kreatif/dms-backend/internal/config"
)

type MailService struct {
	cfg config.Config
}

func NewMailService(cfg config.Config) *MailService {
	return &MailService{cfg: cfg}
}

func (s *MailService) SendEmail(to string, subject string, body string) error {
	log.Printf("[MailService] Sending email to: %s (Subject: %s)", to, subject)

	// SMTP configuration
	host := s.cfg.SMTPHost
	port := s.cfg.SMTPPort
	user := s.cfg.SMTPUser
	password := s.cfg.SMTPPass

	// Message
	message := fmt.Sprintf("From: %s\nTo: %s\nSubject: %s\n\n%s", user, to, subject, body)

	// Auth (optional for Mailpit)
	var auth smtp.Auth
	if user != "" && password != "" {
		auth = smtp.PlainAuth("", user, password, host)
	}

	// Address
	addr := fmt.Sprintf("%s:%d", host, port)

	// Send
	err := smtp.SendMail(addr, auth, user, []string{to}, []byte(message))
	if err != nil {
		log.Printf("[MailService] Error sending email: %v", err)
		return err
	}

	log.Printf("[MailService] Email sent successfully to: %s", to)
	return nil
}
