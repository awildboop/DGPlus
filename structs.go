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
	newMessage, err := msg.DGSession.ChannelMessageSend(msg.ChannelID, message)
	return newMessage, err
}

func (msg *MessageCreateData) SendEmbed(embed *discordgo.MessageEmbed) (*discordgo.Message, error) {
	newMessage, err := msg.DGSession.ChannelMessageSendEmbed(msg.ChannelID, embed)
	return newMessage, err
}

func (msg *MessageCreateData) SendComplex(content *discordgo.MessageSend) (*discordgo.Message, error) {
	newMessage, err := msg.DGSession.ChannelMessageSendComplex(msg.ChannelID, content)
	return newMessage, err
}

func (msg *MessageCreateData) SendReply(content string) (*discordgo.Message, error) {
	newMessage, err := msg.DGSession.ChannelMessageSendReply(msg.ChannelID, content, msg.Message.Reference())
	return newMessage, err
}
