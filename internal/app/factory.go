package app

import (
	"boilerplate-one/internal/config"
	"boilerplate-one/internal/domain/user"
	"boilerplate-one/internal/infrastructure/db"
	"boilerplate-one/internal/infrastructure/migration"
	"boilerplate-one/internal/infrastructure/repo"
	"os"

	"gorm.io/gorm"
)

type Factory struct {
	DB *gorm.DB

	// Domain
	UserHandler *user.UserHandler
}

func NewFactory(config *config.Config) *Factory {
	database := db.Connect(config.DB)

	// Migration database if MIGRATE env is true
	args := os.Args
	if len(args) > 1 && args[1] == "migrate" {
		migration.Run(database)
		os.Exit(0)
	}

	// Inisialisasi domain user
	userRepo := repo.NewUserRepository(database)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	return &Factory{
		DB:          database,
		UserHandler: userHandler,
	}
}
