package userModel

import (
	"github.com/google/uuid"
)

type idValue string

func newIdValue() (*idValue, error) {
	id := idValue(uuid.New().String())

	return &id, nil
}

func (rcv idValue) value() string {
	return string(rcv)
}
