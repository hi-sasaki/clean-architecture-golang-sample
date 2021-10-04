package usecase

import (
	"context"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/admin/inout"
)

type User struct {
	Service *service.User
}

func ProviderUser(service service.User) *User {
	return &User{
		Service: &service,
	}
}

func (u *User) GetByID(ctx context.Context, id string) (*inout.User, error) {

	res, err := u.Service.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}
	user := &inout.User{
		ID:        res.ID,
		FirstName: res.FirstName,
		LastName:  res.LastName,
	}
	return user, nil
}
