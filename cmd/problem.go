package cmd

import (
	"log"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(problemCmd)
}

var problemCmd = &cobra.Command{
	Use:   "problem",
	Short: "open problem in codeforces",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var problemInfo string
		ReadConfig()
		if len(args) == 0 {
			problemInfo = viper.GetString("problem")
			if problemInfo == "" {
				log.Fatal("please random or specify a problem")
			}
		} else {
			if len(args[0]) == 1 {
				contest := viper.GetString("race")
				if contest == "" {
					log.Fatal("please use race command first")
				}
				problemInfo = contest + args[0]
			} else {
				problemInfo = args[0]
				viper.Set("problem", problemInfo)
				err := viper.WriteConfig()
				if err != nil {
					log.Fatal(err)
				}
			}
		}
		contest, index := splitProblem(problemInfo)
		if isGym(contest) {
			OpenWebsite(codeforcesDomain + "/gym/" + contest + "/problem/" + index)
		} else {
			OpenWebsite(codeforcesDomain + "/contest/" + contest + "/problem/" + index)
		}
		GetTestcases(problemInfo)
	},
}

func splitProblem(in string) (string, string) {
	len := len(in)
	if unicode.IsLetter(rune(in[len-2])) {
		return in[:len-2], in[len-2:]
	} else {
		return in[:len-1], in[len-1:]
	}
}
