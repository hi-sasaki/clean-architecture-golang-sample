package usecase

import (
	"github.com/pkg/errors"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/presentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/repository"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/controller"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/domain/model/userModel"
	log "github.com/sirupsen/logrus"
)

type GetUserByIdInteractor struct {
	userRepository  repository.IUserRepository
	userPresentator presentator.IUserPresentator
}

func (rcv *GetUserByIdInteractor) Handle(command controller.GetUserByIdCommand) (userPresentator.UserOutputData, error) {
	log.Debug("GetUserInteractor Handle start")

	user, err := rcv.userRepository.GetUGetUserById(command.Id)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in Handle GetUserInteractor: ")
	}
	userData := userModel.Data(*user)

	result, err := rcv.userPresentator.Serialize(userData)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in Handle GetUserInteractor: ")
	}

	log.WithFields(log.Fields{
		"result": result,
	}).Debug("GetUserInteractor Handle end")

	return result, nil
}

func NewGetUserByIdInteractor(userRepository repository.IUserRepository, userPresentator presentator.IUserPresentator) (*GetUserByIdInteractor, error) {
	useCace := GetUserByIdInteractor{userRepository, userPresentator}
	return &useCace, nil
}
