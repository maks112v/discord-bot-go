package commands

import (
	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type SubCommandExampleCommand struct {
	logger *zap.SugaredLogger
}

var _ Command = (*SubCommandExampleCommand)(nil)

func NewSubCommandExampleCommand(logger *zap.SugaredLogger) *SubCommandExampleCommand {
	return &SubCommandExampleCommand{
		logger: logger,
	}
}

func (c *SubCommandExampleCommand) Name() string {
	return "subcommand"
}

func (c *SubCommandExampleCommand) Description() string {
	return "Example subcommand"
}

func (c *SubCommandExampleCommand) Options() []*discordgo.ApplicationCommandOption {
	return []*discordgo.ApplicationCommandOption{
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "list",
			Description: "List all subcommands",
		},
		{
			Type:        discordgo.ApplicationCommandOptionSubCommand,
			Name:        "create",
			Description: "Create a new subcommand",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Type:        discordgo.ApplicationCommandOptionString,
					Name:        "name",
					Description: "Name of the subcommand",
					Required:    true,
				},
			},
		},
	}
}

func (c *SubCommandExampleCommand) Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error {
	return nil
}

func (c *SubCommandExampleCommand) Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error {
	subCommand := i.ApplicationCommandData().Options[0]

	var content string
	switch subCommand.Name {
	case "list":
		content = "List of subcommands"
	case "create":
		name := subCommand.Options[0].StringValue()
		content = "Creating subcommand: " + name
	}

	err := s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: content,
		},
	})
	if err != nil {
		return err
	}

	return err
}
