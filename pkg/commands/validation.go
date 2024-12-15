package commands

import (
	"errors"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type ValidationExampleCommand struct {
	logger *zap.SugaredLogger
}

var _ Command = (*ValidationExampleCommand)(nil)

func NewValidationExampleCommand(logger *zap.SugaredLogger) *ValidationExampleCommand {
	return &ValidationExampleCommand{
		logger: logger,
	}
}

func (c *ValidationExampleCommand) Name() string {
	return "validation"
}

func (c *ValidationExampleCommand) Description() string {
	return "Example command with validation"
}

func (c *ValidationExampleCommand) Options() []*discordgo.ApplicationCommandOption {
	return nil
}

func (c *ValidationExampleCommand) Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error {
	// Only allow this command to be used in a server
	if i.GuildID == "" {
		return errors.New("command can only be used in a server")
	}

	return nil
}

func (c *ValidationExampleCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "successfully ran in a server",
		},
	})

	return err
}
