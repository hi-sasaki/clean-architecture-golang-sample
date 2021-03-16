package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/handler/chi"
	"net/http"
)

func handler(controller *userController.UserJsonController) http.Handler {
	return chi.NewHandler().Create(controller)
}

