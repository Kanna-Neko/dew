package cmd

import (
	"io/ioutil"
	"log"

	"github.com/jaxleof/cf-helper/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(submitCommand)
}

var submitCommand = &cobra.Command{
	Use:   "submit",
	Short: "submit",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		file := viper.GetString("codeFile")
		if file == "" {
			log.Fatal("please check codeFile field in ./codeforces/config.yaml")
		}
		var problem string
		if len(args) == 1 {
			problem = args[0]
		} else {
			problem = viper.GetString("random")
		}
		contest, index := splitProblem(problem)
		code, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		link.SubmitCode(contest, index, code)
	},
}
