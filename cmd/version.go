package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

const VERSION = "v1.2.11"

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "show the version of dew program",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(VERSION)
	},
}
