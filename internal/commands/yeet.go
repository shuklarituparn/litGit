package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var yeetCmd = &cobra.Command{
	Use:   "yeet",
	Short: "Pushes the change to remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitPush()
		if err != nil {
			fmt.Printf("Ayo, push failed! 😩 Something went wrong: %v\n", err)
		} else {
			fmt.Println("Yeet! Changes pushed like a boss! 🚀💯")
		}
	},
}
