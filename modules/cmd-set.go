package modules

import (
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"

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

	if value == "joinmessage" || value == "leavemessage" {
		return functions.NewGenericEmbed("Set", "Successfully set `"+args[0]+"` to `"+data+"`"+"\nTip: if your message contains `{user}` I will replace it by mentioning the user")
	}

	return functions.NewGenericEmbed("Set", "Successfully set `"+args[0]+"` to `"+data+"`")
}
