package userController

import (
	"github.com/pkg/errors"
	"go-study-sample/src/lib/adaptor/presentator/userPresentator"
	"go-study-sample/src/lib/application/usecase"
)

type UserJsonController struct {
	useCase usecase.GetUserInteractor
}

func (rcv UserJsonController) GetUser() (userPresentator.UserOutputData, error) {
	result, err := rcv.useCase.Handle()
	if err != nil {
		return "", errors.Wrap(err, "Wrap in GetUser UserJsonController: ")
	}

	return result, nil

}

func New(useCase usecase.GetUserInteractor) (*UserJsonController, error) {
	controller := UserJsonController{useCase}
	return &controller, nil
}
