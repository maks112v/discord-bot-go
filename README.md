# Discord Bot Template (Go)

A template for creating Discord bots using Go, featuring a clean command structure, logging, and environment configuration.

## Trying it out

To try out the example bot, you can install it using the following link:

[Install Example Bot](https://discord.com/oauth2/authorize?client_id=1317678697834221640)

## Features

- Slash command support with validation
- Subcommand handling
- Structured logging with Zap
- Environment configuration with .env
- Clean, modular command architecture
- Example commands included

## Prerequisites

- Go 1.16 or higher
- A Discord bot token
- Discord application with slash commands enabled

## Setup

1. Clone the repository:
```bash
git clone https://github.com/maks112v/discord-bot-go
cd discord-bot-go
```

2. Install dependencies:
```bash
go mod download
```

3. Copy the example environment file:
```bash
cp .env.example .env
```

4. Add your Discord bot token to `.env`:
```bash
DISCORD_TOKEN="your-bot-token-here"
```

5. Run the bot:
```bash
go run main.go
```

## Debugging

This template includes VS Code debugging configuration. To debug the bot:

1. Open the project in VS Code
2. Set any breakpoints in your code
3. Press F5 or select "Run and Debug" from the sidebar
4. Select "Launch Bot" from the debug configuration dropdown

The bot will start in debug mode, allowing you to:
- Step through code execution
- Inspect variables
- Set breakpoints
- Use debug console

## Project Structure

```
discord-bot-go/
├── main.go                 # Entry point
├── .env                    # Environment configuration
└── pkg/
    └── commands/          # Command implementations
        ├── commands.go    # Command handler and registration
        ├── ping.go        # Simple ping command
        ├── subcommands.go # Command with subcommands
        ├── validation.go  # Command with validation
        └── followup.go    # Command with follow-up messages
```

## Adding New Commands

1. Create a new command file in `pkg/commands/`
2. Implement the `Command` interface:
```go
type Command interface {
    Name() string
    Description() string
    Options() []*discordgo.ApplicationCommandOption
    Validate(s *discordgo.Session, i *discordgo.InteractionCreate, username string) error
    Execute(s *discordgo.Session, i *discordgo.InteractionCreate) error
}
```
3. Register your command in `commands.go`

## Example Commands

- `/ping` - Basic ping command
- `/validation` - Example of command with guild-only validation
- `/followup` - Example of using follow-up messages in threads
- `/subcommand` - Example of command with subcommands
  - `/subcommand list`
  - `/subcommand create`

## License

Unlicense license. See [LICENSE](LICENSE) for more information.

## Contributing

Pull requests are welcome. For major changes, please open an issue first to discuss what you would like to change.