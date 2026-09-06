package console

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/lulzshadowwalker/aggregator/internal/database"
	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type Unfollow struct {
	//
}

func (c Unfollow) Name() string {
	return "unfollow"
}

func (c Unfollow) Description() string {
	return "unfollow an existing feed"
}

func (c Unfollow) Handle(state *state, args []string, user database.User) (string, int, error) {
	if len(args) != 1 {
		return "", 1, errors.New("usage: unfollow <url>")
	}

	u, err := url.Parse(args[0])
	if err != nil {
		return "", 1, err
	}

	feed, err := rss.Unfollow(context.Background(), state.database, user, *u)
	if err != nil {
		return "", 1, err
	}

	return fmt.Sprintf("feed %q unfollowed successfully", feed.Name), 0, nil
}
