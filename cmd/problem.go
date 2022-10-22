package cmd

import (
	"log"
	"unicode"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(problem)
}

var problem = &cobra.Command{
	Use:   "problem",
	Short: "open problem in codeforces",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var problem string
		ReadConfig()
		if len(args) == 0 {
			problem = viper.GetString("problem")
			if problem == "" {
				log.Fatal("please random or specify a problem")
			}
		} else {
			problem = args[0]
			viper.Set("problem", problem)
			err := viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
		contest, index := splitProblem(problem)
		OpenWebsite("https://codeforces.com/problemset/problem/" + contest + "/" + index)
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
