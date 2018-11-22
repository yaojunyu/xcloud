package cmd

import "github.com/spf13/cobra"

type DeleteOptions struct {

}

type DeleteFlags struct {
	name string
}

func NewDeleteCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "delete [NAME]",
		Short: "Delete a resource",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	deleteFlags := &DeleteFlags{}
	addDeleteFlags(command, deleteFlags)
	return command
}

func addDeleteFlags(command *cobra.Command, flags *DeleteFlags) {

	command.Flags().StringVarP(&flags.name, "name", "n", "", "what is a name")
}