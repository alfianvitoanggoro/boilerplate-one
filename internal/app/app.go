package app

import (
	"boilerplate-one/internal/app/factory"
	"boilerplate-one/internal/app/middleware"
	"boilerplate-one/internal/app/router"
	"boilerplate-one/internal/config"

	"github.com/gofiber/fiber/v2"
)

func BuildApp(config *config.Config) *fiber.App {
	app := fiber.New()

	middleware.RegisterMiddleware(app)

	fac := factory.NewFactory(config) // Inisialisasi semua dependency

	router.InitRouter(app, fac, config) // Daftarkan semua route

	return app
}
