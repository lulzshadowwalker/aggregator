package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Unfollow(ctx context.Context, db *database.Queries, user database.User, url url.URL) (*database.Feed, error) {
	feed, err := db.GetFeedByURL(ctx, url.String())
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("feed with url %q: %w", url.String(), ErrNotFound)
		}

		return nil, err
	}

	err = db.DeleteFeedFollow(ctx, feed.ID)
	if err != nil {
		// it does not error out when you try to unfollow a feed you are not already following 
		return nil, err
	}

	// for now, I don't think we should delete the feed itself since other users might be following it

	return &feed, nil
}
