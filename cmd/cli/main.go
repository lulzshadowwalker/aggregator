package main

import (
	"fmt"
	"os"

	"github.com/lulzshadowwalker/aggregator/internal/config"
)

func main() {
	fmt.Println("--- Before ---")
	fmt.Println("Database URL:", config.Instance.Database.URL, "Username:", config.Instance.Username)

	config.Instance.Username = "John Doe"
	if err := config.Instance.Write(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if err := config.Instance.Read(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("--- After ---")
	fmt.Println("Database URL:", config.Instance.Database.URL, "Username:", config.Instance.Username)

	fmt.Println("welcome, to the aggregator cli!")
}
