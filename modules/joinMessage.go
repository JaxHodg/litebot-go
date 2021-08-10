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
	manager.RegisterEnable("JoinMessage", false)
	manager.RegisterModule(
		&manager.Module{
			Name:        "JoinMessage",
			Description: "",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Message",
			ModuleName:   "JoinMessage",
			DefaultValue: "Welcome {user}",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Channel",
			ModuleName:   "JoinMessage",
			DefaultValue: "",
		},
	)
}

// JoinMessage announces when a user joins the guild
func JoinMessage(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	isEnabled, err := state.GetEnabled(event.GuildID, "JoinMessage")

	if err != nil || !isEnabled {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)

	joinChannel, err := state.GetData(event.GuildID, "JoinMessage", "joinchannel")
	if err != nil {
		return
	}
	submatch := re.FindStringSubmatch(joinChannel)
	if len(submatch) == 0 {
		return
	}
	channelID := submatch[1]

	_, err = session.Channel(channelID)
	if err != nil {
		log.Println(err)
		return
	}

	message, err := state.GetData(event.GuildID, "JoinMessage", "JoinMessage")
	if err != nil {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := functions.NewGenericEmbed("New Member", message)

	if response != nil {
		_, err := session.ChannelMessageSendEmbed(channelID, response)
		if err != nil {
			log.Println(err)
		}
	}
}
