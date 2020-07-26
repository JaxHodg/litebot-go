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
			Name:                "Block",
			Function:            cmdBlock,
			Description:         "Blocks the specified term",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdBlock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	data := strings.ToLower(strings.Join(args, " "))
	if data == "" {
		return functions.NewErrorEmbed("You must specify a term to block")
	}
	if functions.Find(state.CheckList(event.Message.GuildID, "blocked"), data) != -1 {
		return functions.NewErrorEmbed("`" + data + "` is already blocked")
	}

	state.AddToList(event.Message.GuildID, "blocked", data)
	return functions.NewGenericEmbed("Blocked", "Successfully blocked `"+data+"`")
}
