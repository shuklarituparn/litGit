package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for git add

var slayCmd = &cobra.Command{
	Use:   "slay",
	Short: "Adds the file contents",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("slay called")
	},
}
