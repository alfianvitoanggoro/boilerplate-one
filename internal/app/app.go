package app

import (
	"boilerplate-one/internal/config"

	"github.com/gofiber/fiber/v2"
)

func BuildApp(config *config.Config) *fiber.App {
	app := fiber.New()

	RegisterMiddleware(app)

	factory := NewFactory(config) // Inisialisasi semua dependency

	InitRouter(app, factory, config) // Daftarkan semua route

	return app
}
