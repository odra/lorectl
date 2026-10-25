package main

import (
	"os"
	"github.com/odra-ai-tools/lorectl/cmd/commands/root"
)

func main() {
	rootcmd := root.Command()

	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}
