package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(NewCmd)
}

var NewCmd = &cobra.Command{
	Use:   "new",
	Short: "create a contest",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}
