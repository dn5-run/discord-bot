package captcha

import (
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

func readyListener(s *discordgo.Session, e *discordgo.Ready) {
	messages, err := s.ChannelMessages(os.Getenv(CAPTCHA_AUTH_CHANNEL_ID), 10, "", "", "")
	if err != nil {
		log.Printf("CAPTCHA | Error getting messages: %s", err.Error())
		return
	}

	for _, message := range messages {
		if message.Author.ID == s.State.User.ID {
			return
		}
	}

	_, err = s.ChannelMessageSendComplex(os.Getenv(CAPTCHA_AUTH_CHANNEL_ID), &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			{
				Title: "Captcha",
				Description: "下のボタンを押して認証してください。\n" +
					"上手く動かない場合は <@485065877750939649> にDMを送信してください。",
				Color: 0x00ff00,
			},
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    "Captcha",
						Style:    discordgo.PrimaryButton,
						CustomID: CAPTCHA_BUTTON_ID,
					},
				},
			},
		},
	})

	if err != nil {
		log.Printf("CAPTCHA | Error sending message: %s", err.Error())
		return
	}
}
