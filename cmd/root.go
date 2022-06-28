package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "cf is a contest generator",
	Long:  `A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random or custom`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Usage()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
