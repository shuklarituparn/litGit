package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var betCommand = &cobra.Command{
	Use:   "bet",
	Short: "Commits the changes to a repository",
	Run: func(cmd *cobra.Command, args []string) {
		commitMessage, _ := cmd.Flags().GetString("message")
		all, _ := cmd.Flags().GetBool("all")
		if commitMessage == "" {
			fmt.Println("Bruh, you forgot the commit message. Use --message to add one.")
			return
		}
		err := utils.GitCommit(commitMessage, all)
		if err != nil {
			fmt.Printf("Bruh, so this shit happened:%s", err)
			return
		}
		fmt.Println("Bet. Commit made successfully. Vibes secured.")
	},
}
