package mashup

import (
	"bytes"
	"fmt"
	"log"
	"time"

	"github.com/briandowns/spinner"

	"github.com/PuerkitoBio/goquery"
	resty "github.com/go-resty/resty/v2"
	"github.com/jaxleof/uispinner"
)

var (
	me       *resty.Client
	csrf     string
	handle   string
	password string
)

func init() {
	me = resty.New()
	me.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	me.SetContentLength(true)
	handle = "jaxleof"
	password = "xinxin6635"
}
func GetCsrf() {
	res, err := me.R().Get("https://codeforces.com/")
	if err != nil {
		log.Fatal(err)
	}
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		log.Fatal(err)
	}
	var exist bool
	csrf, exist = doc.Find(".csrf-token").First().Attr("data-csrf")
	if !exist {
		fmt.Println("obtain csrf failed")
		return
	}
	me.SetCookies(res.Cookies())
}

func Login() {
	cj := uispinner.New()
	cj.Start()
	defer cj.Stop()
	login := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("logining").SetComplete("login complete")
	defer login.Done()
	GetCsrf()
	res, err := me.R().SetQueryParams(map[string]string{
		"csrf_token":    csrf,
		"action":        "enter",
		"handleOrEmail": handle,
		"password":      password,
		"remember":      "on",
	}).Post("https://codeforces.com/enter?back=%2F")
	if err != nil {
		fmt.Println("login failed")
	}
	me.SetCookies(res.Cookies())
	GetCsrf()
	me.SetHeader("x-csrf-token", csrf)
}
