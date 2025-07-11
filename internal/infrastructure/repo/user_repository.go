package repo

import (
	"context"

	"boilerplate-one/internal/domain/user"

	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) user.Repository {
	return &userRepository{db: db}
}

// Create user (sudah ada)
func (r *userRepository) Create(ctx context.Context, u *user.User) error {
	return r.db.WithContext(ctx).Create(u).Error
}

// Find by ID (sudah ada)
func (r *userRepository) FindByID(ctx context.Context, id uint) (*user.User, error) {
	var u user.User
	if err := r.db.WithContext(ctx).First(&u, id).Error; err != nil {
		return nil, err
	}
	return &u, nil
}

// Find all (sudah ada)
func (r *userRepository) FindAll(ctx context.Context) ([]user.User, error) {
	var users []user.User
	if err := r.db.WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

// Update user
func (r *userRepository) Update(ctx context.Context, id uint, u *user.User) (*user.User, error) {
	var existing user.User
	if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
		return nil, err
	}
	if err := r.db.WithContext(ctx).Model(&existing).Updates(u).Error; err != nil {
		return nil, err
	}
	return &existing, nil
}

// Delete user
func (r *userRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&user.User{}, id).Error
}
