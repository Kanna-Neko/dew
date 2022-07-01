package main

import (
	"cf/cmd"
	"log"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	cmd.Execute()
}
