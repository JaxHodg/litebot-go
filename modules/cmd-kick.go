package modules

import (
	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterModule(
		&manager.Module{
			Name:        "Kick",
			Description: "Kicks the user from the server",
		},
	)
	manager.RegisterCommand(
		&manager.Command{
			Name:       "Kick",
			ModuleName: "Kick",

			Function:    cmdKick,
			Description: "Kicks the user from the server",
			HelpText:    "`{PREFIX}kick @user#1234`",

			RequiredPermissions: discordgo.PermissionKickMembers,
			GuildOnly:           true,
		},
	)
	manager.RegisterEnable("Kick", true)
}

func cmdKick(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	userID := ""
	if len(args) >= 1 {
		userID = functions.ExtractUserID(args[0])
	}

	if userID == "" {
		return functions.NewErrorEmbed("You must specify a user to kick")
	}

	if _, err := session.GuildMember(event.Message.GuildID, userID); err != nil {
		return functions.NewErrorEmbed("Invalid user")
	}

	if err := session.GuildMemberDelete(event.Message.GuildID, userID); err != nil {
		return functions.NewErrorEmbed("Unable to kick user")
	}

	return functions.NewGenericEmbed("Kick", "Kicked  "+args[0])
}
