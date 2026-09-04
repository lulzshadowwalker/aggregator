package auth

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

type Credentials struct {
	Username string
}

func Login(credentials Credentials, db *database.Queries) error {
	username := strings.TrimSpace(credentials.Username)
	if _, err := db.GetUser(context.Background(), credentials.Username); err != nil {
		if errors.Is(sql.ErrNoRows, err) {
			return ErrInvalidCredentials
		}

		return err
	}

	config.Instance.Username = username
	return config.Instance.Write()
}
