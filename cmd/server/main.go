package main

import (
	"boilerplate-one/internal/app"
	"boilerplate-one/internal/config"
	"boilerplate-one/pkg/logger"
)

func main() {
	config := config.Load()

	fiberApp := app.BuildApp(config)

	logger.Infof("%s Server running on %s:%s", config.App.Name, config.App.Env, config.App.Port)
	fiberApp.Listen(":" + config.App.Port)
}
