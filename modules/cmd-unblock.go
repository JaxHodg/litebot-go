package modules

import (
	"strings"

	"../functions"
	"../manager"
	"../state"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Unblock",
			Function:            cmdUnblock,
			Description:         "Unblocks the specified term",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdUnblock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	data := strings.Join(args, " ")
	if data == "" {
		return functions.NewErrorEmbed("You must specify a term to unblock")
	}
	pos := functions.Find(state.CheckList(event.Message.GuildID, "blocked"), data)
	if pos < 0 {
		return functions.NewErrorEmbed("`" + data + "` is not currently blocked")
	}

	state.RemoveFromList(event.Message.GuildID, "blocked", pos)
	return functions.NewGenericEmbed("Blocked", "Successfully unblocked `"+data+"`")
}
