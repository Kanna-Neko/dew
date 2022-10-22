package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	luoguDomain = "https://www.luogu.com.cn"
)

func init() {
	rootCmd.AddCommand(tutorial)
}

var tutorial = &cobra.Command{
	Use:   "tutorial",
	Short: "as the name says",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			OpenWebsite(luoguDomain + "/problem/solution/CF" + args[0])
		} else {
			_, err := os.Stat("./codeforces/config.yaml")
			isExist := os.IsNotExist(err)
			if isExist {
				log.Fatal("config file is not exist, please use cf init command")
			}
			ReadConfig()
			if viper.GetString("problem") == "" {
				log.Fatal("please specify a problem first")
			}
			OpenWebsite(luoguDomain + "/problem/solution/CF" + viper.GetString("problem"))
		}
	},
}
