package link

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/briandowns/spinner"
)

func SubmitCode(contest string, index string, code []byte, programTypeId string) {
	Login()
	sp := spinner.New(spinner.CharSets[43], 100*time.Millisecond)
	sp.Prefix = "AC!!! "
	sp.Start()
	sp.FinalMSG = "submission complete\n"
	var path string
	if isGym(contest) {
		path = fmt.Sprintf(codeforcesDomain+"/gym/%s/submit?csrf_token=%s", contest, csrf)
	} else {
		path = fmt.Sprintf(codeforcesDomain+"/contest/%s/submit?csrf_token=%s", contest, csrf)
	}
	res, err := me.R().SetFormData(map[string]string{
		"csrf_token":            csrf,
		"action":                "submitSolutionFormSubmitted",
		"contestId":             contest,
		"submittedProblemIndex": index,
		"programTypeId":         programTypeId,
		"source":                string(code),
		"tabSize":               "4",
	}).Post(path)
	if err != nil {
		log.Fatal(err)
	}
	sp.Stop()
	doc, err := goquery.NewDocumentFromReader(bytes.NewBuffer(res.Body()))
	if err != nil {
		log.Fatal(err)
	}
	errText := doc.Find(".error.for__source").First().Text()
	if errText != "" {
		log.Fatal(errText)
	}
	if err != nil {
		log.Fatal(err)
	}
}
