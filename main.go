package main

import (
	"github.com/webdelo/tradebot/cmd"
	"log"
)

func main() {
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}
