package cmd

import (
	"io/ioutil"
	"log"

	"github.com/jaxleof/dew/lang"
	"github.com/jaxleof/dew/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var file string

func init() {
	rootCmd.AddCommand(submitCommand)
	submitCommand.PersistentFlags().StringVarP(&file, "file", "f", "", "specify a codefile name which will be submit")
}

var submitCommand = &cobra.Command{
	Use:   "submit",
	Short: "submit problem",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		ReadConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if file == "" {
			file = viper.GetString("codeFile." + viper.GetString("lang"))
			if file == "" {
				log.Fatal("please check codeFile field in ./codeforces/config.yaml")
			}
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
		language := viper.GetString("lang")
		lan, ok := lang.LangDic[language]
		if !ok {
			log.Fatal("don't support language: " + language)
		}
		link.SubmitCode(contest, index, code, lan.ProgramTypeId)
		OpenWebsite(codeforcesDomain + "/contest/" + contest + "/my")
	},
}
