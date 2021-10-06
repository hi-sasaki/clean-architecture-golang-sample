package service

import (
	"context"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/repository"
)

type User struct {
	User repository.User
}

func (u *User) GetByID(ctx context.Context, id string) (*model.User, error) {
	return u.User.GetByID(ctx, id)
}

func (u *User) CreateUser(ctx context.Context, user *model.User) (*model.User, error) {
	return u.User.CreateUser(ctx, user)
}
