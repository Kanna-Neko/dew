package cmd

import (
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(InitCmd)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init somethings",
	Run: func(cmd *cobra.Command, args []string) {
		configFunc()
		updateFunc()
	},
}
