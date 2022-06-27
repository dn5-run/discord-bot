package category

import (
	"dn5.run/discord/filter"
	"github.com/bwmarrin/discordgo"
)

func findSubCommand(data discordgo.ApplicationCommandInteractionData) *discordgo.ApplicationCommandInteractionDataOption {
	for _, option := range data.Options {
		if option.Type == discordgo.ApplicationCommandOptionSubCommand {
			return option
		}
	}
	return nil
}

func commandListener(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !filter.GuildFilter(i.GuildID) {
		return
	}

	if i.Type != discordgo.InteractionApplicationCommand {
		return
	}

	sub := findSubCommand(i.ApplicationCommandData())
	if sub == nil {
		return
	}

	switch sub.Name {
	case "create":
		createCategory(s, i, sub)
	default:
		s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
			Type: discordgo.InteractionResponseChannelMessageWithSource,
			Data: &discordgo.InteractionResponseData{
				Content: "Unknown subcommand",
			},
		})
	}
}
