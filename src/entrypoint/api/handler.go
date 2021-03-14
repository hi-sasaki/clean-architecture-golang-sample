package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/infrastracture/handler/chi"
	"net/http"
)

func handler(controller *userController.UserJsonController) http.Handler {
	return chi.NewHandler().Create(controller)
}

