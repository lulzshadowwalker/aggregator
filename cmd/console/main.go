package main

import (
	"fmt"
	"os"

	"github.com/lulzshadowwalker/aggregator/internal/console"
)

func main() {
	output, status, err := console.New().Run(os.Args[1:])
	if err != nil {
		fmt.Println(err)
	}

	if output != "" {
		fmt.Println(output)
	}

	os.Exit(status)
}
