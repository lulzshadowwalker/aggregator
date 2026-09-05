package console

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type Follow struct {
	//
}

func (c Follow) Name() string {
	return "follow"
}

func (c Follow) Description() string {
	return "follow an existing feed"
}

func (c Follow)  Handle(state *state, args []string) (string, int, error) {
	if len(args) != 1 {
		return "", 1, errors.New("usage: follow <url>")
	}

	u, err := url.Parse(args[0])
	if err != nil {
		return "", 1, err
	}

	feed, err := rss.Follow(context.Background(), state.database, *u)
	if err != nil {
		return "", 1, err
	}

	return fmt.Sprintf("feed %q followed successfully", feed.Name), 0, nil
}
