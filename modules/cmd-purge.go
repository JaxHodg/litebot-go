package modules

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"../functions"
	"../manager"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Purge",
			Function:            cmdPurge,
			Description:         "Purges the specified number of messages",
			RequiredPermissions: discordgo.PermissionManageMessages,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Purge",
			Description: "Purges the specified number of messages",
		},
	)
}

func cmdPurge(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a number of messages to delete")
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return functions.NewErrorEmbed("You must specify a number of messages to delete")
	}

	if num < 1 || num > 99 {
		return functions.NewErrorEmbed("You can delete between 1 and 99 messages")
	}

	messagesToDelete := make([]string, 1)
	messagesToDelete[0] = event.Message.ID

	messages, err := session.ChannelMessages(event.Message.ChannelID, num, event.Message.ID, "", "")
	if err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Unable to purge messages")
	}

	for _, m := range messages {
		messagesToDelete = append(messagesToDelete, m.ID)
	}

	err = session.ChannelMessagesBulkDelete(event.Message.ChannelID, messagesToDelete)
	if err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Unable to purge messages")
	}
	return nil
}
