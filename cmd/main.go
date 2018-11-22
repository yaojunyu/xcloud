package main

import (
	"fmt"
	"os"

	"github.com/xcloudnative/xcloud/cmd/xcloud/cmd"
)

func main() {
	if err := cmd.NewRootCommand().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
