package cmd

import (
	"unicode"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(problem)
}

var problem = &cobra.Command{
	Use:   "problem",
	Short: "open problem in codeforces",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		contest, index := splitProblem(args[0])
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
