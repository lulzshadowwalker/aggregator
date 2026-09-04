package console

import (
	"fmt"
	"strings"
)

type Help struct {
	//
}

func (c Help) Name() string {
	return "help"
}

func (c Help) Description() string {
	return "show this help message"
}

func (c Help) Handle(state *state, args []string) (string, error) {
	var str strings.Builder

	str.WriteString("welcome to the aggregator cli\n\n")
	str.WriteString("Usage:\n")
	str.WriteString("\taggregator <command>\n\n")
	str.WriteString("Available Commands:\n")

	for _, command := range state.console.commands {
		str.WriteString(fmt.Sprintf("\t%s: %s\n", command.Name(), command.Description()))
	}

	return str.String(), ErrHelp
}
