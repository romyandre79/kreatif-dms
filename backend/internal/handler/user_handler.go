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
