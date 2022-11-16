package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(templateCmd)
	templateCmd.Flags().BoolVarP(&seeAll, "all", "a", false, "see all template when you use -a")
}

var seeAll bool

var templateCmd = &cobra.Command{
	Use:     "template",
	Short:   "generate template",
	Aliases: []string{"tmp"},
	Run: func(cmd *cobra.Command, args []string) {
		if seeAll {
			info, err := ioutil.ReadDir(templateDir)
			if err != nil {
				log.Fatal(err)
			}
			var templateName []string
			for _, v := range info {
				if v.IsDir() {
					templateName = append(templateName, v.Name())
				}
			}
			sort.Strings(templateName)
			fmt.Println("existing template\n-----------------")
			for _, v := range templateName {
				fmt.Println("• " + v)
			}
			return
		}
		if len(args) == 0 {
			info, err := ioutil.ReadDir(templateDir + "default")
			if err != nil {
				log.Fatal(err)
			}
			for _, v := range info {
				if !v.IsDir() {
					data, err := ioutil.ReadFile(templateDir + "default/" + v.Name())
					if err != nil {
						log.Fatal(err)
					}
					ioutil.WriteFile(v.Name(), data, 0777)
				}
			}
		} else {
			for _, v := range args {
				if !checkIstemplateExist(v) {
					log.Fatalf("template %s is not exist", v)
				}
			}
			var used map[string]bool = make(map[string]bool)
			for _, v := range args {
				info, err := ioutil.ReadDir(templateDir + v)
				if err != nil {
					log.Fatal(err)
				}
				for _, vv := range info {
					if used[vv.Name()] {
						continue
					}
					used[vv.Name()] = true
					if !vv.IsDir() {
						data, err := ioutil.ReadFile(templateDir + v + "/" + vv.Name())
						if err != nil {
							log.Fatal(err)
						}
						ioutil.WriteFile(vv.Name(), data, 0777)
					}
				}

			}

		}
	},
}

func checkIstemplateExist(templateName string) bool {
	_, err := os.Stat(templateDir + templateName)
	if err == nil {
		return true
	} else {
		return false
	}
}
