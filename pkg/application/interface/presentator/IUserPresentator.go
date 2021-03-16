package presentator

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
)

type IUserPresentator interface {
	Serialize(entity *userModel.UserData) (userPresentator.UserOutputData, error)
}