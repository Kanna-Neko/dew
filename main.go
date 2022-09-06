package main

import (
	"log"

	"github.com/jaxleof/cf-helper/cmd"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	cmd.Execute()
}
