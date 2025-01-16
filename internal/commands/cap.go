package commands

import (
	"github.com/spf13/cobra"
)

// for git reset

var Cap = &cobra.Command{
	Use:   "cap",
	Short: "Resets the commit",
	Long:  `Resets the commit that you made`,
	Run: func(cmd *cobra.Command, args []string) {

	},
}
