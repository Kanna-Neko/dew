package cmd

import (
	"log"

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
		ReadConfig()
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
		OpenWebsite(luoguDomain + "/problem/solution/CF" + problem)
	},
}
