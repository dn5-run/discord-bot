package category

import (
	"os"

	"dn5.run/discord/constants"
	"github.com/bwmarrin/discordgo"
)

type Category struct {
	ID          string
	Name        string
	Description string
	RoleID      string
	EntranceID  string
}

const ENTER_BUTTON_ID_PREFIX = "enter_button_"
const LEAVE_BUTTON_ID_PREFIX = "leave_button_"

func RegisterCategory(s *discordgo.Session) {
	commands := []*discordgo.ApplicationCommand{
		{
			Name:        "category",
			Description: "カテゴリの管理",
			Options: []*discordgo.ApplicationCommandOption{
				{
					Name:        "create",
					Description: "カテゴリを作成",
					Options: []*discordgo.ApplicationCommandOption{
						{
							Name:        "name",
							Description: "カテゴリ名",
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionString,
						},
						{
							Name:        "description",
							Description: "カテゴリの説明",
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionString,
						},
						{
							Name:        "role",
							Description: "カテゴリのロール",
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionRole,
						},
						{
							Name:        "entrance",
							Description: "カテゴリの入り口チャンネル",
							Required:    true,
							Type:        discordgo.ApplicationCommandOptionChannel,
						},
						{
							Name:        "image",
							Description: "カテゴリのイメージ",
							Type:        discordgo.ApplicationCommandOptionString,
						},
					},
					Type: discordgo.ApplicationCommandOptionSubCommand,
				},
			},
		},
	}

	for _, command := range commands {
		s.ApplicationCommandCreate(os.Getenv(constants.DISCORD_APP_ID), os.Getenv(constants.GUILD_ID), command)
	}

	s.AddHandler(commandListener)
	s.AddHandler(buttonListener)

	initDatabase()
}
