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