package factory

import (
	"boilerplate-one/internal/domain/user"
	"boilerplate-one/internal/infrastructure/repo"

	"gorm.io/gorm"
)

func NewUserFactory(database *gorm.DB) *user.UserHandler {
	userRepo := repo.NewUserRepository(database)
	userService := user.NewService(userRepo)
	return user.NewUserHandler(userService)
}
