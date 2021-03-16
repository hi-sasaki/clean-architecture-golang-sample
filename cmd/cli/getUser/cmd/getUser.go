package cmd

import (
	"fmt"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
)

var controller *userController.UserJsonController
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		log.SetLevel(log.Level(5))
		log.Print(controller)
		result, err := controller.GetUser()
		if err != nil {
			log.Fatal(err)

		}
		log.Print(result)
	},
}


func Execute(ctl *userController.UserJsonController) {
	if ctl == nil {
		log.Fatal("controller is nil")
	}
	controller = ctl

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

