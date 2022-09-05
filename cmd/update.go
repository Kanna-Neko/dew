package cmd

import (
	"cf/mashup"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const problemApi = "https://codeforces.com/api/problemset.problems"

func init() {
	rootCmd.AddCommand(UpdateCmd)
}

var UpdateCmd = &cobra.Command{
	Use:   "update",
	Short: "update problem data",
	Run: func(cmd *cobra.Command, args []string) {
		updateFunc()
	},
}

func updateFunc() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	cj := uispinner.New()
	no1 := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("problem data updating").SetComplete("problem data update complete")
	cj.Start()
	defer cj.Stop()
	data, err := findProblemList([]string{})
	if err != nil {
		log.Fatal(err)
	}
	err = saveProblem(data.Result.Problems, "problems.json")
	if err != nil {
		log.Fatal(err)
	}
	for key, val := range divideProblems(data.Result.Problems) {
		err = saveProblem(val, key+".json")
		if err != nil {
			log.Fatal(err)
		}
	}
	if viper.GetString("handle") == "" {
		log.Fatal("config.yaml: handle is empty")
	}
	good := mashup.QueryStatus(viper.GetString("handle"))
	mashup.SaveStatus(good)
	no1.Done()
	no2 := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("user rating updating").SetComplete("user rating update complete")
	SaveRating(viper.GetString("handle"))
	no2.Done()
}

func findProblemList(tags []string) (problemList, error) {
	var tag = strings.Join(tags, ";")
	res, err := http.Get(problemApi + "?tags=" + tag)
	if err != nil {
		log.Fatalln(err)
		return problemList{}, err
	}
	var data problemList
	body, _ := ioutil.ReadAll(res.Body)
	err = json.Unmarshal(body, &data)
	if err != nil {
		log.Fatalln(err)
		return problemList{}, err
	}
	if data.Status != "OK" {
		return problemList{}, errors.New("request status is not OK")
	}
	return data, nil
}

func saveProblem(data []problemInfo, fileName string) error {
	os.Mkdir("./codeforces", 0777)
	file, err := os.OpenFile("./codeforces/"+fileName, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0777)
	if err != nil {
		return err
	}
	js, err := json.Marshal(data)
	if err != nil {
		return err
	}
	fmt.Fprint(file, string(js))
	return nil
}

func divideProblems(data []problemInfo) map[string][]problemInfo {
	var res = make(map[string][]problemInfo)
	for i := 0; i < len(data); i++ {
		var rating = strconv.Itoa(data[i].Rating)
		res[rating] = append(res[rating], data[i])
	}
	return res
}

type problemInfo struct {
	ContestId int      `json:"contestId"`
	Index     string   `json:"index"`
	Name      string   `json:"name"`
	Type      string   `json:"type"`
	Rating    int      `json:"rating"`
	Tags      []string `json:"tags"`
}
type problemStatistic struct {
	ContestId   int    `json:"contestId"`
	Index       string `json:"index"`
	SolvedCount int    `json:"solvedCount"`
}

type Result struct {
	Problems          []problemInfo      `json:"problems"`
	ProblemStatistics []problemStatistic `json:"problemStatistics"`
}
type problemList struct {
	Status string `json:"status"`
	Result Result `json:"result"`
}

func getRating(handle string) int {
	res, err := http.Get("https://codeforces.com/api/user.info?handles=" + handle)
	if err != nil {
		log.Fatal(err)
	}
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var info UserInfo
	err = json.Unmarshal(data, &info)
	if err != nil {
		log.Fatal(err)
	}
	return info.Result[0].Rating
}

type UserInfo struct {
	Status string           `json:"status"`
	Result []UserInfoResult `json:"result"`
}
type UserInfoResult struct {
	Rating int `json:"rating"`
}

func SaveRating(handle string) {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	viper.Set("rating", getRating(handle))
	viper.WriteConfig()
}
