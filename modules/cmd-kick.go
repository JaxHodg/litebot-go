package modules

import (
	"regexp"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterEnable("Kick", true)
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Kick",
			Function:            cmdKick,
			Description:         "Kicks the user from the server",
			RequiredPermissions: discordgo.PermissionKickMembers,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Kick",
			Description: "Kicks the user from the server",
		},
	)
}

func cmdKick(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a user")
	}

	re := regexp.MustCompile(`[<][@](\d*)[>]`)
	substring := re.FindStringSubmatch(args[0])

	if len(substring) == 0 {
		return functions.NewErrorEmbed("You must specify a user")
	}

	userID := substring[1]

	user, err := session.GuildMember(event.Message.GuildID, userID)
	if err != nil {
		return functions.NewErrorEmbed("Invalid user")
	}

	err = session.GuildMemberDelete(event.Message.GuildID, userID)
	if err != nil {
		return functions.NewErrorEmbed("Unable to kick user")
	}

	return functions.NewGenericEmbed("Kick", "Kicked "+user.Mention())
}
