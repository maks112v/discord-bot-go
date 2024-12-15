package commands

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type PingCommand struct {
	logger *zap.SugaredLogger
}

var _ Command = (*PingCommand)(nil)

func NewPingCommand(logger *zap.SugaredLogger) *PingCommand {
	return &PingCommand{
		logger: logger,
	}
}

func (c *PingCommand) Name() string {
	return "ping"
}

func (c *PingCommand) Description() string {
	return "Ping the bot"
}

func (c *PingCommand) Options() []*discordgo.ApplicationCommandOption {
	return nil
}

func (c *PingCommand) Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error {
	return nil
}

func (c *PingCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "pong",
		},
	})

	return err
}
