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
			Name:        "BlockTerm",
			Description: "Blocks the specified term",
		},
	)
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "BlockedTerms",
			ModuleName:   "BlockTerm",
			DefaultValue: []string{},
		},
	)
	manager.RegisterEnable("BlockTerm", false)
}

// BlockTerm checks if message contains any blocked terms
func BlockTerm(session *discordgo.Session, message *discordgo.Message) {
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
	list, _ := state.GetList(message.GuildID, "BlockTerm", "BlockedTerms")

	text := message.Content
	text = strings.TrimSpace(text)
	text = functions.NormaliseString(text)
	text = strings.ToLower(text)

	for _, s := range list {
		if strings.Contains(text, s) {
			err := session.ChannelMessageDelete(message.ChannelID, message.ID)
			if err != nil {
				log.Println(err)
			}
			pm, err := session.UserChannelCreate(message.Author.ID)
			if err != nil {
				log.Println(err)
				return
			}
			_, err = session.ChannelMessageSendEmbed(pm.ID, functions.NewGenericEmbed("Message Blocked", "Your message was removed because it contained a blocked term```"+message.Content+"```"))
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
}
