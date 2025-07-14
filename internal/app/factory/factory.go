package factory

import (
	"boilerplate-one/internal/config"
	"boilerplate-one/internal/domain/auth"
	"boilerplate-one/internal/domain/user"

	"boilerplate-one/internal/infrastructure/db"
	"boilerplate-one/internal/infrastructure/migration"
	"os"

	"gorm.io/gorm"
)

type Factory struct {
	DB *gorm.DB

	// Repositories
	UserRepo user.Repository

	// Domain
	UserHandler *user.UserHandler
	AuthHandler *auth.AuthHandler
}

func NewFactory(config *config.Config) *Factory {
	database := db.Connect(config.DB)

	// Migration database if MIGRATE env is true
	args := os.Args
	if len(args) > 1 && args[1] == "migrate" {
		migration.Run(database)
		os.Exit(0)
	}

	userHandler, userRepo := NewUserFactory(database)
	authHandler := NewAuthFactory(database)
	return &Factory{
		DB:          database,
		UserRepo:    userRepo,
		UserHandler: userHandler,
		AuthHandler: authHandler,
	}
}
