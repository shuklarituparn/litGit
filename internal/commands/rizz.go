package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for the git Init

var rizzCmd = &cobra.Command{
	Use:   "rizz",
	Short: "Init git in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rizz called")
	},
}
