package console

import (
	"context"
	"fmt"
)

type Reset struct {
	//
}

func (c Reset) Name() string {
	return "reset"
}

func (c Reset) Description() string {
	return "truncate all users"
}

func (c Reset) Handle(state *state, args []string) (string, int, error) {
	if err := state.database.TruncateUsers(context.Background()); err != nil {
		return "", 1, fmt.Errorf("failed to truncate users: %w", err)
	}

	return "users truncated successfully", 0, nil
}
