package usecase

import (
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
)

type User struct {
	service *service.User
}

func (u *User) GetByID() (*model.User, error) {
	u.service.GetByID()
	return nil, nil
}
