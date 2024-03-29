package link

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"time"

	"github.com/briandowns/spinner"
	"github.com/dlclark/regexp2"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/viper"

	"github.com/PuerkitoBio/goquery"
	resty "github.com/go-resty/resty/v2"
)

const codeforcesDomain = "https://codeforces.com"

var (
	me        *resty.Client
	csrf      string
	handle    string
	password  string
	cookieJar *cookiejar.Jar
	logined   bool = false
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
	setProxy()
	if checkLoginAgain() {
		loginAgain()
		logined = true
	} else {
		reloadCookie()
	}
}

func checkLoginAgain() bool {
	if logined {
		return false
	}
	if !viper.IsSet("cookie.csrf") || !viper.IsSet("cookie.JSESSIONID") || !viper.IsSet("cookie.39ce7") || !viper.IsSet("cookie.rcpc") {
		return true
	}
	t := viper.GetInt64("cookie.expire")
	return time.Now().Unix() > t
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
	rcpc, isExist := getRCPC()
	if isExist {
		viper.Set("cookie.rcpc", rcpc)
		me.SetCookie(&http.Cookie{
			Name:  "RCPC",
			Value: rcpc,
		})
	}
	getCsrf(codeforcesDomain + "/enter")
	_, err := me.R().SetFormData(map[string]string{
		"action":        "enter",
		"handleOrEmail": handle,
		"password":      password,
		"remember":      "on",
		"csrf_token":    csrf,
	}).Post(codeforcesDomain + "/enter?back")
	if err != nil {
		log.Fatal(err)
	}
	urL, _ := url.Parse(codeforcesDomain)
	for _, val := range cookieJar.Cookies(urL) {
		me.SetCookie(&http.Cookie{Name: val.Name, Value: val.Value})
		viper.Set("cookie."+val.Name, val.Value)
	}
	viper.Set("cookie.expire", time.Now().AddDate(0, 0, 1).Unix())
	if err != nil {
		log.Fatal(err)
	}
	getCsrf(codeforcesDomain)
	err = viper.WriteConfig()
	if err != nil {
		log.Fatal(err)
	}
}

func reloadCookie() {
	cook := viper.GetStringMapString("cookie")
	for k, v := range cook {
		me.SetCookie(&http.Cookie{
			Name:  k,
			Value: v,
		})
	}
	me.SetCookie(&http.Cookie{
		Name:  "39ce7",
		Value: viper.GetString("cookie.39ce7"),
	})
	me.SetCookie(&http.Cookie{
		Name:  "JSESSIONID",
		Value: viper.GetString("cookie.JSESSIONID"),
	})
	me.SetCookie(&http.Cookie{
		Name:  "RCPC",
		Value: viper.GetString("cookie.rcpc"),
	})
	csrf = viper.GetString("cookie.csrf")
}

func setProxy() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
	if viper.IsSet("proxy") {
		me.SetProxy(viper.GetString("proxy"))
	}
}
func getRCPC() (string, bool) {
	res, err := me.R().Get(codeforcesDomain + "/enter")
	if err != nil {
		log.Fatal(err)
	}
	if len(res.Body()) > 800 {
		return "", false
	}
	return decrypterRCPC(getParameter(string(res.Body()))), true
}
func getParameter(in string) (string, string, string) {
	res, err := regexp2.Compile(`(?<=a=toNumbers\(")\w+(?="\))`, 0)
	if err != nil {
		log.Fatal(err)
	}
	m, err := res.FindStringMatch(in)
	if err != nil {
		log.Fatal(err)
	}
	key := m.Group.Captures[0].String()
	res, err = regexp2.Compile(`(?<=b=toNumbers\(")\w+(?="\))`, 0)
	if err != nil {
		log.Fatal(err)
	}
	m, err = res.FindStringMatch(in)
	if err != nil {
		log.Fatal(err)
	}
	iv := m.Group.Captures[0].String()
	res, err = regexp2.Compile(`(?<=c=toNumbers\(")\w+(?="\))`, 0)
	if err != nil {
		log.Fatal(err)
	}
	m, err = res.FindStringMatch(in)
	if err != nil {
		log.Fatal(err)
	}
	cipher := m.Group.Captures[0].String()
	return cipher, key, iv
}

func decrypterRCPC(ocipher, okey, oiv string) string {
	key, _ := hex.DecodeString(okey)
	iv, _ := hex.DecodeString(oiv)
	ciphertext, _ := hex.DecodeString(ocipher)

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}
	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(ciphertext, ciphertext)
	return fmt.Sprintf("%x", ciphertext)
}

func getCsrf(path string) {
	res, err := me.R().Get(path)
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
		log.Fatal(string(res.Body()), "obtain csrf failed")
		return
	}
	viper.Set("cookie.csrf", csrf)
}
