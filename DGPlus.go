package DGPlus

import (
	"github.com/bwmarrin/discordgo"
	"strings"
)

func New(token string, CommandsCaseSensitive bool, RequirePrefix bool, CommandPrefix string) (s *DGPSession, e error) {
	client, err := discordgo.New("Bot " + token)
	s = &DGPSession{client, CommandsCaseSensitive, RequirePrefix, CommandPrefix, []*Command{}}
	e = err
	return
}

func messageHandler(dgop *DGPSession) func(*discordgo.Session, *discordgo.MessageCreate) {
	return func(s *discordgo.Session, msg *discordgo.MessageCreate) {
		args := strings.Split(msg.Content, " ")
		cmd := args[0]
		args = args[1:]

		if (dgop.RequirePrefix && strings.HasPrefix(cmd, dgop.CommandPrefix)) || !dgop.RequirePrefix {
			cmd = cmd[1:]

			for _, command := range dgop.Commands {
				if (dgop.CommandsCaseSensitive && command.Name == cmd) ||
					((!dgop.CommandsCaseSensitive) && strings.EqualFold(command.Name, cmd)) ||
					(command.IgnoreCaseSensitive && strings.EqualFold(command.Name, cmd)) {

					messageData := &MessageCreateData{
						DGPSession: dgop,
						DGSession:  s,
						Message:    msg,
						Command:    cmd,
						Args:       args,
						Content:    msg.Content,
						ChannelID:  msg.ChannelID,
						AuthorID:   msg.Author.ID,
						GuildID:    msg.GuildID,
					}

					command.HandlerFunc(messageData)
					return
				}

				for _, alias := range command.Aliases {
					if (dgop.CommandsCaseSensitive && alias == cmd) ||
						((!dgop.CommandsCaseSensitive) && strings.EqualFold(alias, cmd)) ||
						(command.IgnoreCaseSensitive && strings.EqualFold(alias, cmd)) {

						messageData := &MessageCreateData{
							DGPSession: dgop,
							DGSession:  s,
							Message:    msg,
							Command:    cmd,
							Args:       args,
							Content:    msg.Content,
							ChannelID:  msg.ChannelID,
							AuthorID:   msg.Author.ID,
							GuildID:    msg.GuildID,
						}

						command.HandlerFunc(messageData)
						return
					}
				}
			}
		}
	}
}
