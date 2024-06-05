/*
Copyright © 2024 Amir Khaki <amirkhaki995@gmail.com>
*/
package cmd

import (
	"log"
	"strings"

	"github.com/Kanna-Neko/dew/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ladderProblemCmd = &cobra.Command{
	Use:     "ladderProblem",
	Aliases: []string{"lp"},
	Short:   "set current problem to next question in ladder",
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		ladderKey := viper.GetString("ladder")
		ladder, ok := initialLadders[ladderKey]
		if !ok {
			log.Fatal("invalid ladder in config " + ladderKey)
		}
		link.SaveStatus(viper.GetString("handle"))
		status := link.GetStatus()
		var p Problem
		found := false
		for _, v := range ladder {
			info := strings.ReplaceAll(v.Endpoints[0], "/", "")
			if !status[info] {
				found = true
				p = v
				break
			}
		}
		if !found {
			log.Fatal("all problems of this ladder are solved!")
		}
		info := strings.ReplaceAll(p.Endpoints[0], "/", "")
		viper.Set("problem", info)
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		OpenWebsite(p.URL)
		GetTestcases(info)
	},
}

func init() {
	rootCmd.AddCommand(ladderProblemCmd)
}
