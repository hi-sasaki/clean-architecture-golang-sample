package cobra

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	_interface "github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/command/interface"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func NewCommand(controller *userController.UserJsonController) _interface.Command {
	return &cobra.Command{
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
}
