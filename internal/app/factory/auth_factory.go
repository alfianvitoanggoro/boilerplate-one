package factory

import (
	"boilerplate-one/internal/domain/auth"

	"gorm.io/gorm"
)

func NewAuthFactory(db *gorm.DB) *auth.AuthHandler {
	repo := auth.NewRepository() // dummy, ganti dengan repo DB jika perlu
	service := auth.NewService(repo)
	handler := auth.NewAuthHandler(service)
	return handler
}
