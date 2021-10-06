package admin

import (
	"context"

	"github.com/google/uuid"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/model"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/usecase/inout"
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
func (u *User) CreateUser(ctx context.Context, input *inout.User) (*inout.User, error) {

	user := &model.User{
		FirstName: input.FirstName,
		LastName:  input.LastName,
		Password:  input.Password,
		ID:        uuid.Must(uuid.NewRandom()).String(),
	}
	res, err := u.Service.CreateUser(ctx, user)
	if err != nil {
		return nil, err
	}

	return &inout.User{
		FirstName: res.FirstName,
		LastName:  res.LastName,
		ID:        res.ID,
	}, nil
}
