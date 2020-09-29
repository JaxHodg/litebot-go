package modules

import (
	"log"
	"strings"

	"../functions"
	"../manager"
	"../state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterModule(
		&manager.Module{
			Name:        "BlockMessage",
			Description: "Blocks the specified term",
		},
	)
}

func BlockMessage(session *discordgo.Session, message *discordgo.Message) {
	_, isAdmin, err := functions.MemberHasPermission(session, message, discordgo.PermissionAdministrator)
	if err != nil {
		log.Println(err)
		return
	}
	if isAdmin == true {
		return
	}
	for _, s := range state.CheckList(message.GuildID, "blocked") {
		if strings.Contains(strings.ToLower(message.Content), s) {
			err := session.ChannelMessageDelete(message.ChannelID, message.ID)
			if err != nil {
				log.Println(err)
			}
			pm, err := session.UserChannelCreate(message.Author.ID)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = session.ChannelMessageSendEmbed(pm.ID, functions.NewGenericEmbed("Message Blocked", "Your message: ```"+message.Content+"``` was blocked because it contained a blocked term"))
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}
