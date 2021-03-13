package presentator

import (
	"go-study-sample/src/lib/adaptor/presentator/userPresentator"
	"go-study-sample/src/lib/domain/model/userModel"
)

type IUserPresentator interface {
	Serialize(entity *userModel.UserData) (userPresentator.UserOutputData, error)
}