package commands

import (
	"time"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type FollowUpCommand struct {
	logger *zap.SugaredLogger
}

var _ Command = (*FollowUpCommand)(nil)

func NewFollowUpCommand(logger *zap.SugaredLogger) *FollowUpCommand {
	return &FollowUpCommand{
		logger: logger,
	}
}

func (c *FollowUpCommand) Name() string {
	return "followup"
}

func (c *FollowUpCommand) Description() string {
	return "Command that demonstrates a follow-up message"
}

func (c *FollowUpCommand) Options() []*discordgo.ApplicationCommandOption {
	return nil
}

func (c *FollowUpCommand) Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error {
	return nil
}

func (c *FollowUpCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "Starting a command that takes a while to complete",
		},
	})

	time.Sleep(5 * time.Second)

	s.FollowupMessageCreate(i.Interaction, true, &discordgo.WebhookParams{
		Content: "Responding in a thread to follow up with the status",
	})

	return nil
}
