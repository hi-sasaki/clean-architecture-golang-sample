package userController

import (
	"github.com/pkg/errors"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/controller"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/usecase"
)

type UserJsonController struct {
	getUserInteractor usecase.GetUserInteractor
	getUserByIdInteractor usecase.GetUserByIdInteractor
}

func (rcv UserJsonController) GetUser() (userPresentator.UserOutputData, error) {
	result, err := rcv.getUserInteractor.Handle()
	if err != nil {
		return "", errors.Wrap(err, "Wrap in GetUser UserJsonController: ")
	}

	return result, nil

}

func (rcv UserJsonController) GetUserById(id string) (userPresentator.UserOutputData, error) {
	command := controller.GetUserByIdCommand{Id: id}
	result, err := rcv.getUserByIdInteractor.Handle(command)
	if err != nil {
		return "", errors.Wrap(err, "Wrap in GetUser UserJsonController: ")
	}

	return result, nil

}


func New(getUserInteractor *usecase.GetUserInteractor,getUserByIdInteractor *usecase.GetUserByIdInteractor) (*UserJsonController, error) {
	controller := UserJsonController{*getUserInteractor,*getUserByIdInteractor}
	return &controller, nil
}
