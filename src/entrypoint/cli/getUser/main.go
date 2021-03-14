package main

import (
	"github.com/rikodao/clean-architecture-golang-sample/src/entrypoint/cli/getUser/cmd"
	log "github.com/sirupsen/logrus"
)

func main() {
	controller, err := InitializeUserController()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute(controller)
}