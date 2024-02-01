package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/Kanna-Neko/dew/lang"
	"github.com/Kanna-Neko/dew/link"
	testmanager "github.com/Kanna-Neko/dew/testManager"
	"github.com/briandowns/spinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var testAll bool

func init() {
	rootCmd.AddCommand(testCmd)
	testCmd.PersistentFlags().StringVarP(&file, "file", "f", "", "specify a codefile name which will be submit")
	testCmd.Flags().BoolVarP(&testAll, "all", "a", false, "test all tests, default is test until the first wrong test")
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test problem",
	Args:  cobra.MaximumNArgs(1),
	PreRun: func(cmd *cobra.Command, args []string) {
		lang.ImportLangDic()
	},
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		if file == "" {
			file = viper.GetString("language." + viper.GetString("lang") + ".codefile")
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
		if lan.IsComplieLang {
			compile := lan.CompileCode(file)
			compile.Stderr = os.Stderr
			defer os.Remove("./cat")
			defer os.Remove("./cat.exe")
			compile.Run()
		}
		sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		sp.Prefix = "testing "
		sp.Start()
		var isok = true
		for index, v := range tests.Tests {
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
			out = bytes.ReplaceAll(out, []byte("\r"), []byte(""))
			out = bytes.Trim(out, " \n")
			v.Output = strings.Trim(v.Output, " \n")
			v.Output = strings.ReplaceAll(v.Output, "\r", "")
			if !bytes.Equal(out, []byte(v.Output)) {
				sp.Stop()
				fmt.Printf("oops! testcase %d wrong.\n----------in-----------\n%s\n----------out----------\n%s\n---------answer--------\n%s\n\n", index+1, v.Input, string(out), v.Output)
				isok = false
				if !testAll {
					return
				}
			}
		}
		sp.Stop()
		if len(tests.Tests) == 0 {
			fmt.Println("Warning: sample is empty")
		} else {
			if isok {
				fmt.Println("OK")
			}
		}
	},
}

func GetTestcases(problem string) testmanager.Testcases {
	contest, index := splitProblem(problem)
	var manager = testmanager.ExtractManager()
	manager = testmanager.ManagerPush(manager, problem)
	manager = testmanager.ManagerDeleteExtra(manager)
	testmanager.StoreManager(manager)
	if testmanager.IsTestcaseExist(problem) {
		return testmanager.ExtractTestcase(problem)
	} else {
		link.Login()
		sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		sp.Prefix = "geting TestCases "
		sp.Start()
		defer fmt.Println("problem " + problem + " fetch done")
		defer sp.Stop()
		input, output := link.GetSample(contest, index)
		var res testmanager.Testcases
		for i := 0; i < len(input); i++ {
			res.Tests = append(res.Tests, testmanager.Testcase{Input: input[i], Output: output[i]})
		}
		testmanager.StoreProblemTest(problem, res)
		return res
	}
}
