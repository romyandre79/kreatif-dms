package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"

	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
)

type EmailService struct {
	cfg  config.Config
	repo repository.Querier
}
type smtpNodeConfig struct {
	FromEmail string `json:"from_email"`
	Auth      bool   `json:"auth"`
	User      string `json:"user"`
	Pass      string `json:"pass"`
}

func (s *EmailService) getSMTPConfig(ctx context.Context) (string, string, smtpNodeConfig, error) {
	node, err := s.repo.GetIntegrationNodeByType(ctx, "SMTP")
	if err != nil {
		return "", "", smtpNodeConfig{}, fmt.Errorf("SMTP node not found in database: %v", err)
	}

	var nodeCfg smtpNodeConfig
	if err := json.Unmarshal(node.ConfigJson, &nodeCfg); err != nil {
		return "", "", smtpNodeConfig{}, fmt.Errorf("failed to parse SMTP config JSON: %v", err)
	}

	if nodeCfg.FromEmail == "" {
		nodeCfg.FromEmail = "noreply@kreatif-dms.com"
	}

	host := strings.Split(node.Endpoint, ":")[0]

	return node.Endpoint, host, nodeCfg, nil
}

func NewEmailService(cfg config.Config, repo repository.Querier) *EmailService {
	return &EmailService{cfg: cfg, repo: repo}
}

func (s *EmailService) SendEmail(ctx context.Context, to, subject, body string) error {
	addr, host, nodeCfg, err := s.getSMTPConfig(ctx)
	if err != nil {
		// Fallback to config if DB fails
		addr = fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort)
		host = s.cfg.SMTPHost
		nodeCfg.FromEmail = "noreply@kreatif-dms.com"
		nodeCfg.User = s.cfg.SMTPUser
		nodeCfg.Pass = s.cfg.SMTPPass
		nodeCfg.Auth = s.cfg.SMTPUser != ""
	}

	msg := "From: " + nodeCfg.FromEmail + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body

	var auth smtp.Auth
	if nodeCfg.Auth {
		auth = smtp.PlainAuth("", nodeCfg.User, nodeCfg.Pass, host)
	}

	return smtp.SendMail(addr, auth, nodeCfg.FromEmail, []string{to}, []byte(msg))
}

func (s *EmailService) SendRegistrationNotification(ctx context.Context, email, fullName string) error {
	subject := "Kreatif DMS - Registration Successful"
	body := fmt.Sprintf(`
		<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
			<h2 style="color: #1E3A5F;">Welcome to Kreatif DMS</h2>
			<p>Hi <strong>%s</strong>,</p>
			<p>Thank you for registering at Kreatif DMS.</p>
			<div style="background: #f0f7ff; padding: 15px; border-radius: 8px; border-left: 4px solid #1E3A5F; margin: 20px 0;">
				<p style="margin: 0;"><strong>Account Status: Pending Approval</strong></p>
				<p style="margin: 5px 0 0 0; color: #666;">Your account is currently being reviewed by our administrator. You will receive another email once your account is approved.</p>
			</div>
			<p>Best regards,<br>Kreatif DMS Team</p>
		</div>
	`, fullName)
	return s.SendEmail(ctx, email, subject, body)
}

func (s *EmailService) SendPasswordResetEmail(ctx context.Context, email, resetLink string) error {
	subject := "Kreatif DMS - Password Reset Request"
	body := fmt.Sprintf(`
		<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
			<h2 style="color: #1E3A5F;">Reset Your Password</h2>
			<p>We received a request to reset your password for your Kreatif DMS account.</p>
			<p>Click the button below to reset it:</p>
			<div style="text-align: center; margin: 30px 0;">
				<a href="%s" style="background: #1E3A5F; color: white; padding: 12px 25px; text-decoration: none; border-radius: 5px; font-weight: bold; display: inline-block;">Reset Password</a>
			</div>
			<p style="color: #666; font-size: 12px;">If you didn't request this, you can safely ignore this email. The link will expire in 1 hour.</p>
			<p>Best regards,<br>Kreatif DMS Team</p>
		</div>
	`, resetLink)
	return s.SendEmail(ctx, email, subject, body)
}
func (s *EmailService) TestConnection(ctx context.Context, addr, host, user, pass string, useAuth bool) error {
	c, err := smtp.Dial(addr)
	if err != nil {
		return fmt.Errorf("failed to connect to SMTP server: %v", err)
	}
	defer c.Close()

	if useAuth {
		auth := smtp.PlainAuth("", user, pass, host)
		if err := c.Auth(auth); err != nil {
			return fmt.Errorf("SMTP authentication failed: %v", err)
		}
	}

	return nil
}
