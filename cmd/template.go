package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "generate template",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !checkConfigFile() {
			log.Fatal("config file is not exist, please use init command")
		}
		ReadConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		_, err := os.Stat("./codeforces/default.cpp")
		isNotExist := os.IsNotExist(err)
		if isNotExist {
			log.Fatal("./codeforces/default.cpp don't exist")
		}
		codeContent, err := ioutil.ReadFile("./codeforces/default.cpp")
		if err != nil {
			log.Fatal(err)
		}
		ioutil.WriteFile(viper.GetString("codefile"), codeContent, 0777)
	},
}
