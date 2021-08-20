package modules

import (
	"log"
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterEnable("Spoiler", true)
	manager.RegisterCommand(
		&manager.Command{
			Name:       "Spoiler",
			ModuleName: "Spoiler",

			Function:    cmdSpoiler,
			Description: "Marks the previous message as a spoiler",
			HelpText:    "`{PREFIX}spoiler`\nYou can also mark a specific message as a spoiler using its link\n`{PREFIX}spoiler https://discord.com/channels/123456789/123456789/123456789`",

			RequiredPermissions: discordgo.PermissionManageMessages,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Spoiler",
			Description: "Marks the previous message as a spoiler",
		},
	)
}

func cmdSpoiler(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	var spoilMessage *discordgo.Message
	if len(args) == 0 {
		messages, err := session.ChannelMessages(event.Message.ChannelID, 2, event.Message.ID, "", event.Message.ID)
		if err != nil {
			return functions.NewErrorEmbed("Unable to find any messages")
		}
		spoilMessage = messages[1]
	} else if len(args) == 1 {
		messageID := functions.ExtractMessageID(args[0])

		message, err := session.ChannelMessage(event.Message.ChannelID, messageID)
		if err != nil {
			return functions.NewErrorEmbed("Unable to find any messages")
		}
		spoilMessage = message
	}
	/**if len(messages[1].Embeds) > 0 {
		return functions.NewErrorEmbed("Unable to mark an embed as spoiler")
	}**/

	imageURL := ""
	if len(spoilMessage.Attachments) > 0 {
		imageURL = spoilMessage.Attachments[0].URL
	}
	if spoilMessage.Content+imageURL == "" {
		return functions.NewErrorEmbed("Nothing to mark as a spoiler")
	}
	content := "||" + strings.Replace(spoilMessage.Content, "||", "", -1) + imageURL + "||"
	user := spoilMessage.Author

	messagesToDelete := make([]string, 1)
	messagesToDelete[0] = event.Message.ID
	messagesToDelete = append(messagesToDelete, spoilMessage.ID)

	err := session.ChannelMessagesBulkDelete(event.Message.ChannelID, messagesToDelete)
	if err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Unable to delete original message")
	}

	return functions.NewRepostEmbed(content, user)
}
