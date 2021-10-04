package usecase

import (
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
)

type User struct {
	Service *service.User
}

func ProviderUser(service service.User) *User {
	return &User{
		Service: &service,
	}
}

func (u *User) GetByID() (*model.User, error) {
	u.Service.GetByID()
	return nil, nil
}
