package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(config)
}

var config = &cobra.Command{
	Use:   "config",
	Short: "config codeforces handle and password",
	Run: func(cmd *cobra.Command, args []string) {
		configFunc()
	},
}

func configFunc() {
	fmt.Println("please input your codeforces handle(account)")
	var handle string
	fmt.Scanln(&handle)
	fmt.Println("please input your codeforces password")
	var password string
	fmt.Scanln(&password)
	fmt.Printf("your handle is %s\nyour password is %s\n", handle, password)
	viper.Set("handle", handle)
	viper.Set("password", password)
	os.Mkdir("codeforces", 0777)
	viper.WriteConfig()
	fmt.Println("config save as ./codeforces/config.yaml")
}
