package auth

import (
	"errors"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/config"
)

type LoginParams struct {
	Username string
}

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

func Login(params LoginParams) error {
	username := strings.TrimSpace(params.Username)
	if username == "" {
		return ErrInvalidCredentials
	}

	config.Instance.Username = username
	return config.Instance.Write()
}

