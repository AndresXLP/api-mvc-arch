package services

import (
	"context"

	"github.com/andresxlp/api-mvc-arch/db/repo"
	"github.com/andresxlp/api-mvc-arch/dto"
)

type UserServices interface {
	CreateUser(ctx context.Context, user dto.User) error
}

type userServiceImpl struct {
	repo repo.UserRepository
}

func NewUserService(repo repo.UserRepository) UserServices {
	return &userServiceImpl{repo}
}

func (u *userServiceImpl) CreateUser(ctx context.Context, user dto.User) error {
	err := u.repo.CreateUser(ctx, user)
	if err != nil {
		return err
	}
	return nil
}
