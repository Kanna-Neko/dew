package mashup

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

type UsersStatusApi struct {
	Status string   `json:"status"`
	Result []Status `json:"result"`
}
type Status struct {
	Problem Problem `json:"problem"`
	// Id                  int     `json:"id"`
	// ContestId           int     `json:"contestId"`
	// CreationTimeSeconds int     `json:"creationTimeSeconds"`
	// RelativeTimeSeconds int     `json:"relativeTimeSeconds"`
	// Author              any     `json:"author"`
	// ProgrammingLanguage string  `json:"programmingLanguage"`
	Verdict string `json:"verdict"`
	// Testset             string  `json:"testset"`
	// PassedTestCount     int     `json:"passedTestCount"`
	// TimeConsumedMillis  int     `json:"timeConsumedMillis"`
	// MemoryConsumedBytes int     `json:"memoryConsumedBytes"`
}
type Problem struct {
	ContestId int    `json:"contestId"`
	Index     string `json:"index"`
	// Name      string   `json:"name"`
	// Type      string   `json:"type"`
	// Points    float64  `json:"points"`
	// Rating    int      `json:"rating"`
	// Tags      []string `json:"tags"`
}

func QueryStatus(handle string) map[string]bool {
	res, err := http.Get("https://codeforces.com/api/user.status?handle=" + handle)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	var status UsersStatusApi
	err = json.Unmarshal(body, &status)
	if err != nil {
		log.Fatal(err)
	}
	var cj = make(map[string]bool)
	for i := 0; i < len(status.Result); i++ {
		if status.Result[i].Verdict == "OK" {
			cj[strconv.Itoa(status.Result[i].Problem.ContestId)+status.Result[i].Problem.Index] = true
		}
	}
	return cj
}
func SaveStatus(data map[string]bool) {
	var res, err = json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./codeforces/myStatus.json", res, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
