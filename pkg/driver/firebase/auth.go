package firebase

import (
	"context"

	"firebase.google.com/go/auth"
)

type Oparator struct {
	Client *auth.Client
}

func Provider(client *auth.Client) *Oparator {
	return &Oparator{Client: client}
}
func (o *Oparator) GenerateCustomToken(ctx context.Context, userID string) (string, error) {
	return o.Client.CustomToken(ctx, userID)
}
