package repository

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/domain/model/userModel"
)

type IUserRepository interface {
	GetUser() (*userModel.UserData, error)
}
