package posts

import (
	"context"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Browse(ctx context.Context, db *database.Queries, user database.User, limit int) ([]database.Post, error) {
	return db.GetPostsForUser(ctx, database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
}
