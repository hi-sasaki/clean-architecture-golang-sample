package userModel

import "errors"

type nameValue string

const (
	maxNameLength     = 16
	minimumNameLength = 1
)

func newNameValue(name string) (*nameValue, error) {
	if len(name) < minimumNameLength {
		return nil, errors.New("名前を入力して下さい。")
	}

	if len(name) > maxNameLength {
		return nil, errors.New("名前の上限数を超えています。")
	}

	nameValueObject := nameValue(name)
	return &nameValueObject, nil
}

func (rcv nameValue) value() string {
	return string(rcv)
}
