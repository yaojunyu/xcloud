package cmd

import (
	"github.com/xcloudnative/xcloud/pkg/xc/cmd/templates"
	"io"
	"strings"

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

	createCommands := NewCmdCreate(f, in, out, err)
	installCommands := []*cobra.Command{
		NewCmdInstall(f, in, out, err),
	}
	installCommands = append(installCommands, findCommands("cluster", createCommands)...)

	groups := templates.CommandGroups{
		{
			Message:  "Installing:",
			Commands: installCommands,
		},
		{
			Message: "Working with Jenkins X resources:",
			Commands: []*cobra.Command{
				//getCommands,
				//editCommands,
				createCommands,
				//updateCommands,
				//deleteCommands,
				//NewCmdStart(f, in, out, err),
				//NewCmdStop(f, in, out, err),
			},
		},
	}

	groups.Add(cmds)

	return cmds
}

func findCommands(subCommand string, commands ...*cobra.Command) []*cobra.Command {
	answer := []*cobra.Command{}
	for _, parent := range commands {
		for _, c := range parent.Commands() {
			if commandHasParentName(c, subCommand) {
				answer = append(answer, c)
			} else {
				childCommands := findCommands(subCommand, c)
				if len(childCommands) > 0 {
					answer = append(answer, childCommands...)
				}
			}
		}
	}
	return answer
}

func fullPath(command *cobra.Command) string {
	name := command.Name()
	parent := command.Parent()
	if parent != nil {
		return fullPath(parent) + " " + name
	}
	return name
}

func commandHasParentName(command *cobra.Command, name string) bool {
	path := fullPath(command)
	return strings.Contains(path, name)
}

func runHelp(cmd *cobra.Command, args []string) {
	cmd.Help()
}
