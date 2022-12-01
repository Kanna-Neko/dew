# dew

language: en/[zh](./doc/zh.md)

dew is a assistant which can help you test your program you will submit on codeforces. and a contest generator base on codeforces, which will help you better practice. you can use generate command a contest on mashup, whose problems is random.

I suggest you use 'dew random' command, which will generate one problem whose difficulty base your rating [+200, +300] 

## basic Usage
```shell
# relax yourself and enjoy it, it's easy.
dew init
# waiting for your exploration
# original language is c++, you can use dew lang [language] change language program use.
# you can use dew lang command and check all language we support.
# write your program, we promise use dew.cpp as original submit code file.
# you can rewrite codefile in ./codeforces/config.yaml for different language
# codefile is code program will test and submit
vim dew.cpp
dew test 1749A
dew submit 1749A

# race mode
dew race 1749
vim dew.cpp
dew test A
dew submit A

# random problem
dew random
dew test
dew submit

# specify a problem
dew problem 1749A
dew test
dew submit

# specify a file
dew test -f main.cpp
dew submit -f main.cpp
```

## dew command
```shell
Usage:
  dew [flags]
  dew [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  env         print config env
  generate    create a contest
  help        Help about any command
  init        init somethings
  lang        switch program language
  login       manually login #maybe you will use when token is expired
  open        a shortcut of opening codeforces website
  problem     open problem in codeforces
  race        set contest env
  random      alias to dew generate random
  status      a shortcut of opening codeforces status
  submit      submit problem
  template    generate template
  test        test problem
  tutorial    as the name says
  update      update problem data

Flags:
  -h, --help   help for dew

Use "dew [command] --help" for more information about a command.
```

## dew lang
This command will show current and total choice you can choose when nothing after command

If you want to change language, you can use dew lang {language shortcut}

example:
```shell
dew lang
dew lang python3
```

And you can custom lang in the ./codeforces/config.yaml
## dew template
### create your template
1. mkdir in ./codeforces/template/{your template name}
2. You can write your template into ./codeforces/template/{your template name}/{template file name}
3. The program will show all template you create when you use -a flag `dew template -a`

### generate template
The program will create all template file in your template dir when you use template command, for instance `dew template hello`, all file in `./codeforces/hello/` will be created in current path.

If you don't fill anything after dew template, It's alias `dew template default`

hint:You can fill in more than one template name after the command like `dew template hello1 hello2`, all file in `./codeforces/hello1/` will be created first, distinct filename in `./codeforces/hello2/` will be created next.

The shortcut is `dew tmp`

example:
```shell
dew template -a
dew template
# or 
dew tmp -a
dew tmp


# please create template hello first.
dew template hello

# please create template basic and gcd first.
# if there are some filename is coincident same, files that appear after will not be created.
dew template basic gcd
```

## dew random
random command will specify a problem for your current rating \[+200, +300\], and you can specify a detailed rating or range rating you want.

example:
```shell
dew random
dew random 1800
dew random 1000 1200
```

## proxy set
```shell
# after dew init
dew env -w proxy=http://127.0.0.1:41019
# replace http://127.0.0.1:41019 with your proxy server
```
## Technology Stack
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)

## programTypeId

```
<43>GNU GCC C11 5.1.0
<80>Clang++20 Diagnostics
<52>Clang++17 Diagnostics
<50>GNU G++14 6.4.0
<54>GNU G++17 7.3.0
<73>GNU G++20 11.2.0 (64 bit, winlibs)
<59>Microsoft Visual C++ 2017
<61>GNU G++17 9.2.0 (64 bit, msys 2)
<65>C# 8, .NET Core 3.1
<79>C# 10, .NET SDK 6.0
<9>C# Mono 6.8
<28>D DMD32 v2.091.0
<32>Go 1.19
<12>Haskell GHC 8.10.1
<60>Java 11.0.6
<74>Java 17 64bit
<36>Java 1.8.0_241
<77>Kotlin 1.6.10
<83>Kotlin 1.7.20
<19>OCaml 4.02.1
<3>Delphi 7
<4>Free Pascal 3.0.2
<51>PascalABC.NET 3.4.2
<13>Perl 5.20.1
<6>PHP 8.1.7
<7>Python 2.7.18
<31>Python 3.8.10
<40>PyPy 2.7.13 (7.3.0)
<41>PyPy 3.6.9 (7.3.0)
<70>PyPy 3.9.10 (7.3.9, 64bit)
<67>Ruby 3.0.0
<75>Rust 1.65.0 (2021)
<20>Scala 2.12.8
<34>JavaScript V8 4.8.0
<55>Node.js 12.16.3
<14>ActiveTcl 8.5
<15>Io-2008-01-07 (Win32)
<17>Pike 7.8
<18>Befunge
<22>OpenCobol 1.0
<25>Factor
<26>Secret_171
<27>Roco
<33>Ada GNAT 4
<38>Mysterious Language
<39>FALSE
<44>Picat 0.9
<45>GNU C++11 5 ZIP
<46>Java 8 ZIP
<47>J
<56>Microsoft Q#
<57>Text
<62>UnknownX
<68>Secret 2021
```