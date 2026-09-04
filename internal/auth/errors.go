package auth

import "errors"

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
	ErrConflict           = errors.New("user already exists")
)
