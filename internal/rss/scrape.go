package rss

import (
	"context"
	"database/sql"
	"net/url"
	"time"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Scrape(ctx context.Context, db *database.Queries) (*Feed, *database.Feed, error) {
	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return nil, nil, err
	}

	now := time.Now().UTC()
	err = db.MarkFeedFetched(ctx, database.MarkFeedFetchedParams{
		ID: feed.ID,
		LastFetchedAt: sql.NullTime{
			Time:  now,
			Valid: true,
		},
	})
	if err != nil {
		return nil, nil, err
	}

	u, err := url.Parse(feed.Url)
	if err != nil {
		return nil, &feed, err
	}

	rssFeed, err := Fetch(ctx, *u)
	if err != nil {
		return nil, &feed, err
	}

	return rssFeed, &feed, nil
}
