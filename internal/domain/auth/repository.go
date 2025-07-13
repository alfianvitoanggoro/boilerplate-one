package auth

import (
	"boilerplate-one/internal/domain/user"
	"errors"
)

type Repository interface {
	FindByEmail(email string) (*user.User, error)
}

type repository struct {
	users []user.User // dummy data, ganti dengan DB jika perlu
}

func NewRepository() Repository {
	return &repository{
		users: []user.User{
			{ID: 1, Name: "Admin", Email: "admin@mail.com", RoleID: 1},
			{ID: 2, Name: "User", Email: "user@mail.com", RoleID: 2},
		},
	}
}

func (r *repository) FindByEmail(email string) (*user.User, error) {
	for _, u := range r.users {
		if u.Email == email {
			return &u, nil
		}
	}
	return nil, errors.New("user not found")
}
