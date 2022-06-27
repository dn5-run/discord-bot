package filter

import (
	"os"
)

func GuildFilter(ID string) bool {
	guild := os.Getenv("GUILD_ID")
	return ID == guild
}
