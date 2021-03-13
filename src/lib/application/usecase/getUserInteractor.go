package usecase

import (
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"go-study-sample/src/lib/adaptor/presentator/userPresentator"
	"go-study-sample/src/lib/application/interface/presentator"
	"go-study-sample/src/lib/application/interface/repository"
)

type GetUserInteractor struct {
	userRepository  repository.IUserRepository
	userPresentator presentator.IUserPresentator
}

func (rcv *GetUserInteractor) Handle() (userPresentator.UserOutputData, error) {
	log.Debug("GetUserInteractor Handle start")

	user,err := rcv.userRepository.GetUser()
	if err != nil {
		return "", errors.Wrap(err, "Wrap in Handle GetUserInteractor: ")
	}
	result,err := rcv.userPresentator.Serialize(user)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in Handle GetUserInteractor: ")
	}

	log.WithFields(log.Fields{
		"result": result,
	}).Debug("GetUserInteractor Handle end")

	return result,nil
}

func New(userRepository repository.IUserRepository, userPresentator presentator.IUserPresentator) (*GetUserInteractor, error) {
	useCace := GetUserInteractor{userRepository, userPresentator}
	return &useCace, nil
}