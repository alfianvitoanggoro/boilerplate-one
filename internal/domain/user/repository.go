package user

import "context"

type Repository interface {
	Create(ctx context.Context, u *User) error
	FindByID(ctx context.Context, id uint) (*User, error)
	FindAll(ctx context.Context) ([]User, error)
	Update(ctx context.Context, id uint, u *User) (*User, error)
	Delete(ctx context.Context, id uint) error
}
