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
		return functions.NewErrorEmbed("You must specify a command")
	}

	moduleId := strings.ToLower(args[0])

	if !manager.IsValidModule(moduleId) && !manager.IsValidCommand(moduleId) {
		return functions.NewErrorEmbed(moduleId + " is not a valid Module")
	} else if !manager.IsValidModule(moduleId) {
		return functions.NewErrorEmbed(moduleId + " cannot be Enabled")
	}
	module, err := manager.GetModule(moduleId)
	if err != nil {
		return functions.NewErrorEmbed("Unable to enable " + moduleId)
	}
	isEnabled, err := state.GetEnabled(event.Message.GuildID, moduleId)
	if err == nil && isEnabled {
		return functions.NewGenericEmbed("Enabled", module.Name+" is already Enabled")
	}

	state.EnableModule(event.Message.GuildID, moduleId)
	return functions.NewGenericEmbed("Enabled", "Enabled "+module.Name)
}
