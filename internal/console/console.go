package console

import (
	"errors"

	"github.com/lulzshadowwalker/aggregator/internal/config"
	"github.com/lulzshadowwalker/aggregator/internal/database"
)

var (
	ErrConflict = errors.New("command with the same name already exists")
)

type Console struct {
	commands map[string]Commander
}

func New() *Console {
	commands := []Commander{
		Help{},
		Register{},
		Login{},
		Users{},
		Reset{},
	}

	console := &Console{
		commands: make(map[string]Commander),
	}

	for _, command := range commands {
		if err := console.Register(command); err != nil {
			panic(err)
		}
	}

	return console
}

func (c *Console) Run(args []string) (string, int, error) {
	state := &state{
		config:  config.Instance,
		console: *c,
		database: database.New(database.Connection),
	}

	var command Commander = Help{}
	if len(args) > 0 {
		if resolved, ok := c.commands[args[0]]; ok {
			command = resolved
		}

		// slicing beyond the length is only an issue if we exceed the capacity
		// args[1:] on a one-element slice is ok
		args = args[1:]

	}

	return command.Handle(state, args)
}

func (c *Console) Register(command Commander) error {
	if _, ok := c.commands[command.Name()]; ok {
		return ErrConflict
	}

	c.commands[command.Name()] = command
	return nil
}

// could be a struct too I guess but I would prefer an interface
type Commander interface {
	Name() string
	Description() string
	Handle(state *state, args []string) (string, int, error)
}

type state struct {
	config  *config.Config
	console Console
	database *database.Queries
}
