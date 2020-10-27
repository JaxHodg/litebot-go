package modules

import (
	"log"
	"regexp"
	"strings"

	"../functions"
	"../manager"
	"../state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterModule(
		&manager.Module{
			Name:        "LeaveMessage",
			Description: "",
		},
	)
}

// LeaveMessage announces when a user leaves the guild
func LeaveMessage(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	if !state.CheckEnabled(event.GuildID, "leavemessage") {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)

	submatch := re.FindStringSubmatch(state.CheckData(event.GuildID, "leavechannel"))
	if len(submatch) == 0 {
		return
	}
	channelID := submatch[1]

	_, err := session.Channel(channelID)
	if err != nil {
		log.Println(err)
		return
	}

	message := state.CheckData(event.GuildID, "leavemessage")
	if message == "" {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := functions.NewGenericEmbed("Member left", message)

	if response != nil {
		_, err := session.ChannelMessageSendEmbed(channelID, response)
		if err != nil {
			log.Println(err)
		}
	}
}
