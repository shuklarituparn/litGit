package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var vibesCmd = &cobra.Command{
	Use:   "vibes",
	Short: "Shows the working tree status",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitStatus()
		if err != nil {
			fmt.Printf("Oops, the vibes ain't right! 😬 Something went wrong: %v\n", err)
		} else {
			fmt.Println("Vibes are on point! ✅ All good!")
		}
	},
}
