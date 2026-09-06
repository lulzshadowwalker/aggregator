package rss

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Following(ctx context.Context, db *database.Queries, user database.User) ([]database.GetFeedsByUserIDRow, error) {
	feeds, err := db.GetFeedsByUserID(ctx, user.ID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return make([]database.GetFeedsByUserIDRow, 0), nil
		}

		return nil, err
	}

	return feeds, nil
}
