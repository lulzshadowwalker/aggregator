package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/lulzshadowwalker/aggregator/internal/console"
)

func main() {
	output, err := console.New().Run(os.Args[1:])
	if err != nil {
		if errors.Is(console.ErrHelp, err) {
			fmt.Println(output)
		} else {
			fmt.Println(err)
		}

		os.Exit(1)
	}

	fmt.Println(output)
}
