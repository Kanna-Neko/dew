package link

import (
	"bytes"
	"crypto/sha512"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
	"github.com/go-resty/resty/v2"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/viper"
)

func CloneContest(title string, id string, duration string) {
	cj := uispinner.New()
	cj.Start()
	defer cj.Stop()
	login := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("cloning").SetComplete("clone complete")
	_, err := me.R().SetFormData(map[string]string{
		"action":                 "saveMashup",
		"isCloneContest":         "true",
		"parentContestIdAndName": id,
		"parentContestId":        id,
		"contestName":            title,
		"contestDuration":        duration,
		"problemsJson":           "[]",
		"csrf_token":             csrf,
	}).Post(codeforcesDomain + "/data/mashup")
	if err != nil {
		log.Fatal(err)
	}
	login.Done()
}

func QueryProbelmId(problem string) (string, error) {
	var info ProblemInfos
	_, err := me.R().SetFormData(map[string]string{
		"action":                      "problemQuery",
		"problemQuery":                problem,
		"previouslyAddedProblemCount": "0",
		"csrf_token":                  csrf,
	}).SetResult(&info).Post(codeforcesDomain + "/data/mashup")
	if err != nil {
		log.Fatal(err)
	}
	if len(info.Problems) == 0 {
		log.Fatal(errors.New(problem + " isn't exist"))
	}
	return info.Problems[0].Id, nil
}

type ProblemInfo struct {
	EnglishName   string   `json:"englishName"`
	Id            string   `json:"id"`
	LocalizedName string   `json:"localizedName"`
	Rating        int      `json:"rating"`
	RussianName   string   `json:"russianName"`
	SolutionsUrl  string   `json:"solutionsUrl"`
	SolvedCount   int      `json:"solvedCount"`
	StatementUrl  string   `json:"statementUrl"`
	Tags          []string `json:"tags"`
}
type ProblemInfos struct {
	Problems []ProblemInfo `json:"problems"`
	Success  string        `json:"success"`
}

func CreateContest(title string, duration string, problems []string) {
	cj := uispinner.New()
	login := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("contest creating").SetComplete("contest create complete")
	cj.Start()
	var problemsJson = make([]ProblemJson, len(problems))
	var group = new(sync.WaitGroup)
	for i, problem := range problems {
		group.Add(1)
		x := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix(problem + " problem clawing").SetComplete(problem + " claw complete")
		go func(index int, problem string) {
			id, err := QueryProbelmId(problems[index])
			if err != nil {
				cj.Stop()
				log.Fatalln(err)
				return
			}
			problemsJson[index].Id = id
			problemsJson[index].Index = string(rune(('A' + index)))
			group.Done()
			x.Done()
		}(i, problem)
	}
	group.Wait()
	data, err := json.Marshal(problemsJson)
	if err != nil {
		log.Fatal(err)
	}
	_, err = me.R().SetFormData(map[string]string{
		"action":                 "saveMashup",
		"isCloneContest":         "false",
		"parentContestIdAndName": "",
		"parentContestId":        "",
		"contestName":            title,
		"contestDuration":        duration,
		"problemsJson":           string(data),
		"csrf_token":             csrf,
	}).Post(codeforcesDomain + "/data/mashup")
	if err != nil {
		log.Fatal(err)
	}
	login.Done()
	cj.Stop()
}

func GetContestInfo(contestId string) ContestStandingResult {
	var res ContestStandingInterface
	var response *resty.Response
	var err error
	if isGym(contestId) {
		apikey := viper.GetString("apikey")
		secret := viper.GetString("secret")
		if apikey == "" || secret == "" {
			log.Fatal("please use init apikey command first")
		}
		tim := time.Now().Unix()
		url := "contest.standings?apiKey=%s&contestId=%s&count=5&from=1&showUnofficial=true&time=%d"
		data := "987654/" + fmt.Sprintf(url, apikey,contestId, tim) + "#" + secret
		hash := sha512.Sum512([]byte(data))
		path := "https://codeforces.com/api/" + fmt.Sprintf(url, apikey,contestId, tim) + "&apiSig=987654" + hex.EncodeToString(hash[:])
		response, err = me.R().Get(path)
	} else {
		response, err = me.R().Get("https://codeforces.com/api/contest.standings?contestId=" + contestId + "&from=1&handles=jaxleof&showUnofficial=true")
	}
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response.Body(), &res)
	if res.Status != "OK" {
		log.Fatal(res.Comment)
	}
	return res.Result
}

func GetContestCountdown(contestId string) int {

	var res *resty.Response
	var err error
	if isGym(contestId) {
		res, err = me.R().Get(fmt.Sprintf("https://codeforces.com/gym/%s/countdown", contestId))
	} else {
		res, err = me.R().Get(fmt.Sprintf("https://codeforces.com/contest/%s/countdown", contestId))
	}
	if err != nil {
		log.Fatal(err)
	}
	var ioo = bytes.NewReader(res.Body())
	doc, err := goquery.NewDocumentFromReader(ioo)
	if err != nil {
		log.Fatal(err)
	}
	val, exist := doc.Find(".countdown>span").First().Attr("title")
	if !exist {
		val = doc.Find(".countdown").First().Text()
		if val == "" {
			return 0
		}
	}
	var slice = strings.Split(val, ":")
	s, err := strconv.Atoi(slice[2])
	if err != nil {
		log.Fatal(err)
	}
	m, err := strconv.Atoi(slice[1])
	if err != nil {
		log.Fatal(err)
	}
	h, err := strconv.Atoi(slice[0])
	if err != nil {
		log.Fatal(err)
	}
	return s + m*60 + h*60*60
}
