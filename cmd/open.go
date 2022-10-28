package cmd

import (
	"log"
	"os/exec"
	"runtime"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	OpenCmd.AddCommand(OpenGyms)
	OpenCmd.AddCommand(OpenMashup)
	rootCmd.AddCommand(OpenStatus)
	rootCmd.AddCommand(OpenCmd)
}

var OpenCmd = &cobra.Command{
	Use:   "open",
	Short: "a shortcut of opening codeforces website",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite(codeforcesDomain)
	},
}

var OpenGyms = &cobra.Command{
	Use:   "gym",
	Short: "a shortcut of opening codeforces gyms",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			OpenWebsite(codeforcesDomain + "/gyms")
		} else {
			OpenWebsite(codeforcesDomain + "/gym/" + args[0])
		}

	},
	Args: cobra.MaximumNArgs(1),
}

var OpenMashup = &cobra.Command{
	Use:   "mashup",
	Short: "a shortcut of opening codeforces mashup",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite(codeforcesDomain + "/mashups")
	},
}

var OpenStatus = &cobra.Command{
	Use:   "status",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		OpenWebsite(codeforcesDomain + "/problemset/status?my=on")
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
	var cmd *exec.Cmd
	if runtime.GOOS == "windows" {
		cmd = exec.Command("cmd.exe", "/C", "start", path)
	} else {
		cmd = exec.Command(run, path)
	}
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func OpenRandomFunc(info problemInfo) {
	OpenWebsite(codeforcesDomain + "/contest/" + strconv.Itoa(info.ContestId) + "/problem/" + info.Index)
}
