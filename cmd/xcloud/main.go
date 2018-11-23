package main

import (
	"fmt"
	"os"

	"github.com/xcloudnative/xcloud/cmd/xcloud/commands"
)

func main() {
	//if err := cmd.NewRootCommand().Execute(); err != nil {
	//	fmt.Println(err)
	//	os.Exit(1)
	//}
	if err := commands.NewCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}
