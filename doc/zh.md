# dew

语言： [en](../README.md)/zh

很高兴可以向你介绍dew，dew是一个可以帮助你测试并提交代码到codeforces的软件，并且它还可以随机生成比赛或题目帮助你更好的进行训练。

我推荐你使用`dew random` 命令， 它可以随机从题库中挑选一道符合你水平的题目，这道题目的分数是你的评分[+200,+300]

## 哪些人需要dew
1. 认为重复进行测试并且点击提交是件很麻烦的事情
2. 你需要为你或你的团队创建更多的训练

## 安装
我将介绍两种方式给你

### 下载二进制程序并将他添加在path中
  1. 打开 [releases](https://github.com/jaxleof/dew/releases)
  2. 下载压缩包
  3. 解压并将他加入到path中，如果你不知道什么是path，你可以尝试百度一下，或者询问你的学长学姐。

### 或者使用go install 命令
  1. 下载golang(https://golang.google.cn/)
  2. 使用命令 `go install github.com/jaxleof/dew@latest`

## 特性
|name|describe|
|---|---|
|跨平台|支持windows,macos,linux|
|多种语言e|理论上支持所有语言|
|快速提交以及测试|使用命令行，你只需要敲击几个单词，你就可以测试你的代码并进行提交|
|支持gym|不止是正式比赛，并且支持mushups和gym|
|模板生成|你可以存储你的模板并且在你需要他们的时候进行生成|
|为你随机选择题目|you can randomly pick one problem in codeforces, you can use it like a tool, everyday one problem你可以在codeforces中随机选择一道题目，你可以把它当作每日一题|
|比赛生成器|你可以为你自己或你的团队生成一场比赛在mushups中，其中的题目来源codeforces|

## Some snapshot
![](../snapshot/1.png)
![](../snapshot/2.png)
![](../snapshot/3.png)
![](../snapshot/4.png)


## 基础使用
```shell
# 放轻松，这很简单
dew init
# 原始的语言是c++，你可以通过使用dew lang [language]来改变软件使用的编程语言
# 你可以使用dew lang命令来查看所有支持的语言
# 编写你的程序，我们约定默认使用dew.cpp作为原始的提交代码文件
# 你可以修改这个名字在.codeforces/config.yaml，每个语言提交文件的名字可以不同。
# 我们将根据codefile来选择文件进行提交和测试
vim dew.cpp
dew test 1749A
dew submit 1749A

# 比赛模式
dew race 1749
vim dew.cpp
dew test A
dew submit A

# 随机题目
dew random
dew test
dew submit

# 特定一道题目
dew problem 1749A
dew test
dew submit

# 指定一个文件
dew test -f main.cpp
dew submit -f main.cpp
```

## dew command
```shell
Usage:
  dew [flags]
  dew [command]

Available Commands:
  completion  生成补充脚本
  env         输出配置环境
  generate    创建一场虚拟比赛
  help        获取更多帮助
  init        初始化
  lang        切换编程语言
  login       手动登陆 #你可能在令牌过期的时候使用到
  open        打开codeforces网站的快捷方式
  problem     在codeforces网站上打开一道题目
  race        设置比赛模式
  random      随机一道题目，是generate random命令的别名
  status      打开codeforces status的快捷方式
  submit      提交代码
  template    生产模板
  test        测试代码
  version     查看版本
  tutorial    在luogu中打开题解
  update      更新题目数据到本地

Flags:
  -h, --help   help for dew

Use "dew [command] --help" for more information about a command.
```

## dew lang
当你没有添加任何东西在lang后面时，这个命令将会显示当前你使用的语言，和所有你可以选择的语言

如果你想改变语言，你可以使用dew lang {language shortcut}

example:
```shell
dew lang
dew lang python3
```
### 如何自定义语言

1. 确保你的语言程序被你安装进环境变量里了
2. 让我们来看看配置文件，它的位置在 "./codeoforces.config.yaml" ,以下是python3的配置信息
```
python3: ## the shortcut, you can change to python3 when you use command "dew lang ，这是shortcut，你可以通过dew lang shortcut来改变你使用的语言
   name: Python 3.8.10 #语言名
   codefile: dew.py ## the file dew will test and submit
   isCompileLang: false 
   compileCommand:
   RunCommand: python3 $codefile ##dew将会在test命令中运行这个命令，$codefile是一个变量，dew将会在运行时自动替换掉他
   programTypeId: 31 ## 当你使用提交命令的使用，programTypeId决定了dew使用哪种语言提交，你可以在最下方查看所有的语言
```
## dew template
### 创建你自己的模板
1. 创建一个文件夹在 ./codeforces/template/{你的模板的名字}
2. 将所有模板文件填写进这个文件夹,例如 ./codeforces/template/{your template name}/{template file name}
3. 如果你想显示当前有多少模板可以选择，你可以使用`dew template -a`命令

### 生成模板

程序将生成所有存在于模板文件夹中的文件，例如当你使用`dew template hello`，所有在`./codeforces/hello/`中的文件将会被创建在当前文件夹。

如果你没有添加任何模板名，那么相当于你使用了`dew template default`

提示: 你可以同时使用超过一个模板名像`dew template hello1 hello2`，所有在`./codeforces/hello1/`中的文件将会被先创建出来，所有不同于前者但存在于`./codeforces/hello2/`中的文件将会被后创建出来

快捷命令是`dew tmp`

example:
```shell
dew template -a
dew template
# or 
dew tmp -a
dew tmp


# 请先创建hello模板
dew template hello

# 请先创建basic和gcd模板前
# 如果说两个模板存在相同的文件名，那么在后面的文件将不会被创建
dew template basic gcd
```

## dew random
random命令将会指定一道题目根据你现在的rating[+200,+300]，并且你可以使用一个具体的rating或者一个范围来随机题目。

例子:
```shell
dew random
dew random 1800
dew random 1000 1200
```

## dew generate
1. this command can generate a contest on mushups, I built three option in it for you, you can use `dew generate div1` `dew generate div2` `dew generate div3` to generate.
这个命令将会生成一个比赛在mushups中，我内置了三个选项在这个命令中， 你可以使用`dew generate div1` `dew generate div2` `dew generate div3` 来生成比赛。
2. If you hope custom difficult and filter some tag, you can use `dew generate custom` command, you can config it in "./codeforces/contestTemplate.json", dew will filter all problem satisfy thoes difficult, tag you choose in **good** field and without **bad** field
如果你想自定义难度并且筛选一些标签，你可以使用`dew generate custom` 命令， 你可以配置一些选项在"./codeforces/contestTemplate.json",dew 将会筛选出你所选择的难度并且所有tag满足good中的标签，并且不包含bad中的标签。
## 开发计划
[link](https://miaonei.notion.site/45b6802260cb479896640a06d521c99e?v=83fa5f001404427fa645aa5009ada702)
## 代理设置
```shell
# 在初始化之后
dew env -w proxy=http://127.0.0.1:41019
# 用你的代理服务替换 http://127.0.0.1:41019
```
## 技术栈
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)