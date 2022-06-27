package main

import (
	"log"
	"os"
	"os/signal"

	"dn5.run/discord/captcha"
	"dn5.run/discord/category"
	"dn5.run/discord/constants"
	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

const DISCORD_TOKEN = "DISCORD_TOKEN"
const DISCORD_APP_ID = "DISCORD_APP_ID"
const GUILD_ID = "GUILD_ID"

func env() {
	if _, err := os.Stat(".env"); err != nil {
		return
	}
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	if os.Getenv(constants.DISCORD_TOKEN) == "" {
		log.Fatalf("DISCORD_TOKEN not set")
	}
	if os.Getenv(constants.DISCORD_APP_ID) == "" {
		log.Fatalf("DISCORD_APP_ID not set")
	}
	if os.Getenv(constants.GUILD_ID) == "" {
		log.Fatalf("GUILD_ID not set")
	}
}

func register(discord *discordgo.Session) {
	captcha.RegisterCaptcha(discord)
	category.RegisterCategory(discord)
}

func main() {
	env()

	s, err := discordgo.New("Bot " + os.Getenv("DISCORD_TOKEN"))
	if err != nil {
		log.Fatalf("Error creating Discord session: %s", err.Error())
	}
	defer s.Close()

	s.Identify.Intents = discordgo.IntentGuildMembers | discordgo.IntentGuildMessages
	s.AddHandlerOnce(func(s *discordgo.Session, e *discordgo.Ready) {
		log.Printf("Ready: %s", e.User.Username)
	})

	register(s)

	if err := s.Open(); err != nil {
		log.Fatalf("Error opening Discord session: %s", err.Error())
	}

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)
	<-stop
	log.Println("Stopping...")
}
