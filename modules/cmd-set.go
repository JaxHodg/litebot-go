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
			Name:                "Set",
			Function:            cmdSet,
			Description:         "Sets the specified value",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdSet(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewGenericEmbed("Set values", "You may set the following values: "+strings.Join(state.DataValues, ", "))
	} else if len(args) < 2 {
		return functions.NewErrorEmbed("You must specify what to set " + args[0] + " to")
	}

	value := strings.ToLower(args[0])

	if !functions.Contains(state.DataValues, value) {
		return functions.NewErrorEmbed(args[0] + " is an invalid value")
	}

	data := strings.Join(args[1:], " ")

	state.SetData(event.Message.GuildID, args[0], data)

	return functions.NewGenericEmbed("Set", "Successfully set `"+args[0]+"` to `"+data+"`")
}
