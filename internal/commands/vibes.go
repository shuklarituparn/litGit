package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for the git status

var vibesCmd = &cobra.Command{
	Use:   "vibes",
	Short: "Shows the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("vibes called")
	},
}
