package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for git clone

var snatchCmd = &cobra.Command{
	Use:   "snatch",
	Short: "Clone a repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("clone")
	},
}
