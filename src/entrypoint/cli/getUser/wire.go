//+ wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/repository/userRepository"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/application/interface/presentator"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/application/interface/repository"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/application/usecase"
)

func InitializeUserController() (*userController.UserJsonController,error) {

	wire.Build(
		userController.New,
		usecase.New,
		userRepository.New,
		userPresentator.New,
		wire.Bind(new(repository.IUserRepository), new(*userRepository.UserInMemoryRepository)),
		wire.Bind(new(presentator.IUserPresentator), new(*userPresentator.UserJsonPresentator)),
	)
	return nil, nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}