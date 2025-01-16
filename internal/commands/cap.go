package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var Cap = &cobra.Command{
	Use:   "cap",
	Short: "Resets the commit",
	Long:  `Resets the commit that you made`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		commitHash := args[0]
		if commitHash == "" {
			fmt.Println("Bruh, give the commit hash atleast")
			return
		}
		err := utils.GitReset(commitHash)
		if err != nil {
			fmt.Printf("double cap fr fr: %s", err)
			return
		}
	},
}
