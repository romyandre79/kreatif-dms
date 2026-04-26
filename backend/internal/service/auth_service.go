package service

import (
	"context"
	"errors"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/auth"
	"github.com/kreatif/dms-backend/internal/config"
	"github.com/kreatif/dms-backend/internal/repository"
	"golang.org/x/crypto/bcrypt"
	"log"
	"fmt"

	"github.com/kreatif/dms-backend/internal/infra"
)

type AuthService struct {
	repo     repository.Querier
	cfg      config.Config
	ldapSvc  *infra.LDAPService
	emailSvc *infra.EmailService
}

func NewAuthService(repo repository.Querier, cfg config.Config, ldapSvc *infra.LDAPService, emailSvc *infra.EmailService) *AuthService {
	return &AuthService{repo: repo, cfg: cfg, ldapSvc: ldapSvc, emailSvc: emailSvc}
}

type LoginResponse struct {
	AccessToken  string    `json:"access_token"`
	RefreshToken string    `json:"refresh_token"`
	UserID       uuid.UUID `json:"user_id"`
	FullName     string    `json:"full_name"`
	Role         string    `json:"role"`
	AvatarUrl    string    `json:"avatar_url"`
	SignatureUrl string    `json:"signature_url"`
}

func (s *AuthService) Login(ctx context.Context, identifier, password, authType string) (*LoginResponse, error) {
	log.Printf("[AuthService] Login attempt for: %s (Type: %s)", identifier, authType)

	if authType == "sso" && s.cfg.LDAPEnabled {
		// 1. Corporate LDAP Authentication
		log.Printf("[AuthService] Attempting Corporate LDAP login for: %s", identifier)
		ldapUser, ldapErr := s.ldapSvc.Authenticate(identifier, password)
		if ldapErr == nil {
			// LDAP Success! Find or create user in our DB
			row, err := s.repo.GetUserByEmail(ctx, ldapUser.Email)
			if err != nil {
				// JIT (Just-In-Time) Provisioning
				log.Printf("[AuthService] Creating JIT user for LDAP: %s", ldapUser.Email)
				roleID, _ := s.repo.GetRoleIDByName(ctx, "user")
				user, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
					FullName:     ldapUser.FullName,
					Email:        ldapUser.Email,
					PasswordHash: "LDAP_AUTH",
					RoleID:       roleID,
					Status:       "approved",
				})
				if err != nil {
					return nil, err
				}
				return s.generateLoginResponse(user.ID, user.FullName, "user", user.DepartmentID, user.AvatarUrl.String, user.SignatureUrl.String)
			}
			roleName := "user"
			if row.RoleName.Valid {
				roleName = row.RoleName.String
			}
			return s.generateLoginResponse(row.ID, row.FullName, roleName, row.DepartmentID, row.AvatarUrl.String, row.SignatureUrl.String)
		}
		return nil, fmt.Errorf("corporate login failed: %v", ldapErr)
	}

	// 2. Local Database Authentication (for Emails or fallback)
	log.Printf("[AuthService] Attempting Local DB login for email: [%s]", identifier)
	row, err := s.repo.GetUserByEmail(ctx, identifier)
	if err != nil {
		log.Printf("[AuthService] User not found: [%s]. Error: %v", identifier, err)
		
		// Diagnostic: List all users to see if we're in the right DB
		users, _ := s.repo.ListUsers(ctx)
		var emails []string
		for _, u := range users {
			emails = append(emails, u.Email)
		}
		log.Printf("[AuthService] DIAGNOSTIC: Total users in DB: %d. Available emails: %v", len(users), emails)
		
		return nil, errors.New("user not found in local database")
	}

	err = bcrypt.CompareHashAndPassword([]byte(row.PasswordHash), []byte(password))
	if err != nil {
		log.Printf("[AuthService] Password mismatch for user: %s", identifier)
		return nil, errors.New("invalid password")
	}

	if row.Status != "approved" {
		return nil, fmt.Errorf("your account status is %s. Please wait for administrator approval", row.Status)
	}

	roleName := "user"
	if row.RoleName.Valid {
		roleName = row.RoleName.String
	}

	return s.generateLoginResponse(row.ID, row.FullName, roleName, row.DepartmentID, row.AvatarUrl.String, row.SignatureUrl.String)
}

func (s *AuthService) Register(ctx context.Context, fullName, email, password string) error {
	log.Printf("[AuthService] Registration request for: %s", email)

	// 1. Hash password
	hashed, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 2. Get default role ID (user)
	roleID, err := s.repo.GetRoleIDByName(ctx, "user")
	if err != nil {
		return err
	}

	// 3. Create user with 'approved' status (auto-approve for testing)
	user, err := s.repo.CreateUser(ctx, repository.CreateUserParams{
		FullName:     fullName,
		Email:        email,
		PasswordHash: string(hashed),
		RoleID:       roleID,
		Status:       "approved",
	})
	if err != nil {
		return err
	}

	// 4. Send email notification
	go func() {
		err := s.emailSvc.SendRegistrationNotification(user.Email, user.FullName)
		if err != nil {
			log.Printf("[AuthService] Failed to send registration email to %s: %v", user.Email, err)
		}
	}()

	return nil
}

func (s *AuthService) RequestPasswordReset(ctx context.Context, email string) error {
	user, err := s.repo.GetUserByEmail(ctx, email)
	if err != nil {
		// Don't reveal if user exists or not for security, but we'll just return nil or error
		return nil 
	}

	// In a real app, generate a token and save to DB
	resetLink := fmt.Sprintf("http://localhost:3000/reset-password?token=%s", uuid.New().String())

	go func() {
		err := s.emailSvc.SendPasswordResetEmail(user.Email, resetLink)
		if err != nil {
			log.Printf("[AuthService] Failed to send reset email to %s: %v", user.Email, err)
		}
	}()

	return nil
}

func (s *AuthService) ListPendingUsers(ctx context.Context) ([]repository.User, error) {
	return s.repo.ListPendingUsers(ctx)
}

func (s *AuthService) ApproveUser(ctx context.Context, userID uuid.UUID) error {
	user, err := s.repo.UpdateUserStatus(ctx, repository.UpdateUserStatusParams{
		ID:     userID,
		Status: "approved",
	})
	if err != nil {
		return err
	}

	// Send approval notification
	go func() {
		subject := "Kreatif DMS - Account Approved"
		body := fmt.Sprintf(`
			<div style="font-family: sans-serif; max-width: 600px; margin: auto; padding: 20px; border: 1px solid #eee; border-radius: 10px;">
				<h2 style="color: #22c55e;">Account Approved!</h2>
				<p>Hi <strong>%s</strong>,</p>
				<p>Great news! Your Kreatif DMS account has been approved by the administrator.</p>
				<p>You can now log in using your email and password.</p>
				<div style="text-align: center; margin: 30px 0;">
					<a href="http://localhost:3000/login" style="background: #1E3A5F; color: white; padding: 12px 25px; text-decoration: none; border-radius: 5px; font-weight: bold; display: inline-block;">Login Now</a>
				</div>
				<p>Best regards,<br>Kreatif DMS Team</p>
			</div>
		`, user.FullName)
		_ = s.emailSvc.SendEmail(user.Email, subject, body)
	}()

	return nil
}

func (s *AuthService) generateLoginResponse(userID uuid.UUID, fullName, roleName string, deptIDRaw interface{}, avatarUrl, signatureUrl string) (*LoginResponse, error) {
	var deptID uuid.UUID
	
	// Handle pgtype.UUID or other types if necessary
	if val, ok := deptIDRaw.(pgtype.UUID); ok && val.Valid {
		deptID = val.Bytes
	}

	payload := auth.TokenPayload{
		UserID:       userID,
		Role:         roleName,
		DepartmentID: deptID,
	}

	access, err := auth.GenerateToken(payload, s.cfg.JWTSecret, s.cfg.AccessTokenTTL)
	if err != nil {
		return nil, err
	}

	refresh, err := auth.GenerateToken(payload, s.cfg.JWTSecret, s.cfg.RefreshTokenTTL)
	if err != nil {
		return nil, err
	}

	return &LoginResponse{
		AccessToken:  access,
		RefreshToken: refresh,
		UserID:       userID,
		FullName:     fullName,
		Role:         roleName,
		AvatarUrl:    avatarUrl,
		SignatureUrl: signatureUrl,
	}, nil
}
func (s *AuthService) SetPIN(ctx context.Context, userID uuid.UUID, pin string) error {
	if len(pin) != 6 {
		return errors.New("PIN must be 6 digits")
	}

	// Hash the PIN using bcrypt
	hashed, err := bcrypt.GenerateFromPassword([]byte(pin), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	return s.repo.UpdateUserPIN(ctx, repository.UpdateUserPINParams{
		ID:  userID,
		Pin: pgtype.Text{String: string(hashed), Valid: true},
	})
}

func (s *AuthService) VerifyPIN(ctx context.Context, userID uuid.UUID, pin string) (bool, error) {
	user, err := s.repo.GetUserByID(ctx, userID)
	if err != nil {
		return false, err
	}

	if !user.Pin.Valid {
		return false, errors.New("PIN not set for this user")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Pin.String), []byte(pin))
	if err != nil {
		return false, nil // Invalid PIN
	}

	return true, nil
}
