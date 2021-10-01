package repository

import "context"

type Security interface {
	GenerateCustomToken(ctx context.Context, userID string) (string, error)
}
