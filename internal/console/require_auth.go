package console

import (
	"context"
	"database/sql"
	"errors"

	"github.com/lulzshadowwalker/aggregator/internal/database"
)

type handler func(state *state, args []string) (string, int, error)

type authedHandler func(state *state, args []string, user database.User) (string, int, error)

var ErrUnauthorized = errors.New("you must be logged in to run this command")

func RequireAuth(next authedHandler) handler {
	return func(state *state, args []string) (string, int, error) {
		if state.config.Username == "" {
			return "", 1, ErrUnauthorized
		}

		user, err := state.database.GetUser(context.Background(), state.config.Username)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return "", 1, errors.New("user not found; please login again")
			}

			return "", 1, err
		}

		return next(state, args, user)
	}
}

type AuthedCommander interface {
	Name() string
	Description() string
	Handle(state *state, args []string, user database.User) (string, int, error)
}

type authedAdapter struct {
	command AuthedCommander
}

func (a authedAdapter) Name() string {
	return a.command.Name()
}

func (a authedAdapter) Description() string {
	return a.command.Description()
}

func (a authedAdapter) Handle(state *state, args []string) (string, int, error) {
	return RequireAuth(a.command.Handle)(state, args)
}

func WithAuth(command AuthedCommander) Commander {
	return authedAdapter{command: command}
}
