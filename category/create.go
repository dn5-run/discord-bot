package category

import (
	"log"

	"github.com/bwmarrin/discordgo"
	"github.com/google/uuid"
)

func createCategory(s *discordgo.Session, i *discordgo.InteractionCreate, c *discordgo.ApplicationCommandInteractionDataOption) {
	var name string
	var description string
	var role *discordgo.Role
	var entrance *discordgo.Channel
	var image interface{}

	for _, option := range c.Options {
		switch option.Name {
		case "name":
			name = option.StringValue()
		case "description":
			description = option.StringValue()
		case "role":
			role = option.RoleValue(s, i.GuildID)
		case "entrance":
			entrance = option.ChannelValue(s)
		case "image":
			image = option.StringValue()
		}
	}

	if name == "" || description == "" || role == nil || entrance == nil {
		return
	}

	id, err := uuid.NewUUID()
	if err != nil {
		log.Println(err)
		return
	}
	category := &Category{
		ID:          id.String(),
		Name:        name,
		Description: description,
		RoleID:      role.ID,
		EntranceID:  entrance.ID,
	}

	addCategory(category)

	embed := &discordgo.MessageEmbed{
		Title:       name,
		Description: description,
		Color:       0x00ff00,
	}

	if image != nil {
		embed.Image = &discordgo.MessageEmbedImage{
			URL: image.(string),
		}
	}

	_, err = s.ChannelMessageSendComplex(entrance.ID, &discordgo.MessageSend{
		Embeds: []*discordgo.MessageEmbed{
			embed,
		},
		Components: []discordgo.MessageComponent{
			discordgo.ActionsRow{
				Components: []discordgo.MessageComponent{
					discordgo.Button{
						Label:    name + "に入る",
						Style:    discordgo.PrimaryButton,
						CustomID: ENTER_BUTTON_ID_PREFIX + id.String(),
					},
					discordgo.Button{
						Label:    name + "から退出",
						Style:    discordgo.DangerButton,
						CustomID: LEAVE_BUTTON_ID_PREFIX + id.String(),
					},
				},
			},
		},
	})

	if err != nil {
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "カテゴリ作成中にエラーが発生しました。" + err.Error(),
			},
		})
		return
	}

	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Content: "カテゴリを作成しました。",
		},
	})
}
