package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(problem)
}

var problem = &cobra.Command{
	Use:   "problem",
	Short: "open problem in codeforces",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/problemset/problem/" + args[0][:len(args[0])-1] + "/" + args[0][len(args[0])-1:])
	},
}
