# cf-helper

A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random

I suggest you use 'cf random' command, which will generate one problem whose difficulty base your rating [+200, +500] 

## Usage
```shell
cf init
# waiting for your exploration
vim main.cpp # write your program, we promise use main.cpp as submit code file.
cf test 1749A
cf submit 1749A

# race mode
cf race 1749
vim main.cpp
cf test A
cf submit A

# random problem
cf random
cf test
cf submit

# specify a problem
cf problem 1749A
cf test
cf submit
```

## cf command
```shell
A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random or custom

Usage:
  cf [flags]
  cf [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    create a contest
  help        Help about any command
  info        print config info
  init        init somethings
  login       manually login
  open        a shortcut of opening codeforces website
  problem     open problem in codeforces
  race        set contest env
  random      alias to cf generate random
  submit      submit problem
  test        test problem
  tutorial    as the name says
  update      update problem data

Flags:
  -h, --help   help for cf

Use "cf [command] --help" for more information about a command.
```

## Technology Stack
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)