package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(info)
}

var info = &cobra.Command{
	Use:   "info",
	Short: "print config info",
	Run: func(cmd *cobra.Command, args []string) {
		if !checkConfigFile() {
			log.Fatal("config file is not exist, please use cf init command")
		}
		ReadConfig()
		fmt.Printf("handle: %v\npassword: %v\nrating:%v\nrandom problem: %v\n", viper.Get("handle"), viper.Get("password"), viper.Get("rating"), viper.Get("random"))
	},
}

func checkConfigFile() bool {
	_, err := os.Stat("./codeforces/config.yaml")
	isExist := !os.IsNotExist(err)
	return isExist
}
