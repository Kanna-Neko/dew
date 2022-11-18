package link

import (
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"sync"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/uispinner"
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
		go func(index int) {
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
		}(i)
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

func GetContestProblems(contestId string) []string {
	var res ContestStandingInterface
	response, err := me.R().Get("https://codeforces.com/api/contest.standings?contestId=" + contestId + "&from=1&handles=jaxleof&showUnofficial=true")
	if err != nil {
		log.Fatal(err)
	}
	json.Unmarshal(response.Body(), &res)
	if res.Status != "OK" {
		log.Fatal(res.Comment)
	}
	var problems []string
	for _, v := range res.Result.Problems {
		problems = append(problems, strconv.Itoa(v.ContestId)+v.Index)
	}
	return problems
}
