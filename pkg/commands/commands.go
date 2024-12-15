package commands

import (
	"fmt"

	"github.com/bwmarrin/discordgo"
	"go.uber.org/zap"
)

type Command interface {
	Name() string
	Description() string
	Options() []*discordgo.ApplicationCommandOption
	Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error
	Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error
}

type Commands struct {
	logger  *zap.SugaredLogger
	session *discordgo.Session

	activeCommands map[string]Command
}

func NewCommands(logger *zap.SugaredLogger, session *discordgo.Session) *Commands {
	// Create instances of the active commands
	ping := NewPingCommand(logger)
	validation := NewValidationExampleCommand(logger)
	subCommand := NewSubCommandExampleCommand(logger)

	// Create a map of active commands
	activeCommands := map[string]Command{
		ping.Name():       ping,
		validation.Name(): validation,
		subCommand.Name(): subCommand,
	}

	return &Commands{
		logger:         logger,
		session:        session,
		activeCommands: activeCommands,
	}
}

// Register commands with Discord
func (c *Commands) RegisterCommands() {
	for _, cmd := range c.activeCommands {
		_, err := c.session.ApplicationCommandCreate(c.session.State.User.ID, "", &discordgo.ApplicationCommand{
			Name:        cmd.Name(),
			Description: cmd.Description(),
			Options:     cmd.Options(),
		})
		if err != nil {
			c.logger.Fatalf("Failed to create command %s: %v", cmd.Name(), err)
		}
		c.logger.Infow("Registered command", "name", cmd.Name())
	}
}

// RegisterHandlers registers the command handlers with the Discord session
func (c *Commands) RegisterHandlers() {
	c.session.AddHandler(func(s *discordgo.Session, i *discordgo.InteractionCreate) {
		// Get the username and user ID from the interaction for validation and logging
		var username string
		var userId string
		if i.User != nil {
			username = i.User.String()
			userId = i.User.ID
		} else if i.Member != nil && i.Member.User != nil {
			username = i.Member.User.String()
			userId = i.Member.User.ID
		}

		// Check if the interaction is an application command
		if i.Type == discordgo.InteractionApplicationCommand {
			c.logger.Infow("Received interaction command",
				"command", i.ApplicationCommandData().Name,
				"user", username,
				"userId", userId,
			)

			// Get the command from the active commands map
			cmd, exists := (c.activeCommands)[i.ApplicationCommandData().Name]

			// If the command doesn't exist, log an error and respond to the interaction
			if !exists {
				c.logger.Errorf("Unknown command: %s", i.ApplicationCommandData().Name)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Unknown command: %s", i.ApplicationCommandData().Name),
					},
				})
				return
			}

			// Validate the command
			if err := cmd.Validate(s, i, username); err != nil {
				c.logger.Errorf("Error validating command %s: %v", cmd.Name(), err)
				s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
					Type: discordgo.InteractionResponseChannelMessageWithSource,
					Data: &discordgo.InteractionResponseData{
						Content: fmt.Sprintf("Validation failed: %s", err.Error()),
					},
				})
				return
			}

			// Execute the command
			if err := cmd.Execute(s, i); err != nil {
				c.logger.Errorf("Error executing command %s: %v", cmd.Name(), err)
				return
			}

			c.logger.Infof("Command %s executed successfully", cmd.Name())
		}
	})

}
