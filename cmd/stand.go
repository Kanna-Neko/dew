package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(standCmd)
	standCmd.Flags().BoolVarP(&standAll, "all", "a", false, "open all standing if you add this flag")
}

var standAll bool

var standCmd = &cobra.Command{
	Use:   "stand [contestId]",
	Short: "open standing",
	PreRun: func(cmd *cobra.Command, args []string) {
		ReadConfig()
	},
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var contest string
		if len(args) == 1 {
			contest = args[0]
		} else {
			contest = viper.GetString("race")
		}
		if contest == "" {
			log.Fatal("please use race command to specify a contest first")
		}
		if standAll {
			if isGym(contest) {
				OpenWebsite(fmt.Sprintf("https://codeforces.com/gym/%s/standings", contest))

			} else {
				OpenWebsite(fmt.Sprintf("https://codeforces.com/contest/%s/standings", contest))
			}
		} else {
			if isGym(contest) {
				OpenWebsite(fmt.Sprintf("https://codeforces.com/gym/%s/standings/friends/true", contest))
			} else {
				OpenWebsite(fmt.Sprintf("https://codeforces.com/contest/%s/standings/friends/true", contest))
			}
		}
	},
}
