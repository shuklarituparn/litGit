package commands

import (
	"fmt"
	"github.com/shuklarituparn/litGit/pkg/version"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:     "yh",
	Short:   "A genZ git cli",
	Aliases: []string{"yh", "yuh"},
	Long:    "litGit is a genZ git cli that works and vibes fr fr no cap",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Help()
		} else {
			fmt.Fprintln(os.Stderr, "Usage: yh [command]")
		}
	},
}

func init() {
	rootCmd.PersistentFlags().BoolP("help", "h", false, "Help for yh")
	betCommand.Flags().BoolP("all", "a", false, "stage and commit all changes")
	betCommand.Flags().StringP("message", "m", "", "commit message")
	rootCmd.SetHelpCommand(&cobra.Command{
		Use:   "help",
		Short: `help for yh`,
		Run: func(cmd *cobra.Command, args []string) {
			err := cmd.Parent().Help()
			if err != nil {
				fmt.Fprintf(os.Stderr, "following error happended: %s", err)
			}
		},
	})
	rootCmd.AddCommand(
		Cap,
		rizzCmd,
		slayCmd,
		snatchCmd,
		vibesCmd,
		yeetCmd,
		yoinkCommand,
		betCommand,
		version.VersionCmd,
	)

}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "nah we don't roll this way '%s'", err)
		os.Exit(1)
	}
}
