package modules

import (
	"github.com/bwmarrin/discordgo"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
)

func init() {
	manager.RegisterEnable("Ban", true)
	manager.RegisterCommand(
		&manager.Command{
			Name:       "Ban",
			ModuleName: "Ban",

			Function:    cmdBan,
			Description: "Bans the user from the server",
			HelpText:    "`{PREFIX}ban @user#1234`",

			RequiredPermissions: discordgo.PermissionBanMembers,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Ban",
			Description: "Bans the user from the server",
		},
	)
}

func cmdBan(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	userID := ""
	if len(args) >= 1 {
		userID = functions.ExtractUserID(args[0])
	}

	if userID == "" {
		return functions.NewErrorEmbed("You must specify a user to ban")
	}

	if _, err := session.GuildMember(event.Message.GuildID, userID); err != nil {
		return functions.NewErrorEmbed("Invalid user")
	}

	if err := session.GuildBanCreate(event.Message.GuildID, userID, 0); err != nil {
		return functions.NewErrorEmbed("Unable to ban user")
	}

	return functions.NewGenericEmbed("Ban", "Banned  "+args[0])
}
