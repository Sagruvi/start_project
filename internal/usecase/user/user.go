package usecase_user

import (
	"context"
	model_user "start/internal/model/user"
)

type UserRepository interface {
	GetUser(ctx context.Context, id int) (model_user.User, error)
	CreateUser(ctx context.Context, user model_user.User) (model_user.User, error)
	DeleteUser(ctx context.Context, id int) error
	UpdateUser(ctx context.Context, user model_user.User) (model_user.User, error)
}

type UserUsecase struct {
	repo UserRepository
}

func New(repo UserRepository) *UserUsecase {
	return &UserUsecase{repo: repo}
}

func (u *UserUsecase) GetUser(ctx context.Context, id int) (model_user.User, error) {
	return u.repo.GetUser(ctx, id)
}

func (u *UserUsecase) CreateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	return u.repo.CreateUser(ctx, user)
}

func (u *UserUsecase) DeleteUser(ctx context.Context, id int) error {
	return u.repo.DeleteUser(ctx, id)
}

func (u *UserUsecase) UpdateUser(ctx context.Context, user model_user.User) (model_user.User, error) {
	return u.repo.UpdateUser(ctx, user)
}
