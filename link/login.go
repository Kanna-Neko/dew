package link

import (
	"bytes"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/viper"

	"github.com/PuerkitoBio/goquery"
	resty "github.com/go-resty/resty/v2"
)

const codeforcesDomain = "https://www.codeforces.com"

var (
	me        *resty.Client
	csrf      string
	handle    string
	password  string
	cookieJar *cookiejar.Jar
)

func init() {
	me = resty.New()
	cookieJar, _ = cookiejar.New(nil)
	me.SetCookieJar(cookieJar)
	me.SetHeader("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/103.0.0.0 Safari/537.36")
	me.SetContentLength(true)
}
func ManuallyLogin() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	loginAgain()
}

func Login() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if checkLoginAgain() {
		loginAgain()
	} else {
		reloadCookie()
	}
}

func checkLoginAgain() bool {
	t := viper.GetInt64("cookie.expire")
	if time.Now().Unix() > t {
		return true
	}
	if !viper.IsSet("cookie.csrf") || !viper.IsSet("cookie.JSESSIONID") || !viper.IsSet("cookie.39ce7") {
		return true
	}
	return false
}

func loginAgain() {
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
		"remember":      "on",
	}).Post(codeforcesDomain + "/enter?back=%2F")
	if err != nil {
		log.Fatal(err)
	}
	urL, _ := url.Parse(codeforcesDomain)
	for _, val := range cookieJar.Cookies(urL) {
		if val.Name == "39ce7" {
			viper.Set("cookie.39ce7", val.Value)
		} else if val.Name == "JSESSIONID" {
			viper.Set("cookie.JSESSIONID", val.Value)
		}
	}
	viper.Set("cookie.expire", time.Now().AddDate(0, 0, 29).Unix())
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
		log.Fatal("obtain csrf failed")
		return
	}
	viper.Set("cookie.csrf", csrf)
	err = viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func reloadCookie() {
	me.SetCookie(&http.Cookie{
		Name:  "39ce7",
		Value: viper.GetString("cookie.39ce7"),
	})
	me.SetCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: viper.GetString("cookie.JSESSIONID"),
	})
	csrf = viper.GetString("cookie.csrf")
}
