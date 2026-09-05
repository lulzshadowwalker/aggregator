package console

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	"github.com/lulzshadowwalker/aggregator/internal/rss"
)

type AddFeed struct {
	//
}

func (c AddFeed) Name() string {
	return "addfeed"
}

func (c AddFeed) Description() string {
	return "subscribe to a certain feed"
}

func (c AddFeed) Handle(state *state, args []string) (string, int, error) {
	if len(args) != 2 {
		return "", 1, errors.New("usage: addfeed <name> <url>")
	}

	u, err := url.Parse(args[1])
	if err != nil {
		return "", 1, err
	}

	ctx := context.Background()

	tx, err := state.connection.BeginTx(ctx, nil)
	if err != nil {
		return "", 1, err
	}
	defer tx.Rollback()

	if err := rss.Subscribe(ctx, state.database.WithTx(tx), args[0], *u); err != nil {
		return "", 1, err
	}

	if err := tx.Commit(); err != nil {
		return "", 1, err
	}

	return fmt.Sprintf("subscribed to %s successfully", args[0]), 0, nil
}
