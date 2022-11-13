package cmd

import (
	"log"

	"github.com/jaxleof/cf-helper/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(raceCmd)
}

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "set contest env",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		viper.Set("race", args[0])
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		problems := link.GetContestProblems(args[0])
		for _, v := range problems {
			GetTestcases(v)
		}
	},
}
