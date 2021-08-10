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
	manager.RegisterEnable("BlockMessage", false)
	manager.RegisterModule(
		&manager.Module{
			Name:        "BlockMessage",
			Description: "Blocks the specified term",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name: "Blocked",

			ModuleName: "BlockMessage",

			DefaultValue: []string{},
		},
	)
}

// BlockMessage checks if message contains any blocked terms
func BlockMessage(session *discordgo.Session, message *discordgo.Message) {
	if message.Member == nil {
		return
	}
	_, isAdmin, err := functions.MemberHasPermission(session, message, discordgo.PermissionAdministrator)
	if err != nil {
		log.Println(err)
		return
	}
	if isAdmin {
		return
	}
	list, err := state.GetList(message.GuildID, "BlockMessage", "Blocked")
	if err == nil {
		for _, s := range list {
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
}
