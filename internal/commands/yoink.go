package commands

import (
	"fmt"
	"github.com/spf13/cobra"
)

// for the git pull

var yoinkCommand = &cobra.Command{
	Use:   "yoink",
	Short: "Pulls the changes from the remote repository",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("yoink called")
	},
}
