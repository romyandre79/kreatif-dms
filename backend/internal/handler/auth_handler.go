package handler

import (
	"github.com/google/uuid"
	"github.com/gofiber/fiber/v3"
	"github.com/kreatif/dms-backend/internal/service"
	"github.com/kreatif/dms-backend/pkg/response"
	"github.com/kreatif/dms-backend/pkg/utils"
)

type AuthHandler struct {
	svc *service.AuthService
}

func NewAuthHandler(svc *service.AuthService) *AuthHandler {
	return &AuthHandler{svc: svc}
}

type loginRequest struct {
	Identifier string `json:"identifier" validate:"required"`
	Password   string `json:"password" validate:"required"`
	AuthType   string `json:"auth_type"` // "sso" or "local"
}

// Login godoc
// @Summary Login user
// @Description Login with email and password to get JWT tokens
// @Tags Auth
// @Accept json
// @Produce json
// @Param body body loginRequest true "Login Credentials"
// @Success 200 {object} response.APIResponse{data=service.LoginResponse}
// @Failure 401 {object} response.APIResponse
// @Router /auth/login [post]
func (h *AuthHandler) Login(c fiber.Ctx) error {
	req := new(loginRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Validate request
	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	res, err := h.svc.Login(c.Context(), req.Identifier, req.Password, req.AuthType)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Login failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Login successful", res)
}

type registerRequest struct {
	FullName string `json:"full_name" validate:"required,min=3"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=6"`
}

func (h *AuthHandler) Register(c fiber.Ctx) error {
	req := new(registerRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	err := h.svc.Register(c.Context(), req.FullName, req.Email, req.Password)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Registration failed", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "Registration successful. Please wait for admin approval.", nil)
}

func (h *AuthHandler) ListPendingUsers(c fiber.Ctx) error {
	users, err := h.svc.ListPendingUsers(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list pending users", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Pending users retrieved", users)
}

func (h *AuthHandler) ApproveUser(c fiber.Ctx) error {
	idStr := c.Params("id")
	userID, err := uuid.Parse(idStr)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	err = h.svc.ApproveUser(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to approve user", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User approved successfully", nil)
}

type forgotPasswordRequest struct {
	Email string `json:"email" validate:"required,email"`
}

func (h *AuthHandler) ForgotPassword(c fiber.Ctx) error {
	req := new(forgotPasswordRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	err := h.svc.RequestPasswordReset(c.Context(), req.Email)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to request password reset", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Reset link sent if email exists", nil)
}

type pinRequest struct {
	PIN string `json:"pin" validate:"required,len=6,numeric"`
}

func (h *AuthHandler) SetPIN(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	req := new(pinRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	err := h.svc.SetPIN(c.Context(), userID, req.PIN)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to set PIN", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "PIN set successfully", nil)
}

type setUserPinRequest struct {
	UserID uuid.UUID `json:"user_id" validate:"required"`
	PIN    string    `json:"pin" validate:"required,len=6,numeric"`
}

func (h *AuthHandler) SetUserPIN(c fiber.Ctx) error {
	req := new(setUserPinRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	err := h.svc.SetPIN(c.Context(), req.UserID, req.PIN)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to set user PIN", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User PIN set successfully", nil)
}

func (h *AuthHandler) VerifyPIN(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	req := new(pinRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	valid, err := h.svc.VerifyPIN(c.Context(), userID, req.PIN)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "PIN verification failed", err.Error())
	}

	if !valid {
		return response.Error(c, fiber.StatusUnauthorized, "Invalid PIN", "")
	}

	return response.Success(c, fiber.StatusOK, "PIN verified", nil)
}

func (h *AuthHandler) SetupMFA(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	secret, url, err := h.svc.GenerateMFASecret(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to generate MFA secret", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "MFA setup initiated", fiber.Map{
		"secret": secret,
		"url":    url,
	})
}

func (h *AuthHandler) VerifyMFA(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	type request struct {
		Secret string `json:"secret" validate:"required"`
		Code   string `json:"code" validate:"required"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	err := h.svc.VerifyMFAAndEnable(c.Context(), userID, req.Secret, req.Code)
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "MFA verification failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "MFA enabled successfully", nil)
}

func (h *AuthHandler) LoginMFA(c fiber.Ctx) error {
	type request struct {
		UserID uuid.UUID `json:"user_id" validate:"required"`
		Code   string    `json:"code" validate:"required"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	res, err := h.svc.LoginMFA(c.Context(), req.UserID, req.Code)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "MFA login failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Login successful", res)
}

func (h *AuthHandler) Refresh(c fiber.Ctx) error {
	type request struct {
		RefreshToken string `json:"refresh_token" validate:"required"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	res, err := h.svc.RefreshToken(c.Context(), req.RefreshToken)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Token refresh failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Token refreshed", res)
}

func (h *AuthHandler) GetMyPermissions(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	perms, err := h.svc.GetPermissions(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get permissions", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Permissions retrieved", perms)
}

func (h *AuthHandler) GetMyMenu(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	menu, err := h.svc.GetMenu(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get menu", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Menu retrieved", menu)
}

func (h *AuthHandler) GetMyProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	profile, err := h.svc.GetProfile(c.Context(), userID)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get profile", err.Error())
	}
	return response.Success(c, fiber.StatusOK, "Profile retrieved", profile)
}

func (h *AuthHandler) ChangePassword(c fiber.Ctx) error {
	userID := c.Locals("user_id").(uuid.UUID)
	type request struct {
		OldPassword string `json:"old_password" validate:"required"`
		NewPassword string `json:"new_password" validate:"required,min=6"`
	}
	req := new(request)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	if errs := utils.ValidateStruct(req); len(errs) > 0 {
		return response.Error(c, fiber.StatusBadRequest, "Validation failed", utils.FormatValidationErrors(errs))
	}

	err := h.svc.ChangePassword(c.Context(), userID, req.OldPassword, req.NewPassword)
	if err != nil {
		return response.Error(c, fiber.StatusUnauthorized, "Change password failed", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Password changed successfully", nil)
}
