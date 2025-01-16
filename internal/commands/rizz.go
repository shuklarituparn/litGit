package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/internal/utils"
	"github.com/spf13/cobra"
)

var rizzCmd = &cobra.Command{
	Use:   "rizz",
	Short: "Init git in the current directory",
	Run: func(cmd *cobra.Command, args []string) {
		err := utils.GitInit()
		if err != nil {
			fmt.Printf("err:%s", err)
			return
		}
	},
}
