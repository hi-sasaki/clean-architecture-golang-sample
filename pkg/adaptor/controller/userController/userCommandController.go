package userController

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/controller"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/usecase"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

type UserCommandController struct {
	getUserInteractor     usecase.GetUserInteractor
	getUserByIdInteractor usecase.GetUserByIdInteractor
}

func (rcv UserCommandController) CommandGetUser(cmd *cobra.Command, args []string) {
	log.SetLevel(log.Level(5))
	result, err := rcv.getUserInteractor.Handle()
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}

func (rcv UserCommandController) CommandGetUserById(cmd *cobra.Command, args []string) {

	log.Print(args[0])
	log.SetLevel(log.Level(5))
	command := controller.GetUserByIdCommand{Id: args[0]}
	result, err := rcv.getUserByIdInteractor.Handle(command)
	if err != nil {
		log.Fatal(err)
	}
	log.Print(result)
}

func NewComanndController(getUserInteractor *usecase.GetUserInteractor, getUserByIdInteractor *usecase.GetUserByIdInteractor) *UserCommandController {
	controller := UserCommandController{*getUserInteractor, *getUserByIdInteractor}
	return &controller
}
