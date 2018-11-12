package app

import (
	"os"

	"github.com/xcloudnative/xcloud/pkg/xc/cmd"
)

// Run runs the command
func Run() error {
	cmd := cmd.NewXCCommand(cmd.NewFactory(), os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
