package link

import (
	"fmt"
	"log"

	"github.com/PuerkitoBio/goquery"
)

func GetSample(contest string, index string) ([]string, []string) {
	setProxy()
	res, err := me.R().SetDoNotParseResponse(true).Get(codeforcesDomain + "/contest/" + contest + "/problem/" + index)
	if err != nil {
		log.Fatal(err)
	}
	if !res.IsSuccess() {
		log.Fatal("request error: " + res.Status())
	}
	defer res.RawBody().Close()
	doc, err := goquery.NewDocumentFromReader(res.RawBody())
	if err != nil {
		log.Fatal(err)
	}
	var input []string
	var output []string
	if doc.Find(".sample-test>.input>pre>div").Length() == 0 {
		doc.Find(".sample-test>.input").Each(func(i int, dom *goquery.Selection) {
			dom.Find("pre>br").ReplaceWithHtml("\n")
			input = append(input, dom.Find("pre").Text())
		})
	} else {
		doc.Find(".sample-test>.input").Each(func(i int, dom *goquery.Selection) {
			var sam string
			dom.Find("pre>br").ReplaceWithHtml("\n")
			dom.Find("pre>.test-example-line").Each(func(i int, s *goquery.Selection) {
				sam += fmt.Sprintf("%s\n", s.Text())
			})
			input = append(input, sam)
		})
	}
	doc.Find(".sample-test>.output").Each(func(i int, dom *goquery.Selection) {
		dom.Find("pre>br").ReplaceWithHtml("\n")
		output = append(output, dom.Find("pre").Text())
	})
	return input, output
}
