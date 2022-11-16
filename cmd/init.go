package cmd

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
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

func init() {
	rootCmd.AddCommand(InitCmd)
	InitCmd.AddCommand(InitDefaultCpp)
}

var InitCmd = &cobra.Command{
	Use:   "init",
	Short: "init somethings",
	Run: func(cmd *cobra.Command, args []string) {
		configFunc()
		updateFunc()
		setDefaultCpp("./codeforces/default.cpp")
	},
}
var InitDefaultCpp = &cobra.Command{
	Use:   "cpp",
	Short: "init default cpp file",
	Run: func(cmd *cobra.Command, args []string) {
		setDefaultCpp("./codeforces/default.cpp")
	},
}

func setDefaultCpp(path string) {
	err := ioutil.WriteFile(path, []byte(defaultCpp), 0666)
	if err != nil {
		log.Fatal(err)
	}
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
	for k, v := range lang.LangDic {
		viper.Set("codefile."+k, v.OriginalCodefile)
	}
}
