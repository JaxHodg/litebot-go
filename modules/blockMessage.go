package modules

import (
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

func BlockMessage(session *discordgo.Session, event *discordgo.MessageCreate) {
	_, isAdmin, _ := functions.MemberHasPermission(session, event, discordgo.PermissionAdministrator)

	for _, s := range state.CheckList(event.GuildID, "blocked") {
		if strings.Contains(strings.ToLower(event.Message.Content), s) && !isAdmin {
			session.ChannelMessageDelete(event.Message.ChannelID, event.Message.ID)
			pm, err := session.UserChannelCreate(event.Message.Author.ID)
			if err != nil {
				return
			}
			session.ChannelMessageSendEmbed(pm.ID, functions.NewGenericEmbed("Message Blocked", "Your message: ```"+event.Message.Content+"``` was blocked because it contained a blocked term"))
			return
		}
	}
}
