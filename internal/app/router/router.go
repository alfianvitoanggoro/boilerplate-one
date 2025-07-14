package router

import (
	"boilerplate-one/internal/app/factory"
	"boilerplate-one/internal/app/middleware"
	"boilerplate-one/internal/config"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, f *factory.Factory, config *config.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Welcome to the %s API", config.App.Name))
	})

	api := app.Group("/api")

	// Register auth router
	authGroup := api.Group("/auth")
	authRouter(authGroup, f.AuthHandler)

	// 🔓 Public user routes (tanpa JWT)
	publicUserGroup := api.Group("/users")
	publicUserRouter(publicUserGroup, f.UserHandler)

	// 🔐 Protected user routes (dengan JWT)
	privateUserGroup := api.Group("/users")
	privateUserGroup.Use(middleware.JWTMiddleware(f.UserRepo))
	privateUserRouter(privateUserGroup, f.UserHandler)
}
