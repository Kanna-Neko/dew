package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/Kanna-Neko/dew/lang"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	modify bool
)

var envSet = map[string]bool{
	"handle":   true,
	"password": true,
	"rating":   true,
	"problem":  true,
	"race":     true,
	"proxy":    true,
	"lang":     true,
}
var envSlice = []string{"handle", "password", "rating", "problem", "race", "proxy", "apikey", "secret"}

func init() {
	env.PersistentFlags().BoolVarP(&modify, "write", "w", false, `you can modify env like "env -w proxy=http://127.0.0.1:20245" when you use flag -w or --write`)
	for k := range lang.LangDic {
		envSet["codefile."+k] = true
	}
	rootCmd.AddCommand(env)
}

var env = &cobra.Command{
	Use:   "env",
	Short: "print config env",
	PreRun: func(cmd *cobra.Command, args []string) {
		if !checkConfigFile() {
			log.Fatal("config file is not exist, please use init command")
		}
		ReadConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			for _, arg := range envSlice {
				fmt.Printf("%s: %v\n", arg, viper.Get(arg))
			}
			return
		}
		if !modify {
			for _, arg := range args {
				if envSet[arg] {
					fmt.Println(viper.Get(arg))
				}
			}
		} else {
			for _, arg := range args {
				var splited = strings.SplitN(arg, "=", 2)
				if len(splited) != 2 {
					log.Fatal("= is required")
				}
				if envSet[splited[0]] {
					viper.Set(splited[0], splited[1])
				}
			}
			err := viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}

func checkConfigFile() bool {
	_, err := os.Stat("./codeforces/config.yaml")
	isExist := !os.IsNotExist(err)
	return isExist
}
