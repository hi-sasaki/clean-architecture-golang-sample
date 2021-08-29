package cobra

import (
	"github.com/rikodao/clean-architecture-golang-sample/pkg/adaptor/controller/userController"
	"github.com/rikodao/clean-architecture-golang-sample/pkg/infrastracture/command"
	"github.com/spf13/cobra"
)

func CommandGetUser(controller *userController.UserCommandController) command.ICommand {
	return &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: controller.CommandGetUser,
	}
}

func CommandGetUserById(controller *userController.UserCommandController) command.ICommand {
	return &cobra.Command{
		Use:   "hugo",
		Short: "Hugo is a very fast static site generator",
		Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at http://hugo.spf13.com`,
		Run: controller.CommandGetUserById,
	}
}
