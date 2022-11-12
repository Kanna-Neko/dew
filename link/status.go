package link

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strconv"
)


func QueryStatus(handle string) map[string]bool {
	var status UsersStatusApi
	setProxy()
	_, err := me.R().SetResult(&status).Get(codeforcesDomain + "/api/user.status?handle=" + handle)
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
func saveStatus(data map[string]bool) {
	var res, err = json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./codeforces/myStatus.json", res, 0666)
	if err != nil {
		log.Fatal(err)
	}
}
func SaveStatus(handle string) {
	good := QueryStatus(handle)
	saveStatus(good)
}

func GetStatus() map[string]bool {
	var data, err = ioutil.ReadFile("./codeforces/myStatus.json")
	if err != nil {
		log.Fatal(err)
	}
	var res map[string]bool
	err = json.Unmarshal(data, &res)
	if err != nil {
		log.Fatal(err)
	}
	return res
}
