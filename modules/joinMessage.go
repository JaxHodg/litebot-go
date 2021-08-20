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
	manager.RegisterEnable("JoinMessage", false)
	manager.RegisterModule(
		&manager.Module{
			Name:        "JoinMessage",
			Description: "",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "JoinMessage",
			ModuleName:   "JoinMessage",
			DefaultValue: "Welcome {user}",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "JoinChannel",
			ModuleName:   "JoinMessage",
			DefaultValue: "",
		},
	)
}

// JoinMessage announces when a user joins the guild
func JoinMessage(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	// Checks if JoinMessage is enabled
	if isEnabled, _ := state.GetEnabled(event.GuildID, "JoinMessage"); !isEnabled {
		return
	}
	// Gets the the channelID for the server
	joinChannel, err := state.GetData(event.GuildID, "JoinMessage", "JoinChannel")
	if err != nil {
		return
	}
	channelID := functions.ExtractChannelID(joinChannel)

	// Checks if the channel works, probably not needed
	if _, err := session.Channel(channelID); err != nil {
		log.Println(err)
		return
	}

	// Gets the set message
	message, _ := state.GetData(event.GuildID, "JoinMessage", "JoinMessage")
	message = strings.ReplaceAll(message, "{user}", event.Mention())

	// Sends the message
	response := functions.NewGenericEmbed("New Member", message)
	if _, err = session.ChannelMessageSendEmbed(channelID, response); err != nil {
		log.Println(err)
	}
}
