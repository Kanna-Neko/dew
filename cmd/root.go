package cmd

import (
	"fmt"
	"os"

	"github.com/Kanna-Neko/dew/link"
	"github.com/spf13/cobra"
)

const codeforcesDomain = "https://codeforces.com"

func init() {
	rootCmd.AddCommand(loginCmd)
}

var rootCmd = &cobra.Command{
	Use:   "dew",
	Short: "dew is a contest generator",
	Long:  `A assistant which can help you test your program you will submit on codeforces. and a contest generator base on codeforces, which will help you better practice. you can use generate command a contest on mashup, whose problems is random`,
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
