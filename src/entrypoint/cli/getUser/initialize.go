package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/lib/adaptor/controller/userController"
	log "github.com/sirupsen/logrus"
	)

func initializeController() *userController.UserJsonController  {
	controller,err := InitializeUserController()
	if err != nil {
		log.Fatal(err)
	}
	return controller


}