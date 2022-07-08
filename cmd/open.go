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
		if runtime.GOOS == "darwin" {
			var comm = exec.Command("open", "https://codeforces.com")
			err := comm.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(runtime.GOOS, "have not supported")
		}
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
	Short: "a shortcut of opening codeforces mashup",
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

var OpenStatus = &cobra.Command{
	Use:   "status",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		if runtime.GOOS == "darwin" {
			var comm = exec.Command("open", "https://codeforces.com/problemset/status?my=on")
			err := comm.Run()
			if err != nil {
				log.Fatal(err)
			}
		} else {
			log.Fatal(runtime.GOOS, "have not supported")
		}
	},
}

var OpenRandom = &cobra.Command{
	Use:   "random",
	Short: "a shortcut of opening codeforces status",
	Run: func(cmd *cobra.Command, args []string) {
		OpenRandomFunc()
	},
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
	if runtime.GOOS == "darwin" {
		var comm = exec.Command("open", "https://codeforces.com/problemset/problem/"+randomString[:len(randomString)-1]+"/"+randomString[len(randomString)-1:])
		err := comm.Run()
		if err != nil {
			log.Fatal(err)
		}
	} else {
		log.Fatal(runtime.GOOS, "have not supported")
	}

}
