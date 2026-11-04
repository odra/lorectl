package main

import (
	"os"
	"github.com/odra-ai-tools/lorectl/cmd/commands/root"
	"github.com/odra-ai-tools/lorectl/cmd/commands/version"
)

func main() {
	rootcmd := root.Command()
	rootcmd.AddCommand(version.Command())

	if err := rootcmd.Execute(); err != nil {
		os.Exit(1)
	}
}
