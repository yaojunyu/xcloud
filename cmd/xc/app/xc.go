package app

import (
	"os"

	"github.com/xcloudnative/xcloud/pkg/pkg/xc/cmd"
)

// Run runs the command
func Run() error {
	/*
		logs.InitLogs()
		defer logs.FlushLogs()
	*/

	cmd := cmd.NewXCCommand(cmd.NewFactory(), os.Stdin, os.Stdout, os.Stderr)
	return cmd.Execute()
}
