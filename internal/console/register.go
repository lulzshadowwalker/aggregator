package console
import (
	"fmt"

	"github.com/lulzshadowwalker/aggregator/internal/auth"
)

type Register struct {
	//
}

func (c Register) Name() string {
	return "register"
}

func (c Register) Description() string {
	return "register with username"
}

func (c Register) Handle(state *state, args []string) (string, int, error) {
	var username string
	if len(args) > 0 {
		username = args[0]
	}

	if err := auth.Register(auth.Credentials{
		Username: username,
	}, state.database); err != nil {
		return "", 1, err
	}

	return fmt.Sprintf("welcome, %s!", username), 0, nil
}
