package router

import (
	"boilerplate-one/internal/app/middleware"
	"boilerplate-one/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

// userRouter is a function that registers the user routes
func userRouter(router fiber.Router, handler *user.UserHandler) {
	router.Post("/", handler.Register) // register tanpa JWT
	// Endpoint lain pakai JWT
	router.Get("/", middleware.JWTMiddleware(), handler.GetUsers)
	router.Get("/:id", middleware.JWTMiddleware(), handler.GetUserByID)
	router.Put("/:id", middleware.JWTMiddleware(), handler.UpdateUser)
	router.Delete("/:id", middleware.JWTMiddleware(), handler.DeleteUser)
}
