package cmd

import (
	"fmt"
	"os"

	"github.com/jaxleof/cf-helper/link"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(loginCmd)
}

var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "cf is a contest generator",
	Long:  `A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random or custom`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

var loginCmd = &cobra.Command{
	Use:   "login",
	Short: "manually login",
	Run: func(cmd *cobra.Command, args []string) {
		link.ManuallyLogin()
	},
}
