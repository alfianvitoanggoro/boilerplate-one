package factory

import (
	"boilerplate-one/internal/domain/user"
	"boilerplate-one/internal/infrastructure/repo"

	"gorm.io/gorm"
)

func NewUserFactory(database *gorm.DB) (*user.UserHandler, user.Repository) {
	userRepo := repo.NewUserRepository(database)
	userService := user.NewService(userRepo)
	userHandler := user.NewUserHandler(userService)

	return userHandler, userRepo
}
