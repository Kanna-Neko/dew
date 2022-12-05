package cmd

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/dew/link"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(raceCmd)
}

var raceCmd = &cobra.Command{
	Use:   "race",
	Short: "set contest env",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		viper.Set("race", args[0])
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
		link.Login()
		sp := spinner.New(spinner.CharSets[34], 100*time.Millisecond)
		sp.FinalMSG = "get contest info completed\n"
		sp.Prefix = "geting contest info "
		sp.Start()
		t := link.GetContestCountdown(args[0])
		for t > 0 {
			sp.Prefix = fmt.Sprintf("countdown: %d ", t)
			time.Sleep(time.Second)
			t--
		}
		sp.Prefix = "geting testcases "
		info := link.GetContestInfo(args[0])
		sp.Stop()
		var problems []string
		for _, v := range info.Problems {
			problems = append(problems, strconv.Itoa(v.ContestId)+v.Index)
		}
		for _, v := range problems {
			GetTestcases(v)
		}
		fmt.Println("testcases download complete")
	},
}
