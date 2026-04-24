package middleware

import (
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/google/uuid"
	"github.com/kreatif/dms-backend/internal/auth"
	"github.com/stretchr/testify/assert"
)

func TestAuthMiddleware(t *testing.T) {
	app := fiber.New()
	secret := "test-secret"
	
	app.Use("/protected", AuthMiddleware(secret))
	app.Get("/protected", func(c fiber.Ctx) error {
		return c.SendString("success")
	})

	t.Run("No Header", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/protected", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusUnauthorized, resp.StatusCode)
	})

	t.Run("Valid Token", func(t *testing.T) {
		token, _ := auth.GenerateToken(auth.TokenPayload{
			UserID:       uuid.New(),
			Role:         "user",
			DepartmentID: uuid.New(),
		}, secret, 15*time.Minute)

		req := httptest.NewRequest("GET", "/protected", nil)
		req.Header.Set("Authorization", "Bearer "+token)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})
}

func TestRoleMiddleware(t *testing.T) {
	app := fiber.New()
	
	// Mock Auth Middleware by setting locals
	app.Use("/admin", func(c fiber.Ctx) error {
		c.Locals("user_role", "admin")
		return c.Next()
	})
	app.Get("/admin", RoleMiddleware("admin"), func(c fiber.Ctx) error {
		return c.SendString("admin-area")
	})

	app.Use("/user-area", func(c fiber.Ctx) error {
		c.Locals("user_role", "user")
		return c.Next()
	})
	app.Get("/user-area", RoleMiddleware("admin"), func(c fiber.Ctx) error {
		return c.SendString("admin-only")
	})

	t.Run("Authorized Role", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/admin", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusOK, resp.StatusCode)
	})

	t.Run("Unauthorized Role", func(t *testing.T) {
		req := httptest.NewRequest("GET", "/user-area", nil)
		resp, _ := app.Test(req)
		assert.Equal(t, fiber.StatusForbidden, resp.StatusCode)
	})
}
