package cmd

import (
	"io/ioutil"
	"log"

	"github.com/spf13/cobra"
)

const defaultCpp = `
#include <iostream>
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
#define ls rt << 1
#define rs rt << 1 | 1
template <class T>
void gmin(T &a, T b)
{
    if (a > b)
        a = b;
}
template <class T>
void gmax(T &a, T b)
{
    if (a < b)
        a = b;
}
const ll mod = 998244353;
inline ll gcd(ll a, ll b)
{
    ll r;
    while (b > 0)
    {
        r = a % b;
        a = b;
        b = r;
    }
    return a;
}

inline ll lcm(ll a, ll b)
{
    return a * b / (gcd(a, b));
}
ll ksm(ll x, ll k)
{
    ll res = 1;
    while (k)
    {
        if (k & 1)
            res = 1LL * res * x % mod;
        x = 1LL * x * x % mod;
        k >>= 1;
    }
    return res % mod;
}
int n,m;
int T;
const int maxn = 150005;
int main() {
	
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
