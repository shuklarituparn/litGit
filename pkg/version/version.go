package version

import (
	"fmt"
	"github.com/spf13/cobra"
)

var version = "1.0.0"

var VersionCmd = &cobra.Command{
	Use:   "version",
	Short: "Displays the current version of the app",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("App Version: %s\n", version)
	},
}
