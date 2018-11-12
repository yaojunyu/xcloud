package main

import (
	"os"

	"github.com/xcloudnative/xcloud/cmd/xc/app"
)

func main() {
	if err := app.Run(); err != nil {
		os.Exit(1)
	}
	os.Exit(0)
}
