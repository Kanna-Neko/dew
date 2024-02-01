package main

import (
	"log"

	"github.com/Kanna-Neko/dew/cmd"
)

func main() {
	log.SetFlags(log.Flags() | log.Lshortfile)
	cmd.Execute()
}
