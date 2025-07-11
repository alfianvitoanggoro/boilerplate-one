package user

import (
	"boilerplate-one/internal/domain/user/dto"
	"context"
)

type Service interface {
	Register(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserDetailResponse, error)
	GetUserByID(ctx context.Context, id uint) (*dto.UserDetailResponse, error)
	GetAllUser(ctx context.Context) ([]dto.UserDetailResponse, error)
	UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserDetailResponse, error)
	DeleteUser(ctx context.Context, id uint) error
}

type service struct {
	repo Repository
}

func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) Register(ctx context.Context, req *dto.CreateUserRequest) (*dto.UserDetailResponse, error) {
	user := &User{
		Name:  req.Name,
		Email: req.Email,
	}
	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return &dto.UserDetailResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *service) GetUserByID(ctx context.Context, id uint) (*dto.UserDetailResponse, error) {
	user, err := s.repo.FindByID(ctx, id)
	if err != nil {
		return nil, err
	}
	return &dto.UserDetailResponse{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *service) GetAllUser(ctx context.Context) ([]dto.UserDetailResponse, error) {
	users, err := s.repo.FindAll(ctx)
	if err != nil {
		return nil, err
	}
	var resp []dto.UserDetailResponse
	for _, u := range users {
		resp = append(resp, dto.UserDetailResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		})
	}
	return resp, nil
}

func (s *service) UpdateUser(ctx context.Context, id uint, req *dto.UpdateUserRequest) (*dto.UserDetailResponse, error) {
	user := &User{}
	if req.Name != nil {
		user.Name = *req.Name
	}
	if req.Email != nil {
		user.Email = *req.Email
	}

	updated, err := s.repo.Update(ctx, id, user)
	if err != nil {
		return nil, err
	}
	return &dto.UserDetailResponse{
		ID:    updated.ID,
		Name:  updated.Name,
		Email: updated.Email,
	}, nil
}

func (s *service) DeleteUser(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}
