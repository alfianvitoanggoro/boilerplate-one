package router

import (
	"boilerplate-one/internal/domain/auth"

	"github.com/gofiber/fiber/v2"
)

// authRouter is a function that registers the auth routes
func authRouter(router fiber.Router, handler *auth.AuthHandler) {
	router.Post("/login", handler.Login) // Login user without JWT
}
