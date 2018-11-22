package cmd

import "github.com/spf13/cobra"

type GetOptions struct {

}

type GetFlags struct {
	name string
}

func NewGetCommand() *cobra.Command {
	var command = &cobra.Command{
		Use: "get [NAME]",
		Short: "edit a resource",
		Run: func(cmd *cobra.Command, args []string) {

		},
	}

	getFlags := &GetFlags{}
	addGetFlags(command, getFlags)
	return command
}

func addGetFlags(command *cobra.Command, flags *GetFlags) {

	command.Flags().StringVarP(&flags.name, "name", "n", "", "what is a name")
}