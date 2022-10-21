package mashup

import (
	"bytes"
	"log"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/viper"

	"github.com/PuerkitoBio/goquery"
	resty "github.com/go-resty/resty/v2"
)

var (
	me       *resty.Client
	csrf     string
	handle   string
	password string
)

func init() {
	me = resty.New()
	// me.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	me.SetContentLength(true)
}

func Login() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if !viper.IsSet("handle") {
		log.Fatal("handle info is empty\n you can use cf init config first")
	}
	if !viper.IsSet("password") {
		log.Fatal("password info is empty\n you can use cf config first")
	}
	handle = viper.GetString("handle")
	password = viper.GetString("password")
	cj := uispinner.New()
	cj.Start()
	defer cj.Stop()
	login := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("logining").SetComplete("login complete")
	defer login.Done()
	res, err := me.R().SetFormData(map[string]string{
		"action":        "enter",
		"handleOrEmail": handle,
		"password":      password,
	}).Post("https://codeforces.com/enter?back=%2F")
	if err != nil {
		log.Fatal("login failed")
	}
	me.SetCookies(res.Cookies())
	doc, err := goquery.NewDocumentFromReader(bytes.NewReader(res.Body()))
	if err != nil {
		log.Fatal(err)
	}
	var exist bool
	csrf, exist = doc.Find(".csrf-token").First().Attr("data-csrf")
	if !exist {
		log.Fatal("obtain csrf failed")
		return
	}
}
