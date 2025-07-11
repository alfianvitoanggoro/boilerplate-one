package response

import (
	"github.com/gofiber/fiber/v2"
)

type SuccessResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type ErrorResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Error   any    `json:"error,omitempty"`
}

// SendSuccess sends a success response
func SendSuccess(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusOK).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// SendCreated sends a created success response
func SendCreated(c *fiber.Ctx, message string, data interface{}) error {
	return c.Status(fiber.StatusCreated).JSON(SuccessResponse{
		Status:  "success",
		Message: message,
		Data:    data,
	})
}

// SendError sends an error response
func SendError(c *fiber.Ctx, code int, message string, err interface{}) error {
	return c.Status(code).JSON(ErrorResponse{
		Status:  "error",
		Message: message,
		Error:   err,
	})
}
