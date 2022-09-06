# cf-helper

A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random

I suggest you use 'cf random' command, which will generate one problem whose difficulty base your rating [+200, +500] 

## Usage
```shell
cf init
cf random
cf generate div2
```

## cf command
```shell
Usage:
  cf [flags]
  cf [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    create a contest
  help        Help about any command
  info        print config info
  init        init somethings
  open        a shortcut of opening codeforces website
  problem     open problem in codeforces
  random      alias to cf generate random
  tutorial    as the name says
  update      update problem data

Flags:
  -h, --help   help for cf

Use "cf [command] --help" for more information about a command.
```

## cf new command
```shell
create a contest
Usage:
  cf generate [flags]
  cf generate [command]

Available Commands:
  div1        create a contest, whose difficulty like div1
  div2        create a contest, whose difficulty like div2
  div3        create a contest, whose difficulty like div3
  random      random select one problem
```

## Technology Stack
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)