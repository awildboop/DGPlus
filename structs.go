package dgplus

import "github.com/bwmarrin/discordgo"

type Command struct {
	Name                string
	Aliases             []string
	IgnoreCaseSensitive bool
	IgnoreRequirePrefix bool
	HandlerFunc         func(*MessageCreateData)
}

type MessageCreateData struct {
	DGPSession *DGPSession
	DGSession  *discordgo.Session
	Message    *discordgo.MessageCreate
	Command    string
	Args       []string
	Content    string
	ChannelID  string
	AuthorID   string
	GuildID    string
}

func (msg *MessageCreateData) Send(message string) (*discordgo.Message, error) {
	return msg.DGSession.ChannelMessageSend(msg.ChannelID, message)
}

func (msg *MessageCreateData) SendEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	return msg.DGSession.ChannelMessageSendEmbed(msg.ChannelID, embed)
}

func (msg *MessageCreateData) SendComplex(content *discordgo.MessageSend) (*discordgo.Message, error) {
	return msg.DGSession.ChannelMessageSendComplex(msg.ChannelID, content)
}

func (msg *MessageCreateData) SendReply(content string) (*discordgo.Message, error) {
	return msg.DGSession.ChannelMessageSendReply(msg.ChannelID, content, msg.Message.Reference())
}
