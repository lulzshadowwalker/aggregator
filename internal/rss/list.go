package rss

import (
	"context"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func List(ctx context.Context, db *database.Queries) ([]database.GetFeedsRow, error) {
	return db.GetFeeds(ctx)
}
