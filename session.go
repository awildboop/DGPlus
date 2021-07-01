package dgplus

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"strings"
)

type DGPSession struct {
	DGSession             *discordgo.Session
	CommandsCaseSensitive bool
	RequirePrefix         bool   // Whether a message should be checked for CommandPrefix
	CommandPrefix         string // ignored if RequirePrefix false
	Commands              []*Command
}

func (s *DGPSession) RegisterCommand(name string, aliases []string, ignoreCaseSensitive bool, ignoreRequirePrefix bool, handler func(*MessageCreateData)) (cmd *Command, e error) {
	cmd = &Command{name, aliases, ignoreCaseSensitive, ignoreRequirePrefix, handler}

	// check if existing commands contain new commands name or alias
	for _, c := range s.Commands {
		if strings.EqualFold(c.Name, name) {
			em := fmt.Sprintf("command name (%s) already exists", name)
			e = fmt.Errorf(em)
			return
		}

		for _, alias := range c.Aliases {
			if strings.EqualFold(alias, name) {
				em := fmt.Sprintf("command name (%s) already exists as alias", name)
				e = fmt.Errorf(em)
				return
			}

			for _, alias2 := range aliases {
				if strings.EqualFold(alias, alias2) {
					em := fmt.Sprintf("command alias (%s) already exists as alias", alias2)
					e = fmt.Errorf(em)
				}
			}
		}

	}

	// check if new commands aliases contain the commands name
	for _, alias := range aliases {
		if strings.EqualFold(alias, name) {
			e = fmt.Errorf("command aliases should not contain command name")
			return
		}
	}

	s.Commands = append(s.Commands, cmd)
	return cmd, nil
}

// Open sets intents to IntentsAllWithoutPrivileged, adds the command message handler, and opens the session
func (s *DGPSession) Open() error {
	s.DGSession.Identify.Intents = discordgo.IntentsAllWithoutPrivileged
	s.DGSession.AddHandler(messageHandler(s))

	return s.DGSession.Open()
}

func (s *DGPSession) Close() error {
	return s.DGSession.Close()
}
