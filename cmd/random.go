package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	rootCmd.AddCommand(random)
}

var random = &cobra.Command{
	Use:   "random",
	Short: "alias to cf generate random",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 1 {
			rating, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			randomMinMax(rating, rating)
		} else if len(args) == 2 {
			minRating, err := strconv.Atoi(args[0])
			if err != nil {
				log.Fatal(err)
			}
			maxRating, err := strconv.Atoi(args[1])
			if err != nil {
				log.Fatal(err)
			}
			randomMinMax(minRating, maxRating)
		} else {
			Random()
		}
	},
}

func randomMinMax(min, max int) {
	if min%100 != 0 {
		min = min - min%100 + 100
	}
	max -= max % 100
	ReadConfig()
	var pro []string
	for i := min; i <= max; i += 100 {
		pro = append(pro, strconv.Itoa(i))
	}
	var thisOne = PickOneProblem(pro)
	viper.Set("problem", strconv.Itoa(thisOne.ContestId)+thisOne.Index)
	err := viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
	OpenRandomFunc(thisOne)
}
