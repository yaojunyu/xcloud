package cmd

import (
	"io"

	"github.com/spf13/cobra"
	"gopkg.in/AlecAivazis/survey.v1/terminal"
)

// NewXCCommand create the `xc` command and its all nested children.
func NewXCCommand(f Factory, in terminal.FileReader, out terminal.FileWriter, err io.Writer) *cobra.Command {
	cmds := &cobra.Command{
		Use:   "xc",
		Short: "xc is a command line tool for working with XCloud",
		Long: `
		`,
		Run: runHelp,
	}

	return cmds
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
