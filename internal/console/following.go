package console

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type Following struct {
	//
}

func (c Following) Name() string {
	return "following"
}

func (c Following) Description() string {
	return "list all of the feeds you follow"
}

func (c Following) Handle(state *state, args []string) (string, int, error) {
	if len(args) > 0 {
		return "", 1, errors.New("usage: following")
	}

	feeds, err := rss.Following(context.Background(), state.database)
	if err != nil {
		return "", 1, err
	}

	if len(feeds) == 0 {
		return "No feeds found", 0, nil
	}

	var str strings.Builder
	str.WriteString("Feeds:\n")
	for _, feed := range feeds {
		fmt.Fprintf(&str, "\t* %s: %q\n", feed.FeedName, feed.FeedUrl)
	}

	return str.String(), 0, nil
}
