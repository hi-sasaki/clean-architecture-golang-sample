package userModel

import (
	"github.com/google/uuid"
)

type idValue string

func newIdValue() *idValue {
	id := idValue(uuid.New().String())

	return &id
}

func (rcv idValue) value() string {
	return string(rcv)
}
