package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/repository/userRepository"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/presentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/repository"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/usecase"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/command"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/command/cobra"
	"log"
)

func NewCommand() command.ICommand {
	consroller := newUserController()
	return cobra.CommandGetUser(consroller)
}

func newUserController() *userController.UserCommandController {
	controller := userController.NewComanndController(newGetUserInteractor(), newGetUserByIdInteractor())
	return controller
}

func newUserRepository() repository.IUserRepository {
	repository, err := userRepository.New()
	if err != nil {
		log.Fatal("Error newUserRepository")
	}
	return repository
}

func newUserPresentator() presentator.IUserPresentator {
	presentor, err := userPresentator.New()
	if err != nil {
		log.Fatal("Error newUserPresentator")
	}
	return presentor
}

func newGetUserInteractor() *usecase.GetUserInteractor {
	usecase, err := usecase.NewGetUserInteractor(newUserRepository(), newUserPresentator())
	if err != nil {
		log.Fatal("Error newGetUserInteractor")
	}
	return usecase
}

func newGetUserByIdInteractor() *usecase.GetUserByIdInteractor {
	usecase, err := usecase.NewGetUserByIdInteractor(newUserRepository(), newUserPresentator())
	if err != nil {
		log.Fatal("Error newGetUserInteractor")
	}
	return usecase
}
