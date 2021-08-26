package modules

import (
	"log"
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterModule(
		&manager.Module{
			Name:        "LeaveMessage",
			Description: "",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "LeaveMessage",
			ModuleName:   "LeaveMessage",
			DefaultValue: "Goodbye {user}",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "LeaveChannel",
			ModuleName:   "LeaveMessage",
			DefaultValue: "",
		},
	)
	manager.RegisterEnable("LeaveMessage", false)
}

// LeaveMessage announces when a user leaves the guild
func LeaveMessage(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	// Checks if leaveMessage is enabled
	if isEnabled, _ := state.GetEnabled(event.GuildID, "leaveMessage"); !isEnabled {
		return
	}
	// Gets the the channelID for the server
	leaveChannel, err := state.GetData(event.GuildID, "leaveMessage", "leaveChannel")
	if err != nil {
		return
	}
	channelID := functions.ExtractChannelID(leaveChannel)

	// Checks if the channel works, probably not needed
	if _, err := session.Channel(channelID); err != nil {
		log.Println(err)
		return
	}

	// Gets the set message
	message, _ := state.GetData(event.GuildID, "leaveMessage", "leaveMessage")
	message = strings.ReplaceAll(message, "{user}", event.Mention())

	// Sends the message
	response := functions.NewGenericEmbed("Member Left", message)
	if _, err = session.ChannelMessageSendEmbed(channelID, response); err != nil {
		log.Println(err)
	}
}
