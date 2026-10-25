package root

import (
	"github.com/spf13/cobra"
)

var verbose bool
var cfgFilePath string

var rootCmd = &cobra.Command{
	Use: "lorectl",
	Short: "manage stuff",
}
 
func init() {
	// --config flag
	rootCmd.PersistentFlags().StringVarP(&cfgFilePath, "config", "c", "", "path to config file")
	
	// --verbose boolean flag
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose mode")
}

func Command() *cobra.Command {
	return rootCmd
}
