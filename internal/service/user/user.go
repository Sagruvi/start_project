package service_user

import (
	"context"
	model_user "start/internal/model/user"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id int) (model_user.User, error)
	CreateUser(ctx context.Context, user model_user.User) (model_user.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user model_user.User) (model_user.User, error)
}

type service struct {
	uc UserUsecase
}

func New(uc UserUsecase) *service {
	return &service{uc: uc}
}

func (s *service) GetUser(ctx context.Context, id int) (model_user.User, error) {
	return s.uc.GetUser(ctx, id)
}

func (s *service) CreateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	return s.uc.CreateUser(ctx, user)
}

func (s *service) DeleteUser(ctx context.Context, id int) error {
	return s.uc.DeleteUser(ctx, id)
}

func (s *service) UpdateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	return s.uc.UpdateUser(ctx, user)
}
