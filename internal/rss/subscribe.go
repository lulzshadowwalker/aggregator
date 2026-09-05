package rss

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lib/pq"
	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Subscribe(ctx context.Context, db *database.Queries, name string, url url.URL) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name cannot be empty")
	}

	user, err := db.GetUser(ctx, config.Instance.Username)
	if err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrUnauthorized
		}

		return err
	}

	_, err = db.CreateFeed(ctx, database.CreateFeedParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		Name:      name,
		Url:       url.String(),
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})

	var pqErr *pq.Error
	if errors.As(err, &pqErr) && pqErr.Code == "23505" {
		return fmt.Errorf("rss feed with the same url already exists: %w", ErrConflict)
	}
	if err != nil {
		return err
	}

	_, err = Follow(ctx, db, url)
	return err
}
