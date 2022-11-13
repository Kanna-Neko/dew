# dew

A contest generator base on codeforces, which will help you better practice. you can use generate command a contest on mashup, whose problems is random

I suggest you use 'dew random' command, which will generate one problem whose difficulty base your rating [+200, +500] 

## Usage
```shell
dew init
# waiting for your exploration
vim main.cpp # write your program, we promise use main.cpp as submit code file.
dew test 1749A
dew submit 1749A

# race mode
dew race 1749
vim main.cpp
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
A contest generator base on codeforces, which will help you better practice. you can use generate command a contest on mashup, whose problems is random or custom

Usage:
  dew [flags]
  dew [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  env         print config env
  generate    create a contest
  help        Help about any command
  init        init somethings
  login       manually login
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

## Technology Stack
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)