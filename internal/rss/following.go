package rss

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Following(ctx context.Context, db *database.Queries) ([]database.GetFeedsByUserIDRow, error) {
	user, err := db.GetUser(ctx, config.Instance.Username)
	// this should be removed whenever we introduce some form of a middleware
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return nil, ErrUnauthorized
		}

		return nil, err
	}

	feeds, err := db.GetFeedsByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return make([]database.GetFeedsByUserIDRow, 0), nil
		}

		return nil, err
	}

	return feeds, nil
}
