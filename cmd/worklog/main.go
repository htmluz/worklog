package main

import (
	"log"

	"github.com/htmluz/worklog/internal/cli"
)

func main() {
	if err := cli.NewRootCmd().Execute(); err != nil {
		log.Fatal(err)
	}
}
