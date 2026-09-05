package users

import (
	"context"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func List(db *database.Queries) ([]database.User, error) {
	return db.GetUsers(context.Background())
}
