package modules

import (
	"regexp"

	"../functions"
	"../manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Kick",
			Function:            cmdKick,
			Description:         "Kicks the mentioned user",
			RequiredPermissions: discordgo.PermissionKickMembers,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Kick",
			Description: "Kicks the mentioned user",
		},
	)
}

func cmdKick(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a user")
	}

	re := regexp.MustCompile(`[<][@](\d*)[>]`)
	userID := re.FindStringSubmatch(args[0])[1]

	if userID == "" {
		return functions.NewErrorEmbed("You must specify a user")
	}

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
