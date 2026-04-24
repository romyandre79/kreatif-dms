package infra

import (
	"fmt"
	"net/smtp"

	"github.com/kreatif/dms-backend/internal/config"
)

type EmailService struct {
	cfg config.Config
}

func NewEmailService(cfg config.Config) *EmailService {
	return &EmailService{cfg: cfg}
}

func (s *EmailService) SendEmail(to, subject, body string) error {
	from := "noreply@kreatif-dms.com"
	msg := "From: " + from + "\n" +
		"To: " + to + "\n" +
		"Subject: " + subject + "\n" +
		"MIME-version: 1.0;\n" +
		"Content-Type: text/html; charset=\"UTF-8\";\n\n" +
		body

	addr := fmt.Sprintf("%s:%d", s.cfg.SMTPHost, s.cfg.SMTPPort)
	
	// If no auth needed (like Mailpit)
	var auth smtp.Auth
	if s.cfg.SMTPUser != "" {
		auth = smtp.PlainAuth("", s.cfg.SMTPUser, s.cfg.SMTPPass, s.cfg.SMTPHost)
	}

	return smtp.SendMail(addr, auth, from, []string{to}, []byte(msg))
}

func (s *EmailService) SendRegistrationNotification(email, fullName string) error {
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
	return s.SendEmail(email, subject, body)
}

func (s *EmailService) SendPasswordResetEmail(email, resetLink string) error {
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
	return s.SendEmail(email, subject, body)
}
