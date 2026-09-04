package auth

import (
	"errors"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/config"
)

type Credentials struct {
	Username string
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func Login(credentials Credentials) error {
	username := strings.TrimSpace(credentials.Username)
	if username == "" {
		return ErrInvalidCredentials
	}

	config.Instance.Username = username
	return config.Instance.Write()
}

