package cmd

import (
	"github.com/spf13/cobra"
)

// NewRootCommand create root xcloud command
func NewRootCommand() *cobra.Command {
	var command = &cobra.Command {
		Use: "xcloud",
		Short: "xcloud is cli to XCloud",
		Run: func(cmd *cobra.Command, args []string) {
			cmd.HelpFunc()(cmd, args)
		},
	}

	command.AddCommand(NewCreateCommand())
	command.AddCommand(NewDeleteCommand())
	command.AddCommand(NewGetCommand())
	command.AddCommand(NewEditCommand())

	return command
}
