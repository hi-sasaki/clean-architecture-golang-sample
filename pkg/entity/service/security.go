package service

import (
	"context"

	"github.com/hi-sasaki/clean-architecture-golang-sample/pkg/entity/repository"
)

type Security struct {
	Security repository.Security
}

func (s Security) GenerateCustomToken(ctx context.Context, id string) (string, error) {
	return s.Security.GenerateCustomToken(ctx, id)
}
