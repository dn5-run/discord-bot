package captcha

import (
	"log"
	"os"

	"dn5.run/discord/filter"
	"github.com/bwmarrin/discordgo"
)

func messageListener(s *discordgo.Session, e *discordgo.MessageCreate) {
	if !filter.GuildFilter(e.GuildID) {
		return
	}

	key := getKey(e.Author.ID)
	if key == "" {
		return
	}

	if key != e.Content {
		s.ChannelMessageDelete(e.ChannelID, e.ID)
		return
	}

	s.GuildMemberRoleAdd(e.GuildID, e.Author.ID, os.Getenv(CAPTCHA_AUTH_ID))
	s.ChannelMessageDelete(e.ChannelID, e.ID)
	unregisterKey(e.Author.ID)

	log.Printf("CAPTCHA | %s passed captcha", e.Author.Username)
}
