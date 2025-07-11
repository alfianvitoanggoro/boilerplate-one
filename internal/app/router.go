package app

import (
	"boilerplate-one/internal/config"
	"boilerplate-one/internal/domain/user"
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func InitRouter(app *fiber.App, f *Factory, config *config.Config) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString(fmt.Sprintf("Hello, Welcome to the %s API", config.App.Name))
	})

	api := app.Group("/api")

	// Register user router
	userGroup := api.Group("/users")
	user.UserRouter(userGroup, f.UserHandler)
}
