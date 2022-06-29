package main

import (
	"cf/cmd"
	"log"

	"github.com/spf13/viper"
)

func main() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	cmd.Execute()
}
