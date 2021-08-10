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
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Unblock",
			Function:            cmdUnblock,
			Description:         "Unblocks the specified term",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdUnblock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	data := strings.Join(args, " ")
	if len(args) == 0 {
		pm, err := session.UserChannelCreate(event.Message.Author.ID)
		if err != nil {
			log.Println(err)
			return functions.NewErrorEmbed("Unable to send a DM containing blocked terms")
		}
		blockedList, err := state.GetList(event.Message.GuildID, "BlockMessage", "blocked")
		if err != nil {
			log.Println(err)
			return functions.NewErrorEmbed("Unable to send a DM containing blocked terms")
		}

		embed := &discordgo.MessageEmbed{}
		embed.Color = 0xEBCB8B
		embed.Title = "Blocked Terms"

		for i := range blockedList {
			embed.Description = embed.Description + "```" + blockedList[i] + "```"
		}
		if len(blockedList) == 0 {
			embed.Description = "No blocked terms"
		}

		_, err = session.ChannelMessageSendEmbed(pm.ID, embed)
		if err != nil {
			log.Println(err)
		}
		return functions.NewGenericEmbed("Blocked Terms", "Check your DMs for a list of blocked terms")
	}
	list, err := state.GetList(event.Message.GuildID, "BlockMessage", "Blocked")
	if err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Error getting blocked terms")
	}
	pos := functions.Find(list, data)
	if pos < 0 {
		return functions.NewErrorEmbed("`" + data + "` is not currently blocked")
	}

	state.RemoveToList(event.Message.GuildID, "BlockMessage", "Blocked", data)
	return functions.NewGenericEmbed("Blocked", "Successfully unblocked `"+data+"`")
}
