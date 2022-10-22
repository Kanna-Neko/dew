package cmd

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	InitCmd.AddCommand(config)
}

var config = &cobra.Command{
	Use:   "config",
	Short: "config codeforces handle and password",
	Run: func(cmd *cobra.Command, args []string) {
		configFunc()
	},
}

func configFunc() {
	if !checkConfigFile() {
		os.Mkdir("codeforces", 0777)
		os.Create("./codeforces/config.yaml")
	}
	ReadConfig()
	fmt.Println("please input your codeforces handle(account)")
	var handle string
	fmt.Scanln(&handle)
	fmt.Println("please input your codeforces password")
	var password string
	fmt.Scanln(&password)
	fmt.Printf("your handle is %s\nyour password is %s\n", handle, password)
	viper.Set("handle", handle)
	viper.Set("password", password)
	viper.Set("codeFile", "main.cpp")
	viper.WriteConfig()
	fmt.Println("handle and password save into ./codeforces/config.yaml")
	cj := uispinner.New()
	defer cj.Stop()
	no1 := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("user rating is initing").SetComplete("user rating init complete")
	defer no1.Done()
	cj.Start()
	SaveRating(handle)
}

func ReadConfig() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("please use cf init first, error : %s\n", err)
	}
}
