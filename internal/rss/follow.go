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
	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Follow(ctx context.Context, db *database.Queries, url url.URL) (*database.Feed, error) {
	user, err := db.GetUser(ctx, config.Instance.Username)
	// this should be removed whenever we introduce some form of a middleware
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrUnauthorized
		}

		return nil, err
	}

	feed, err := db.GetFeedByURL(ctx, url.String())
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, fmt.Errorf("feed with url %q: %w", url.String(), ErrNotFound)
		}

		return nil, err
	}

	_, err = db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return nil, fmt.Errorf("you are already following this feed: %w", ErrConflict)
	}
	if err != nil {
		return nil, err
	}

	return &feed, nil
}
