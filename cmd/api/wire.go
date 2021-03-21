//+ wireinject

package main

import (
	"github.com/google/wire"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/presentator/userPresentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/repository/userRepository"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/presentator"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/interface/repository"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/application/usecase"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/handler/chi"
	"net/http"
)

func InitializeRouter() (http.Handler, error) {

	wire.Build(
		chi.NewHandler,
		userController.New,
		usecase.NewGetUserInteractor,
		usecase.NewGetUserByIdInteractor,
		userRepository.New,
		userPresentator.New,
		wire.Bind(new(repository.IUserRepository), new(*userRepository.UserInMemoryRepository)),
		wire.Bind(new(presentator.IUserPresentator), new(*userPresentator.UserJsonPresentator)),
	)
	return nil, nil // wireはこの関数の戻り値を無視するので、nilを返せばよい
}
