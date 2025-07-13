package auth

import (
	"boilerplate-one/pkg/response"

	"github.com/gofiber/fiber/v2"
)

type AuthHandler struct {
	service Service
}

func NewAuthHandler(service Service) *AuthHandler {
	return &AuthHandler{service: service}
}

func (h *AuthHandler) Login(c *fiber.Ctx) error {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := c.BodyParser(&req); err != nil {
		return response.SendError(c, fiber.StatusBadRequest, "Invalid request body", err)
	}
	token, err := h.service.Login(req.Email, req.Password)
	if err != nil {
		return response.SendError(c, fiber.StatusUnauthorized, "Login failed", err)
	}

	return response.SendSuccess(c, "User logged in successfully", token)
}
