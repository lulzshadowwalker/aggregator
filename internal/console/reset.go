package console

import (
	"context"
	"fmt"

	"github.com/lulzshadowwalker/aggregator/internal/config"
)

type Reset struct {
	//
}

func (c Reset) Name() string {
	return "reset"
}

func (c Reset) Description() string {
	return "delete all users"
}

func (c Reset) Handle(state *state, args []string) (string, int, error) {
	if err := state.database.DeleteUsers(context.Background()); err != nil {
		return "", 1, fmt.Errorf("failed to delete all users: %w", err)
	}

	config.Instance.Username = ""
	if err := config.Instance.Write(); err != nil {
		return "", 1, err
	}

	return "all users deleted successfully", 0, nil
}
