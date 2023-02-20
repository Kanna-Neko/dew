package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"time"

	"github.com/briandowns/spinner"
	"github.com/jaxleof/dew/lang"
	"github.com/jaxleof/uispinner"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const defaultCpp = `#include <iostream>
#include <cstdio>
#include <cstring>
#include <algorithm>
#include <string>
#include <vector>
#include <queue>
#include <set>
#include <map>
#include <stack>
#include <cmath>
using namespace std;
#define ll long long
#define inf 0x3f3f3f3f
template <class T> void gmin(T &a, T b) {
    if (a > b) a = b;
}
template <class T> void gmax(T &a, T b) {
    if (a < b) a = b;
}
const ll mod = 998244353;
int n,m;
int T;
const int maxn = 200005;
void solve() {

}
int main() {
    ios::sync_with_stdio(false);
    cin.tie(0);
    // freopen("input.in","r",stdin);
    cin >> T; 
    while(T--) {
        solve();
    }
    return 0;
}
`
const templateDir = "./codeforces/template/"
const testFilesDir = "./codeforces/testFiles/"

func init() {
	rootCmd.AddCommand(InitCmd)
	InitCmd.AddCommand(InitApikeyCmd)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init somethings",
	Run: func(cmd *cobra.Command, args []string) {
		configFunc()
		updateFunc()
	},
}
var InitApikeyCmd = &cobra.Command{
	Use:   "apikey",
	Short: "init apikey and secret",
	PreRun: func(cmd *cobra.Command, args []string) {
		ReadConfig()
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("please type the apikey")
		var apikey string
		fmt.Scanf("%s", &apikey)
		fmt.Println("please type secret")
		var secret string
		fmt.Scanf("%s", &secret)
		apikey = strings.Trim(apikey, " ")
		secret = strings.Trim(secret, " ")
		viper.Set("apikey", apikey)
		viper.Set("secret", secret)
		err := viper.WriteConfig()
		if err != nil {
			log.Fatal(err)
		}
	},
}

func configFunc() {
	if !checkConfigFile() {
		os.Mkdir("codeforces", 0777)
		os.Create("./codeforces/config.yaml")
	}
	ReadConfig()
	fmt.Println("please input your codeforces handle(account)")
	var handle string
	fmt.Scanln(&handle)
	fmt.Println("please input your codeforces password")
	var password string
	fmt.Scanln(&password)
	fmt.Printf("your handle is %s\nyour password is %s\n", handle, password)
	viper.Set("handle", handle)
	viper.Set("password", password)
	initLang()
	viper.WriteConfig()
	fmt.Println("handle and password save into ./codeforces/config.yaml")
	initTmplate()
	initTestManager()
	initContestTemplate()
	cj := uispinner.New()
	defer cj.Stop()
	no1 := cj.AddSpinner(spinner.CharSets[34], 100*time.Millisecond).SetPrefix("user rating is initing").SetComplete("user rating init complete")
	defer no1.Done()
	cj.Start()
	SaveRating(handle)
}

func ReadConfig() {
	viper.SetConfigFile("./codeforces/config.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("please use init command first, error : %s\n", err)
	}
}

func initLang() {
	viper.Set("lang", "c++")
	for _, v := range lang.LangSlice {
		viper.Set("language."+v+".isCompileLang", lang.OriginLangDic[v].IsComplieLang)
		viper.Set("language."+v+".compileCommand", lang.OriginLangDic[v].CompileCommand)
		viper.Set("language."+v+".runCommand", lang.OriginLangDic[v].RunCommand)
		viper.Set("language."+v+".programTypeId", lang.OriginLangDic[v].ProgramTypeId)
		viper.Set("language."+v+".codefile", lang.OriginLangDic[v].Codefile)
		viper.Set("language."+v+".Name", lang.OriginLangDic[v].Name)
	}
}

func initTmplate() {
	err := os.MkdirAll(templateDir+"default", 0777)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile(templateDir+"default/dew.cpp", []byte(defaultCpp), 0777)
	if err != nil {
		log.Fatal(err)
	}
}
func initTestManager() {
	err := os.MkdirAll(testFilesDir, 0777)
	if err != nil {
		log.Fatal(err)
	}
}

func initContestTemplate() {
	var template = ContestInfos{
		Templates: []ContestInfo{{
			Name:              "example",
			Duration:          "120",
			ContestTitle:      "Do you like cat?",
			ProblemConditions: div2Diffculty,
			BanProblems:       []string{"1A", "12B"},
		}},
	}
	data, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./codeforces/contestTemplate.json", data, 0777)
	if err != nil {
		log.Fatal(err)
	}
}
