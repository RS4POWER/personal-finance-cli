package main

import (
	"log"

	"github.com/RS4POWER/personal-finance-cli/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		log.Fatal(err)
	}
}
