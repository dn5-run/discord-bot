package captcha

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

const CAPTCHA_BUTTON_ID = "captcha_button"
const CAPTCHA_AUTH_ID = "CAPTCHA_AUTH_ID"
const CAPTCHA_AUTH_CHANNEL_ID = "CAPTCHA_AUTH_CHANNEL_ID"

func checkEnv() {
	if os.Getenv(CAPTCHA_AUTH_ID) == "" {
		log.Fatalf("CAPTCHA | CAPTCHA_AUTH_ID not set")
	}
	if os.Getenv(CAPTCHA_AUTH_CHANNEL_ID) == "" {
		log.Fatalf("CAPTCHA | CAPTCHA_AUTH_CHANNEL_ID not set")
	}
}

func RegisterCaptcha(s *discordgo.Session) {
	checkEnv()
	s.AddHandler(readyListener)
	s.AddHandler(buttonListener)
	s.AddHandler(messageListener)
}
