package cmd

import (
	"fmt"
	"log"

	"github.com/Kanna-Neko/dew/lang"
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
		lang.ImportLangDic()
	},
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			displayLangDetails()
		} else {
			changeLang(args[0])
		}
	},
}

func displayLangDetails() {
	fmt.Printf("current lang: %s\n", viper.GetString("lang"))
	fmt.Println("----------------------------")
	fmt.Println("support program language:")
	for k, lan := range lang.LangDic {
		fmt.Println(k + ":")
		fmt.Printf("   name: %s\n", lan.Name)
		fmt.Printf("   codefile: %s\n", lan.Codefile)
		fmt.Printf("   isCompileLang: %v\n", lan.IsComplieLang)
		fmt.Printf("   compileCommand: %s\n", lan.CompileCode("$codefile").String())
		fmt.Printf("   RunCommand: %s\n", lan.RunCode("$codefile").String())
		fmt.Printf("   programTypeId: %s\n", lan.ProgramTypeId)
	}
	fmt.Println("----------------------------")
	fmt.Printf("current lang: %s\n", viper.GetString("lang"))
}

func changeLang(langShortcut string) {
	_, ok := lang.LangDic[langShortcut]
	if !ok {
		log.Fatal("don't support language: " + langShortcut)
	}
	viper.Set("lang", langShortcut)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}
