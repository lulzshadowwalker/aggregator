package console

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

type Feeds struct {
	//
}

func (c Feeds) Name() string {
	return "feeds"
}

func (c Feeds) Description() string {
	return "list all feeds"
}

func (c Feeds) Handle(state *state, args []string) (string, int, error) {
	if len(args) > 0 {
		return "", 1, errors.New("usage: feeds")
	}

	feeds, err := state.database.GetFeeds(context.Background())
	if err != nil {
		return "", 1, err
	}

	if len(feeds) == 0 {
		return "No feeds found", 0, nil
	}

	var users []string
	groups := make(map[string][]database.GetFeedsRow)
	for _, feed := range feeds {
		if _, ok := groups[feed.UserName]; !ok {
			users = append(users, feed.UserName)
		}
		groups[feed.UserName] = append(groups[feed.UserName], feed)
	}

	var str strings.Builder
	str.WriteString("Feeds:\n")
	for _, user := range users {
		suffix := ":\n"
		if user == config.Instance.Username {
			suffix = " (current):\n"
		}
		fmt.Fprintf(&str, "\t* %s%s", user, suffix)

		for _, feed := range groups[user] {
			fmt.Fprintf(&str, "\t\t* Name: %s\n", feed.Name)
			fmt.Fprintf(&str, "\t\t  URL:  %s\n", feed.Url)
		}
	}

	return strings.TrimRight(str.String(), "\n"), 0, nil
}
