package captcha

import (
	"log"
	"math/rand"

	"dn5.run/discord/filter"
	"github.com/bwmarrin/discordgo"
)

func randomString(digit uint32) string {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, digit)
	if _, err := rand.Read(b); err != nil {
		return ""
	}
	var result string
	for _, v := range b {
		result += string(letters[int(v)%len(letters)])
	}
	return result

}

func buttonListener(s *discordgo.Session, e *discordgo.InteractionCreate) {
	if !filter.GuildFilter(e.GuildID) {
		return
	}

	if e.Type != discordgo.InteractionMessageComponent || e.MessageComponentData().CustomID != CAPTCHA_BUTTON_ID {
		return
	}

	key := randomString(6)
	image := createCaptchaImage(key)

	log.Printf("Captcha key: %s", key)

	err := s.InteractionRespond(e.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "画像の文字列を送信してください。",
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
			Files: []*discordgo.File{
				{
					Name:        "captcha.png",
					ContentType: "image/png",
					Reader:      image,
				},
			},
		},
	})

	if err != nil {
		log.Printf("CAPTCHA | Error sending message: %s", err.Error())
		return
	}

	registerKey(e.Member.User.ID, key)
	log.Printf("CAPTCHA | Sent captcha to %s", e.Member.User.Username)
}
