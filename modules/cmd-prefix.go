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
			Name:                "Prefix",
			Function:            cmdPrefix,
			Description:         "Configures the prefix",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Prefix",
			ModuleName:   "Prefix",
			DefaultValue: "!",
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name: "Prefix",

			Description: "",
		},
	)
}

func cmdPrefix(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	prefix := ""
	if len(args) >= 1 {
		prefix = args[0]
		prefix = strings.TrimSpace(prefix)
	}
	if prefix == "" {
		return functions.NewErrorEmbed("You must specify a new prefix")
	}
	state.SetData(event.GuildID, "Prefix", "Prefix", prefix)
	return functions.NewGenericEmbed("Set Prefix", "Successfully set prefix to `"+prefix+"`")
}
