package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var yoinkCommand = &cobra.Command{
	Use:   "yoink",
	Short: "Pulls the changes from the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitPull()
		if err != nil {
			fmt.Printf("Bruh, pull failed! 😤 Something went sideways: %v\n", err)
		} else {
			fmt.Println("Yoink! Changes pulled in like a pro! 🙌✨")
		}
	},
}
