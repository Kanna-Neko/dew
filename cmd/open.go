package cmd

import (
	"log"
	"os/exec"
	"runtime"

	"github.com/spf13/cobra"
)

func init() {
	Open.AddCommand(OpenGyms)
	Open.AddCommand(OpenMashup)
	rootCmd.AddCommand(Open)
}

var Open = &cobra.Command{
	Use:   "open",
	Short: "a shortcut of opening codeforces website",
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

var OpenGyms = &cobra.Command{
	Use:   "gym",
	Short: "a shortcut of opening codeforces gyms",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			var comm = exec.Command("open", "https://codeforces.com/gyms")
			err := comm.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(runtime.GOOS, "have not supported")
		}
	},
}

var OpenMashup = &cobra.Command{
	Use:   "mashup",
	Short: "a shortcut of opening codeforces gyms",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			var comm = exec.Command("open", "https://codeforces.com/mashups")
			err := comm.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(runtime.GOOS, "have not supported")
		}
	},
}
