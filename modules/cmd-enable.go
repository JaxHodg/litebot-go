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
			Name:                "Enable",
			Function:            cmdEnable,
			Description:         "Enables the specified command",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdEnable(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a module")
	}

	module := strings.ToLower(args[0])

	if !manager.IsValidModule(module) && !manager.IsValidCommand(module) {
		return functions.NewErrorEmbed(module + " is not a valid module")
	} else if !manager.IsValidModule(module) {
		return functions.NewErrorEmbed(module + " cannot be enabled")
	} else if state.CheckEnabled(event.Message.GuildID, module) {
		return functions.NewGenericEmbed("Enabled", manager.GetModule(module).Name+" is already enabled")
	}

	state.EnableModule(event.Message.GuildID, module)

	return functions.NewGenericEmbed("Enabled", "Enabled "+manager.GetModule(module).Name)
}
