package user

import "github.com/gofiber/fiber/v2"

// UserRouter is a function that registers the user routes
func UserRouter(router fiber.Router, handler *UserHandler) {
	router.Post("/", handler.Register)
	router.Get("/", handler.GetUsers)
	router.Get("/:id", handler.GetUserByID)
	router.Put("/:id", handler.UpdateUser)
	router.Delete("/:id", handler.DeleteUser)
}
