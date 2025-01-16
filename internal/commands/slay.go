package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var slayCmd = &cobra.Command{
	Use:   "slay [files...]",
	Short: "Adds the file contents",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitAdd(args...)
		if err != nil {
			fmt.Printf("Yo, something went wrong while adding the files: %v 😬\n", err)
		} else {
			fmt.Println("Ayy, files added like a boss! 🙌🔥")
		}

	},
}
