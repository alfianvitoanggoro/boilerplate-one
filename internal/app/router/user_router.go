package router

import (
	"boilerplate-one/internal/domain/user"

	"github.com/gofiber/fiber/v2"
)

// userRouter is a function that registers the user routes
func publicUserRouter(router fiber.Router, handler *user.UserHandler) {
	router.Post("/", handler.Register) // register tanpa JWT
	// Endpoint lain pakai JWT
	router.Get("/", handler.GetUsers)
	router.Get("/:id", handler.GetUserByID)
	router.Put("/:id", handler.UpdateUser)
	router.Delete("/:id", handler.DeleteUser)
}

func privateUserRouter(router fiber.Router, handler *user.UserHandler) {
	router.Get("/", handler.GetUsers)
	router.Get("/:id", handler.GetUserByID)
	router.Put("/:id", handler.UpdateUser)
	router.Delete("/:id", handler.DeleteUser)
}
