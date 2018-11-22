package cmd

import "github.com/spf13/cobra"

type CreateOptions struct {

}

type CreateFlags struct {
	name string
}

// NewCommandCreate create command
func NewCreateCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "create [NAME]",
		Short: "Create a resource",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	createFlags := &CreateFlags{}
	addCreateFlags(command, createFlags)
	return command
}

func addCreateFlags(command *cobra.Command, flags *CreateFlags) {

	command.Flags().StringVarP(&flags.name, "name", "n", "", "what is a name")
}