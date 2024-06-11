/*
Copyright © 2024 Amir Khaki <amirkhaki995@gmail.com>
*/
package cmd

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/Kanna-Neko/dew/link"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

//go:embed ladders.json
var initialLadderContent embed.FS

type Problem struct {
	Difficulty int64    `json:"difficulty"`
	Endpoints  []string `json:"endpoints"`
	Name       string   `json:"name"`
	URL        string   `json:"url"`
}
type ladders map[string][]Problem

var initialLadders ladders
var ladderKeys = []string{
	"0-1300", "1300-1399", "1400-1499", "1500-1599", "1600-1699", "1700-1799",
	"1800-1899", "1900-1999", "2000-2099", "2100-2199", "2200-inf",
	"0-1300-ex", "1300-1399-ex", "1400-1499-ex", "1500-1599-ex", "1600-1699-ex",
	"1700-1799-ex", "1800-1899-ex", "1900-1999-ex", "2000-2099-ex", "2100-2199-ex",
	"2200-inf-ex",
}

var problemStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.RoundedBorder()).
	PaddingLeft(3).
	BorderForeground(lipgloss.Color("63")).
	BorderBottom(true).
	Width(40)
var solvedTitleStyle = lipgloss.NewStyle().
	Bold(true).
	Foreground(lipgloss.Color("#7FFFD4")).
	PaddingLeft(3).
	Inherit(problemStyle)

func init() {
	rootCmd.AddCommand(ladderCmd)
	ladderCmd.PersistentFlags().BoolVarP(&listL, "list", "l",
		false, "list of available ladders")
	data, err := initialLadderContent.ReadFile("ladders.json")
	if err != nil {
		log.Fatal(err)
	}
	if err = json.Unmarshal(data, &initialLadders); err != nil {
		log.Fatal(err)
	}
}

var listL bool
var ladderCmd = &cobra.Command{
	Use:   "ladder",
	Short: "go to specific ladder",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		ReadConfig()
		if len(args) == 0 {
			if listL {
				fmt.Println("existing ladders\n-----------------")
				for _, k := range ladderKeys {
					fmt.Println(k, ":", len(initialLadders[k]), "problems")
				}
				return
			}
			// show list of questions
			currentLadder := viper.GetString("ladder")
			if currentLadder == "" {
				currentLadder = ladderKeys[0]
				viper.Set("ladder", currentLadder)
			}
			problems, ok := initialLadders[currentLadder]
			if !ok {
				log.Fatal("invalid ladder in config")
			}
			status := link.GetStatus()
			fmt.Println(currentLadder, "\n-----------------")
			for _, p := range problems {
				x := p.Name + " " + p.Endpoints[0]
				if status[strings.ReplaceAll(p.Endpoints[0], "/", "")] {
					fmt.Println(solvedTitleStyle.Render(x))
				} else {
					fmt.Println(problemStyle.Render(x))
				}
			}

		} else {
			_, ok := initialLadders[args[0]]

			if !ok {
				log.Fatal("no ladder called " + args[0])
			}
			viper.Set("ladder", args[0])
		}
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
	},
}
