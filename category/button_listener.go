package category

import (
	"log"
	"strings"

	"dn5.run/discord/filter"
	"github.com/bwmarrin/discordgo"
)

func getID(s string) string {
	l := strings.Split(s, "_")
	return l[len(l)-1]
}

func getRole(s *discordgo.Session, guildId string, roleId string) *discordgo.Role {
	roles, err := s.GuildRoles(guildId)
	if err != nil {
		return nil
	}
	for _, role := range roles {
		if role.ID == roleId {
			return role
		}
	}
	return nil
}

func enterCategory(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := getID(i.MessageComponentData().CustomID)

	c := getCategory(id)
	if c == nil {
		return
	}

	role := getRole(s, i.GuildID, c.RoleID)

	err := s.GuildMemberRoleAdd(i.GuildID, i.Member.User.ID, role.ID)
	if err != nil {
		log.Printf("Category | Error adding role: %s", err.Error())
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
			Content: "カテゴリに参加しました🍡🍡",
		},
	})
}

func leaveCategory(s *discordgo.Session, i *discordgo.InteractionCreate) {
	id := getID(i.MessageComponentData().CustomID)

	c := getCategory(id)
	if c == nil {
		return
	}

	role := getRole(s, i.GuildID, c.RoleID)

	err := s.GuildMemberRoleRemove(i.GuildID, i.Member.User.ID, role.ID)
	if err != nil {
		log.Printf("Category | Error removing role: %s", err.Error())
		return
	}
	s.InteractionRespond(i.Interaction, &discordgo.InteractionResponse{
		Type: discordgo.InteractionResponseChannelMessageWithSource,
		Data: &discordgo.InteractionResponseData{
			Flags:   uint64(discordgo.MessageFlagsEphemeral),
			Content: "カテゴリから退出しました。🐼🍞",
		},
	})
}

func buttonListener(s *discordgo.Session, i *discordgo.InteractionCreate) {
	if !filter.GuildFilter(i.GuildID) {
		return
	}

	if i.Type != discordgo.InteractionMessageComponent {
		return
	}

	if strings.HasPrefix(i.MessageComponentData().CustomID, ENTER_BUTTON_ID_PREFIX) {
		enterCategory(s, i)
	}
	if strings.HasPrefix(i.MessageComponentData().CustomID, LEAVE_BUTTON_ID_PREFIX) {
		leaveCategory(s, i)
	}
}
