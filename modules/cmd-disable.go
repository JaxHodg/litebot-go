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
			Name:                "Disable",
			Function:            cmdDisable,
			Description:         "Disables the specified command",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdDisable(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a command")
	}

	module := strings.ToLower(args[0])

	if !manager.IsValidModule(module) && !manager.IsValidCommand(module) {
		return functions.NewErrorEmbed(module + " is not a valid Module")
	} else if !manager.IsValidModule(module) {
		return functions.NewErrorEmbed(module + " cannot be disabled")
	} else if !state.CheckEnabled(event.Message.GuildID, module) {
		return functions.NewGenericEmbed("Disabled", manager.GetModule(module).Name+" is already disabled")
	}

	state.DisableModule(event.Message.GuildID, module)
	return functions.NewGenericEmbed("Disabled", "Disabled "+manager.GetModule(module).Name)
}
