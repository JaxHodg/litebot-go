package modules

import (
	"regexp"

	"github.com/bwmarrin/discordgo"

	"../functions"
	"../manager"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Ban",
			Function:            cmdBan,
			Description:         "Bans the mentioned user",
			RequiredPermissions: discordgo.PermissionBanMembers,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Ban",
			Description: "Bans the mentioned user",
		},
	)
}

func cmdBan(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
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

	err = session.GuildBanCreate(event.Message.GuildID, userID, 0)
	if err != nil {
		return functions.NewErrorEmbed("Unable to ban user")
	}

	return functions.NewGenericEmbed("Ban", "Banned  "+user.Mention())
}
