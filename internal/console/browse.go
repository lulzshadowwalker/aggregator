package console

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/database"
	"github.com/lulzshadowwalker/aggregator/internal/posts"
)

type Browse struct {
	//
}

func (c Browse) Name() string {
	return "browse"
}

func (c Browse) Description() string {
	return "browse posts from feeds you follow"
}

func (c Browse) Handle(state *state, args []string, user database.User) (string, int, error) {
	limit := 2
	if len(args) == 1 {
		var err error
		limit, err = strconv.Atoi(args[0])
		if err != nil || limit <= 0 {
			return "", 1, errors.New("limit must be a positive integer")
		}
	} else if len(args) > 1 {
		return "", 1, errors.New("usage: browse [limit]")
	}

	postList, err := posts.Browse(context.Background(), state.database, user, limit)
	if err != nil {
		return "", 1, err
	}

	if len(postList) == 0 {
		return "No posts found", 0, nil
	}

	var str strings.Builder
	for _, post := range postList {
		fmt.Fprintf(&str, "Title:       %s\n", post.Title)
		fmt.Fprintf(&str, "URL:         %s\n", post.Url)
		if post.PublishedAt.Valid && post.PublishedAt.String != "" {
			fmt.Fprintf(&str, "PublishedAt: %s\n", post.PublishedAt.String)
		}
		if post.Description.Valid && post.Description.String != "" {
			fmt.Fprintf(&str, "Description: %s\n", post.Description.String)
		}
		fmt.Fprintf(&str, "========================================\n")
	}

	return strings.TrimRight(str.String(), "\n"), 0, nil
}
