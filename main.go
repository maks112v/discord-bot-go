package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
	"github.com/maks112v/discord-bot-go/pkg/commands"
	"go.uber.org/zap"
)

func main() {
	zap, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("failed to create logger: %v", err)
		return
	}
	logger := zap.Sugar()

	// Load the .env file
	logger.Debug("Loading .env file")
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	logger.Info("Starting discord bot")
	token := os.Getenv("DISCORD_TOKEN")
	session, err := discordgo.New("Bot " + token)
	if err != nil {
		log.Fatalf("failed to create Discord session: %v", err)
	}

	if err := session.Open(); err != nil {
		logger.Fatalf("Failed to open Discord session: %v", err)
	}
	defer session.Close()

	logger.Info("Registering commands")
	cmd := commands.NewCommands(logger, session)
	cmd.RegisterCommands()
	cmd.RegisterHandlers()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	logger.Info("Bot is running, press Ctrl+C to exit")
	<-stop
}
