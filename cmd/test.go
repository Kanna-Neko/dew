package cmd

import (
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
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		problem := viper.GetString("problem")
		contest, index := splitProblem(problem)
		sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
		sp.Prefix = "testing "
		input, output := link.GetSample(contest, index)
		compile := exec.Command("g++", viper.GetString("codeFile"), "-o", "cat")
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
			if string(out) != output[i] {
				sp.Stop()
				fmt.Printf("oops!\nin:\n%s\nout:\n%s\nanswer:\n%s", v, string(out), output[i])
				return
			}
		}
		fmt.Println("OK")
	},
}
