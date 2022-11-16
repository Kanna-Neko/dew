package cmd

import (
	"fmt"
	"log"

	"github.com/jaxleof/dew/lang"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(langCmd)
}

var langCmd = &cobra.Command{
	Use:   "lang [language]",
	Short: "switch program language",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		ReadConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			fmt.Printf("current lang: %s\n", viper.GetString("lang"))
			fmt.Println("----------------------------")
			fmt.Println("support program language:")
			for _, v := range lang.LangSlice {
				fmt.Println(v + ":")
				var lan = lang.LangDic[v]
				fmt.Printf("   name:%s\n", lan.Name)
				fmt.Printf("   codefile:%s\n", viper.GetString("codefile."+v))
			}
		} else {
			_, ok := lang.LangDic[args[0]]
			if !ok {
				log.Fatal("don't support language: " + args[0])
			}
			viper.Set("lang", args[0])
			err := viper.WriteConfig()
			if err != nil {
				log.Fatal(err)
			}
		}
	},
}
