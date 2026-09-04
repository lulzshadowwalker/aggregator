package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

func Register(credentials Credentials, db *database.Queries) error {
	username := strings.TrimSpace(credentials.Username)
	if username == "" {
		return ErrInvalidCredentials
	}

	_, err := db.GetUser(context.Background(), credentials.Username)
	if err == nil {
		return ErrConflict
	} else if !errors.Is(err, sql.ErrNoRows) {
		return err
	}

	if _, err := db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		Name:      credentials.Username,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}); err != nil {
		return err
	}

	config.Instance.Username = username
	return config.Instance.Write()
}
