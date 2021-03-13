package cmd

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"go-study-sample/src/lib/adaptor/controller/userController"
	"go-study-sample/src/lib/adaptor/presentator/userPresentator"
	"go-study-sample/src/lib/adaptor/repository/userRepository"
	"go-study-sample/src/lib/application/usecase"
	"os"
)
var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
	Run: func(cmd *cobra.Command, args []string) {

		log.SetLevel(log.Level(5))
		repository, err := userRepository.New()
		if err != nil {
			log.Fatal(err)
		}
		presentator, err := userPresentator.New()
		if err != nil {
			log.Fatal(err)

			usecase, err := usecase.New(repository, presentator)
			if err != nil {
				log.Fatal(err)
			}

			controller, err := userController.New(*usecase)
			if err != nil {
				log.Fatal(err)
			}
			result, err := controller.GetUser()
			log.Print(result)
			log.Print(&result)

		}
	},
}


func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

