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
			Name: "JoinMessage",

			Function:    cmdJoinMessage,
			Description: "Configures messages when users join",
			HelpText:    "`{PREFIX}joinmessage #general`\nYou can also set a custom message that mentions the new user with {user}\n`{PREFIX}joinmessage #general Welcome {user}`",

			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdJoinMessage(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	channelID := ""
	message := ""

	if len(args) >= 1 {
		channelID = functions.ExtractChannelID(args[0])
	}
	if len(args) >= 2 {
		message = strings.Join(args[1:], " ")
		message = strings.TrimSpace(message)
	}
	if channelID == "" {
		return functions.NewErrorEmbed("The first argument must be the channel")
	}

	if _, err := session.Channel(channelID); err != nil {
		return functions.NewErrorEmbed("Invalid channel")
	} else if !functions.CanSpeak(session, channelID) {
		return functions.NewErrorEmbed("Unable to send messages in " + args[0])
	}

	state.SetData(event.GuildID, "JoinMessage", "JoinChannel", args[0])
	enabled, _ := state.GetEnabled(event.GuildID, "JoinMessage")

	if message == "" {
		return functions.NewModuleGenericEmbed("Set JoinMessage Channel", "Successfully set the JoinMessage channel to "+args[0]+"\nTip: You can also put a custom message after the channel", "JoinMessage", enabled)
	}
	state.SetData(event.GuildID, "JoinMessage", "JoinMessage", message)
	return functions.NewModuleGenericEmbed("Set JoinMessage Channel", "Successfully set the JoinMessage channel to "+args[0]+" and the message to:\n```"+message+"```", "JoinMessage", enabled)
}
