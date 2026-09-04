package console

import (
	"context"
	"fmt"
	"strings"

	"github.com/lulzshadowwalker/aggregator/internal/config"
)

type Users struct {
	//
}

func (c Users) Name() string {
	return "users"
}

func (c Users) Description() string {
	return "list all users"
}

func (c Users) Handle(state *state, args []string) (string, int, error) {
	users, err := state.database.ListUsers(context.Background())
	if err != nil {
		return "", 1, err
	}

	var str strings.Builder
	str.WriteString("Users:\n")
	for _, user := range users {
		suffix := "\n"
		if user.Name == config.Instance.Username {
			suffix = " (current)\n"
		}

		fmt.Fprintf(&str, "\t* %s%s", user.Name, suffix)
	}

	return str.String(), 0, err
}
