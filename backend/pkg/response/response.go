package response

import "github.com/gofiber/fiber/v3"

type APIResponse struct {
	Success bool        `json:"success"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

func Success(c fiber.Ctx, status int, message string, data interface{}) error {
	return c.Status(status).JSON(APIResponse{
		Success: true,
		Message: message,
		Data:    data,
	})
}

func Error(c fiber.Ctx, status int, message string, err string) error {
	return c.Status(status).JSON(APIResponse{
		Success: false,
		Message: message,
		Error:   err,
	})
}
