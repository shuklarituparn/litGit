package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for the git push

var yeetCmd = &cobra.Command{
	Use:   "yeet",
	Short: "Pushes the change to remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yeet called")
	},
}
