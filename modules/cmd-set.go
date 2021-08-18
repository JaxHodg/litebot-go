package modules

import (
	"regexp"
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
			Description:         "Sets a variable for a module",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdSet(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewGenericEmbed("Set", "To set a variable, you must first specify one of the following modules: `"+strings.Join(manager.ListModules(), "`, `")+"`")
	} else if len(args) == 1 {
		if manager.IsValidModule(args[0]) {
			variablesList := manager.ListVariables(args[0])
			if len(variablesList) != 0 {
				return functions.NewGenericEmbed("Set `"+args[0]+"`", "To set a variable, you must specify one of the following variables: `"+strings.Join(manager.ListVariables(args[0]), "`, `")+"`")
			} else {
				return functions.NewGenericEmbed("Set `"+args[0]+"`", "No variables to set")
			}
		} else {
			return functions.NewErrorEmbed("Unknown module")
		}
	} else if len(args) == 2 {
		if manager.IsValidVariable(args[0], args[0]) {
			moduleID := strings.ToLower(args[0])
			variableID := strings.ToLower(args[0])
			variable := strings.Join(args[1:], " ")

			state.SetData(event.Message.GuildID, moduleID, variableID, variable)

			return functions.NewGenericEmbed("Set", "Successfully set `"+variableID+"` to `"+variable+"`")
		}
		return functions.NewErrorEmbed("You must specify a value for `" + args[1] + "`")
	}

	moduleID := strings.ToLower(args[0])
	variableID := strings.ToLower(args[1])
	variable := strings.Join(args[2:], " ")

	if !manager.IsValidVariable(moduleID, variableID) || variableID == "enabled" {
		return functions.NewErrorEmbed(variableID + " is not a valid variable")
	}

	state.SetData(event.Message.GuildID, moduleID, variableID, variable)

	/**if value == "joinmessage" || value == "leavemessage" {
		return functions.NewGenericEmbed("Set", "Successfully set `"+args[0]+"` to `"+data+"`"+"\nTip: if your message contains `{user}` I will replace it by mentioning the user")
	}**/

	// Checks if channel
	re := regexp.MustCompile(`<#(\d*)>`)
	submatch := re.FindStringSubmatch(variable)
	if len(submatch) != 0 {
		return functions.NewGenericEmbed("Set", "Successfully set `"+variableID+"` to "+variable)
	}

	return functions.NewGenericEmbed("Set", "Successfully set `"+variableID+"` to `"+variable+"`")
}
