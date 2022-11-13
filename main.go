package main

import (
	"log"

	"github.com/jaxleof/dew/cmd"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	cmd.Execute()
}
