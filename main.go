package main

import (
	"os"

	"github.com/maxbet1507/l7ping/cmd"
)

func main() {
	if err := cmd.Cmd.Execute(); err != nil {
		os.Exit(1)
	}
	return
}
