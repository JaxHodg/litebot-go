package modules

import (
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
			Name:        "JoinMessage",
			Description: "",
		},
	)
}

func JoinMessage(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	if !state.CheckEnabled(event.GuildID, "JoinMessage") {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)

	submatch := re.FindStringSubmatch(state.CheckData(event.GuildID, "joinchannel"))
	if len(submatch) == 0 {
		return
	}
	channelID := submatch[1]

	_, err := session.Channel(channelID)
	if err != nil {
		return
	}

	message := state.CheckData(event.GuildID, "joinmessage")
	if message == "" {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := functions.NewGenericEmbed("New Member", message)

	if response != nil {
		session.ChannelMessageSendEmbed(channelID, response)
	}
}
