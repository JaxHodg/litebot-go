package modules

import (
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
)

func init() {
	manager.RegisterModule(
		&manager.Module{
			Name:        "Purge",
			Description: "Deletes a specific number of messages",
		},
	)
	manager.RegisterCommand(
		&manager.Command{
			Name:       "Purge",
			ModuleName: "Purge",

			Function:    cmdPurge,
			Description: "Deletes a specific number of messages (Maximum of 99)",
			HelpText:    "`{PREFIX}purge 10`",

			RequiredPermissions: discordgo.PermissionManageMessages,
			GuildOnly:           true,
		},
	)
	manager.RegisterEnable("Purge", true)
}

func cmdPurge(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	num := -1
	if len(args) >= 1 {
		var err error
		num, err = strconv.Atoi(args[0])
		if err != nil || num < 1 || num > 99 {
			num = -1
		}
	}
	if num == -1 {
		return functions.NewErrorEmbed("You must specify a number of messages to delete (Maximum 99)")
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

	if err = session.ChannelMessagesBulkDelete(event.Message.ChannelID, messagesToDelete); err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Unable to purge messages")
	}
	return nil
}
