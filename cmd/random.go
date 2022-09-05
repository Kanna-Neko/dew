package cmd

import "github.com/spf13/cobra"

func init() {
	rootCmd.AddCommand(random)
}

var random = &cobra.Command{
	Use:   "random",
	Short: "alias to cf generate random",
	Run: func(cmd *cobra.Command, args []string) {
		Random()
	},
}
