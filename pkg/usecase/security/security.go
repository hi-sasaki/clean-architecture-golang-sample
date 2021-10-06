package security

import (
	"context"
	"errors"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/service"
)

type Security struct {
	Service     *service.Security
	UserService *service.User
}

func (s *Security) Login(ctx context.Context, userID, pass string) (string, error) {

	user, err := s.UserService.GetByID(ctx, userID)
	if err != nil {
		return "", err
	}
	if user.Password != pass {
		return "", errors.New("invalid password")
	}

	return s.Service.GenerateCustomToken(ctx, userID)
}
