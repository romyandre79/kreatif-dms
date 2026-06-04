package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"github.com/kreatif/dms-backend/internal/repository"
	"github.com/kreatif/dms-backend/pkg/response"
	"golang.org/x/crypto/bcrypt"
)

type UserHandler struct {
	repo repository.Querier
}

func NewUserHandler(repo repository.Querier) *UserHandler {
	return &UserHandler{repo: repo}
}

type createUserRequest struct {
	Email        string    `json:"email" validate:"required,email"`
	Password     string    `json:"password" validate:"required,min=6"`
	FullName     string    `json:"full_name" validate:"required"`
	RoleID       int32     `json:"role_id" validate:"required"`
	DepartmentID uuid.UUID `json:"department_id"`
	AvatarUrl    string    `json:"avatar_url"`
	SignatureUrl string    `json:"signature_url"`
	IsMfaEnabled bool      `json:"is_mfa_enabled"`
}

func (h *UserHandler) Register(c fiber.Ctx) error {
	req := new(createUserRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to hash password", err.Error())
	}

	user, err := h.repo.CreateUser(c.Context(), repository.CreateUserParams{
		Email:        req.Email,
		PasswordHash: string(hashedPassword),
		FullName:     req.FullName,
		RoleID:       req.RoleID,
		DepartmentID: pgtype.UUID{Bytes: req.DepartmentID, Valid: req.DepartmentID != uuid.Nil},
		Status:       "active",
		AvatarUrl:    pgtype.Text{String: req.AvatarUrl, Valid: req.AvatarUrl != ""},
		SignatureUrl: pgtype.Text{String: req.SignatureUrl, Valid: req.SignatureUrl != ""},
		IsMfaEnabled: pgtype.Bool{Bool: req.IsMfaEnabled, Valid: true},
	})

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to create user", err.Error())
	}

	return response.Success(c, fiber.StatusCreated, "User registered successfully", user)
}

func (h *UserHandler) List(c fiber.Ctx) error {
	users, err := h.repo.ListUsers(c.Context())
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to list users", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "Users retrieved successfully", users)
}

func (h *UserHandler) Update(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	type updateRequest struct {
		Email        string    `json:"email"`
		FullName     string    `json:"full_name"`
		RoleID       int32     `json:"role_id"`
		DepartmentID uuid.UUID `json:"department_id"`
		Status       string    `json:"status"`
		AvatarUrl    string    `json:"avatar_url"`
		SignatureUrl string    `json:"signature_url"`
		IsMfaEnabled bool      `json:"is_mfa_enabled"`
	}

	req := new(updateRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	user, err := h.repo.UpdateUser(c.Context(), repository.UpdateUserParams{
		ID:           id,
		Email:        req.Email,
		FullName:     req.FullName,
		RoleID:       req.RoleID,
		DepartmentID: pgtype.UUID{Bytes: req.DepartmentID, Valid: req.DepartmentID != uuid.Nil},
		Status:       req.Status,
		AvatarUrl:    pgtype.Text{String: req.AvatarUrl, Valid: req.AvatarUrl != ""},
		SignatureUrl: pgtype.Text{String: req.SignatureUrl, Valid: req.SignatureUrl != ""},
		IsMfaEnabled: pgtype.Bool{Bool: req.IsMfaEnabled, Valid: true},
	})

	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to update user", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User updated successfully", user)
}

func (h *UserHandler) Delete(c fiber.Ctx) error {
	id, err := uuid.Parse(c.Params("id"))
	if err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid user ID", err.Error())
	}

	err = h.repo.DeleteUser(c.Context(), id)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to delete user", err.Error())
	}

	return response.Success(c, fiber.StatusOK, "User deleted successfully", nil)
}

func (h *UserHandler) GetHierarchy(c fiber.Ctx) error {
	ctx := c.Context()
	
	companies, err := h.repo.GetCompanies(ctx)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get companies", err.Error())
	}

	branches, err := h.repo.GetBranches(ctx)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get branches", err.Error())
	}

	departments, err := h.repo.GetDepartments(ctx)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get departments", err.Error())
	}

	users, err := h.repo.GetUsersForHierarchy(ctx)
	if err != nil {
		return response.Error(c, fiber.StatusInternalServerError, "Failed to get hierarchy users", err.Error())
	}

	result := map[string]interface{}{
		"companies":   companies,
		"branches":    branches,
		"departments": departments,
		"users":       users,
	}

	return response.Success(c, fiber.StatusOK, "Organization structure retrieved successfully", result)
}

type updateHierarchyRequest struct {
	UserUpdates []struct {
		UserID       uuid.UUID  `json:"user_id" validate:"required"`
		DepartmentID *uuid.UUID `json:"department_id"`
	} `json:"user_updates"`
	DepartmentUpdates []struct {
		DepartmentID uuid.UUID  `json:"department_id" validate:"required"`
		HeadID       *uuid.UUID `json:"head_id"`
	} `json:"department_updates"`
}

func (h *UserHandler) UpdateHierarchy(c fiber.Ctx) error {
	req := new(updateHierarchyRequest)
	if err := c.Bind().JSON(req); err != nil {
		return response.Error(c, fiber.StatusBadRequest, "Invalid request body", err.Error())
	}

	// Iterate and update users
	for _, update := range req.UserUpdates {
		var deptID pgtype.UUID
		if update.DepartmentID != nil {
			deptID = pgtype.UUID{Bytes: *update.DepartmentID, Valid: true}
		} else {
			deptID = pgtype.UUID{Valid: false}
		}

		err := h.repo.UpdateUserDepartment(c.Context(), repository.UpdateUserDepartmentParams{
			DepartmentID: deptID,
			ID:           update.UserID,
		})
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update hierarchy", err.Error())
		}
	}

	// Iterate and update departments
	for _, update := range req.DepartmentUpdates {
		var headID pgtype.UUID
		if update.HeadID != nil {
			headID = pgtype.UUID{Bytes: *update.HeadID, Valid: true}
		} else {
			headID = pgtype.UUID{Valid: false}
		}

		err := h.repo.UpdateDepartmentHead(c.Context(), repository.UpdateDepartmentHeadParams{
			HeadID: headID,
			ID:     update.DepartmentID,
		})
		if err != nil {
			return response.Error(c, fiber.StatusInternalServerError, "Failed to update department head", err.Error())
		}
	}

	return response.Success(c, fiber.StatusOK, "Hierarchy updated successfully", nil)
}
