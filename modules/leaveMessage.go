package modules

import (
	"log"
	"regexp"
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterEnable("LeaveMessage", false)
	manager.RegisterModule(
		&manager.Module{
			Name:        "LeaveMessage",
			Description: "",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Message",
			ModuleName:   "LeaveMessage",
			DefaultValue: "Goodbye {user}",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Channel",
			ModuleName:   "LeaveMessage",
			DefaultValue: "",
		},
	)
}

// LeaveMessage announces when a user leaves the guild
func LeaveMessage(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	isEnabled, err := state.GetEnabled(event.GuildID, "LeaveMessage")

	if err != nil || !isEnabled {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)

	leaveChannel, err := state.GetData(event.GuildID, "LeaveMessage", "leavechannel")
	if err != nil {
		return
	}
	submatch := re.FindStringSubmatch(leaveChannel)
	if len(submatch) == 0 {
		return
	}
	channelID := submatch[1]

	_, err = session.Channel(channelID)
	if err != nil {
		log.Println(err)
		return
	}

	message, err := state.GetData(event.GuildID, "LeaveMessage", "LeaveMessage")
	if err != nil {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := functions.NewGenericEmbed("Member Left", message)

	if response != nil {
		_, err := session.ChannelMessageSendEmbed(channelID, response)
		if err != nil {
			log.Println(err)
		}
	}
}
