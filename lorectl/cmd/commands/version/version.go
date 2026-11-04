package version

import (
	"github.com/spf13/cobra"
	"github.com/odra-ai-tools/lorectl/pkg/version"
)

var versionCmd = &cobra.Command{
	Use: "version",
	Short: "Show program's version",
	RunE: func (cmd *cobra.Command, args []string) error {
		cmd.Printf("v%s\n", version.Version)

		return nil
	},
}

func Command() *cobra.Command {
	return versionCmd
}
