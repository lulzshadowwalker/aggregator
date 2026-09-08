package console

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type Agg struct {
	//
}

func (c Agg) Name() string {
	return "agg"
}

func (c Agg) Description() string {
	return "aggregate feeds continuously"
}

func (c Agg) Handle(state *state, args []string) (string, int, error) {
	if len(args) != 1 {
		return "", 1, errors.New("usage: agg <time_between_reqs>")
	}

	interval, err := time.ParseDuration(args[0])
	if err != nil {
		return "", 1, err
	}
	if interval <= 0 {
		return "", 1, errors.New("time_between_reqs must be greater than 0")
	}

	fmt.Printf("Collecting feeds every %s\n", interval)

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		scrape(state)
	}
}

func scrape(state *state) {
	feed, count, err := rss.Scrape(context.Background(), state.database)
	if err != nil {
		fmt.Printf("Couldn't scrape feed: %v\n", err)
		return
	}

	fmt.Printf("Feed %s collected, %d posts found\n", feed.Name, count)
}

