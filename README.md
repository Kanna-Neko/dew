# cf-helper

A contest generator base on codeforces, which will help you better practice. you can use cf generate a contest on mashup, whose problems is random or custom

## cf command
```shell
Usage:
  cf [flags]
  cf [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  config      config codeforces handle and password
  help        Help about any command
  init        init somethings
  new         create a contest
  update      update problem data

Flags:
  -h, --help   help for cf

Use "cf [command] --help" for more information about a command.
```

## cf new command
```shell
create a contest

Usage:
  cf new [flags]
  cf new [command]

Available Commands:
  div1        create a contest, whose difficulty like div1
  div2        create a contest, whose difficulty like div2
  div3        create a contest, whose difficulty like div3

Flags:
  -h, --help   help for new

Use "cf new [command] --help" for more information about a command.
```

## Technology Stack
1. [cobra](https://github.com/spf13/cobra)
2. [resty](https://github.com/go-resty/resty)