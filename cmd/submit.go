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
			if len(args[0]) == 1 {
				if viper.GetString("race") == "" {
					log.Fatal("please use cf race first")
				} else {
					problem = viper.GetString("race") + args[0]
				}
			} else {
				problem = args[0]
			}
		} else {
			problem = viper.GetString("problem")
			if problem == "" {
				log.Fatal("please specify a problem first")
			}
		}
		contest, index := splitProblem(problem)
		code, err := ioutil.ReadFile(file)
		if err != nil {
			log.Fatal(err)
		}
		link.SubmitCode(contest, index, code)
	},
}
