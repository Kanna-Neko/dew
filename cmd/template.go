package cmd

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(templateCmd)
}

var templateCmd = &cobra.Command{
	Use:   "template",
	Short: "generate template",
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
		ioutil.WriteFile("main.cpp", codeContent, 0777)
	},
}
