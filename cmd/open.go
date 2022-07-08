package cmd

import (
	"io/ioutil"
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	Open.AddCommand(OpenGyms)
	Open.AddCommand(OpenMashup)
	Open.AddCommand(OpenStatus)
	Open.AddCommand(OpenRandom)
	rootCmd.AddCommand(Open)
}

var Open = &cobra.Command{
	Use:   "open",
	Short: "a shortcut of opening codeforces website",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com")
	},
}

var OpenGyms = &cobra.Command{
	Use:   "gym",
	Short: "a shortcut of opening codeforces gyms",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/gyms")
	},
}

var OpenMashup = &cobra.Command{
	Use:   "mashup",
	Short: "a shortcut of opening codeforces mashup",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/mashups")
	},
}

var OpenStatus = &cobra.Command{
	Use:   "status",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite("https://codeforces.com/problemset/status?my=on")
	},
}

var OpenRandom = &cobra.Command{
	Use:   "random",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		OpenRandomFunc()
	},
}

var commands = map[string]string{
	"windows": "start",
	"darwin":  "open",
	"linux":   "xdg-open",
}

func OpenWebsite(path string) {
	run, ok := commands[runtime.GOOS]
	if !ok {
		log.Fatalf("don't know how to open things on %s platform", runtime.GOOS)
	}

	cmd := exec.Command(run, path)
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func OpenRandomFunc() {
	var random, err = ioutil.ReadFile("./codeforces/random.helloWorld")
	randomString := string(random)
	if randomString == "" {
		log.Fatal("random.helloWorld is empty")
	}
	if err != nil {
		log.Fatal(err)
	}
	OpenWebsite("https://codeforces.com/problemset/problem/" + randomString[:len(randomString)-1] + "/" + randomString[len(randomString)-1:])
}
