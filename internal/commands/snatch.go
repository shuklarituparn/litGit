package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var snatchCmd = &cobra.Command{
	Use:   "snatch",
	Short: "Clone a repository",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitClone(args[0])
		if err != nil {
			fmt.Printf("Yo, cloning failed! 😕 Here's the deal: %v\n", err)
		} else {
			fmt.Println("Ayy, repo cloned like a champ! 🚀🔥")
		}

	},
}
