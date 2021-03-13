package repository

import (
	"go-study-sample/src/lib/domain/model/userModel"
)

type IUserRepository interface {
	GetUser() (*userModel.UserData, error)
}
