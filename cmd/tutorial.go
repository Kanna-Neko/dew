package cmd

import (
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
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
			OpenWebsite("https://www.luogu.com.cn/problem/solution/CF" + args[0])
		} else {
			_, err := os.Stat("./codeforces/config.yaml")
			isExist := os.IsNotExist(err)
			if isExist {
				log.Fatal("config file is not exist, please use cf init command")
			}
			viper.SetConfigFile("./codeforces/config.yaml")
			err = viper.ReadInConfig()
			if err != nil {
				log.Fatal(err)
			}
			OpenWebsite("https://www.luogu.com.cn/problem/solution/CF" + viper.GetString("random"))

		}
	},
}
