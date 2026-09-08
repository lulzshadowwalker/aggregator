package console

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type RSS struct {
	//
}

func (c RSS) Name() string {
	return "rss"
}

func (c RSS) Description() string {
	return "fetch rss feed from <urL>"
}

func (c RSS) Handle(state *state, args []string) (string, int, error) {
	if len(args) != 1 {
		return "", 1, errors.New("usage: rss <url>")
	}

	u, err := url.Parse(args[0])
	if err != nil {
		return "", 1, err
	}
	feed, err := rss.Fetch(context.Background(), *u)
	if err != nil {
		return "", 1, err
	}

	return fmt.Sprintf("%#v", feed), 0, nil
}
