package cmd

import "github.com/spf13/cobra"

type EditOptions struct {

}

type EditFlags struct {
	name string
}

func NewEditCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "edit [NAME]",
		Short: "Edit a resource",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	editFlags := &EditFlags{}
	addEditFlags(command, editFlags)
	return command
}

func addEditFlags(command *cobra.Command, flags *EditFlags) {

	command.Flags().StringVarP(&flags.name, "name", "n", "", "what is a name")
}