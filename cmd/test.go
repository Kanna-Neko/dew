package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/dew/lang"
	"github.com/jaxleof/dew/link"
	testmanager "github.com/jaxleof/dew/testManager"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "specify a codefile name which will be submit")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test problem",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		if file == "" {
			file = viper.GetString("codeFile." + viper.GetString("lang"))
			if file == "" {
				log.Fatal("please check codeFile field in ./codeforces/config.yaml")
			}
		}
		var problem string
		if len(args) == 1 {
			if len(args[0]) == 1 {
				if viper.GetString("race") == "" {
					log.Fatal("please use race command first")
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
		language := viper.GetString("lang")
		lan, ok := lang.LangDic[language]
		if !ok {
			log.Fatal("don't support language: " + language)
		}
		tests := GetTestcases(problem)
		sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		sp.Prefix = "testing "
		sp.Start()
		if lan.IsComplieLang {
			compile := lan.CompileCode(file)
			compile.Stderr = os.Stderr
			defer os.Remove("./cat")
			compile.Run()
		}
		for _, v := range tests.Tests {
			in := strings.NewReader(v.Input)
			cmd := lan.RunCode(file)
			cmd.Stdin = in
			cmd.Stderr = os.Stderr
			out, err := cmd.Output()
			if err != nil {
				sp.Stop()
				log.Fatal(err)
			}
			v.Input = strings.Trim(v.Input, " \n")
			out = bytes.Trim(out, " \n")
			v.Output = strings.Trim(v.Output, " \n")
			if !bytes.Equal(out, []byte(v.Output)) {
				sp.Stop()
				fmt.Printf("oops!\n----------in-----------\n%s\n----------out----------\n%s\n---------answer--------\n%s", v.Input, string(out), v.Output)
				return
			}
		}
		sp.Stop()
		if len(tests.Tests) == 0 {
			fmt.Println("Warning: sample is empty")
		} else {
			fmt.Println("OK")
		}
	},
}

func GetTestcases(problem string) testmanager.Testcases {
	sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	sp.Prefix = "geting TestCases "
	sp.Start()
	defer sp.Stop()
	contest, index := splitProblem(problem)
	var manager = testmanager.ExtractManager()
	manager = testmanager.ManagerPush(manager, problem)
	manager = testmanager.ManagerDeleteExtra(manager)
	testmanager.StoreManager(manager)
	if testmanager.IsTestcaseExist(problem) {
		return testmanager.ExtractTestcase(problem)
	} else {
		link.Login()
		input, output := link.GetSample(contest, index)
		var res testmanager.Testcases
		for i := 0; i < len(input); i++ {
			res.Tests = append(res.Tests, testmanager.Testcase{Input: input[i], Output: output[i]})
		}
		testmanager.StoreProblemTest(problem, res)
		return res
	}
}
