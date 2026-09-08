package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func nullString(s string) sql.NullString {
	if s == "" {
		return sql.NullString{Valid: false}
	}

	return sql.NullString{
		String: s,
		Valid:  true,
	}
}

func Scrape(ctx context.Context, db *database.Queries) (*database.Feed, int, error) {
	feed, err := db.GetNextFeedToFetch(ctx)
	if err != nil {
		return nil, 0, err
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
		return nil, 0, err
	}

	u, err := url.Parse(feed.Url)
	if err != nil {
		return &feed, 0, err
	}

	rssFeed, err := Fetch(ctx, *u)
	if err != nil {
		return &feed, 0, err
	}

	for _, item := range rssFeed.Channel.Items {
		itemCreatedAt := time.Now().UTC()
		_, err := db.CreatePost(ctx, database.CreatePostParams{
			ID:          uuid.New(),
			FeedID:      feed.ID,
			Url:         item.Link,
			Title:       item.Title,
			Description: nullString(item.Description),
			PublishedAt: nullString(item.PublicationDate),
			CreatedAt:   itemCreatedAt,
			UpdatedAt:   itemCreatedAt,
		})
		if err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" {
				continue
			}

			fmt.Printf("Couldn't create post %q: %v\n", item.Title, err)
		}
	}

	return &feed, len(rssFeed.Channel.Items), nil
}
