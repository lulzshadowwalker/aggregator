package console

import (
	"fmt"

	"github.com/lulzshadowwalker/aggregator/internal/auth"
)

type Login struct {
	//
}

func (c Login) Name() string {
	return "login"
}

func (c Login) Description() string {
	return "login with username and password"
}

func (c Login) Handle(state *state, args []string) (string, error) {
	var username string
	if len(args) > 0 {
		username = args[0]
	}

	if err := auth.Login(auth.LoginParams{
		Username: username,
	}); err != nil {
		return "", err
	}

	return fmt.Sprintf("welcome back, %s!", username), nil
}
