package presentator

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/domain/model/userModel"
)

type IUserPresentator interface {
	Serialize(entity *userModel.UserData) (userPresentator.UserOutputData, error)
}