package cmd

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/cf-helper/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(testCmd)
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "test problem",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		var problem string
		if len(args) == 1 {
			if len(args[0]) == 1 {
				if viper.GetString("race") == "" {
					log.Fatal("please use cf race first")
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
		contest, index := splitProblem(problem)
		sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		sp.Prefix = "testing "
		sp.Start()
		input, output := link.GetSample(contest, index)
		compile := exec.Command("g++", viper.GetString("codeFile"), "-o", "cat")
		compile.Stderr = os.Stderr
		defer os.Remove("./cat")
		compile.Run()
		for i, v := range input {
			in := strings.NewReader(v)
			cmd := exec.Command("./cat")
			cmd.Stdin = in
			cmd.Stderr = os.Stderr
			out, err := cmd.Output()
			if err != nil {
				sp.Stop()
				log.Fatal(err)
			}
			out = bytes.Trim(out, " \n")
			output[i] = strings.Trim(output[i], " \n")
			if !bytes.Equal(out, []byte(output[i])) {
				fmt.Println(out)
				fmt.Println([]byte(output[i]))
				sp.Stop()
				fmt.Printf("oops!\nin:\n%s\nout:\n%s\nanswer:\n%s", v, string(out), output[i])
				return
			}
		}
		sp.Stop()
		if len(input) == 0 {
			fmt.Println("Warning: sample input is empty")
		}
		if len(output) == 0 {
			fmt.Println("Warning: sample output is empty")
		}
		if len(input) == 0 && len(output) == 0 {
			fmt.Println("please check the validity of problem")
		} else {
			fmt.Println("OK")
		}
	},
}
