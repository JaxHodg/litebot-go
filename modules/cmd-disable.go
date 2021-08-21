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
			Name: "Disable",

			Function:    cmdDisable,
			Description: "Disables the specified command",
			HelpText:    "`{PREFIX}disable kick`",

			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdDisable(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a command")
	}

	moduleId := strings.ToLower(args[0])

	if !manager.IsValidModule(moduleId) && !manager.IsValidCommand(moduleId) {
		return functions.NewErrorEmbed(moduleId + " is not a valid Module")
	} else if !manager.IsValidVariable(moduleId, "enabled") {
		return functions.NewErrorEmbed(moduleId + " cannot be disabled")
	}
	module, err := manager.GetModule(moduleId)
	if err != nil {
		return functions.NewErrorEmbed("Unable to disable " + moduleId)
	}
	if isEnabled, _ := state.GetEnabled(event.Message.GuildID, moduleId); !isEnabled {
		return functions.NewGenericEmbed("Disabled", module.Name+" is already disabled")
	}

	state.DisableModule(event.Message.GuildID, moduleId)
	return functions.NewGenericEmbed("Disabled", "Disabled "+module.Name)
}
