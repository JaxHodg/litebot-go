package modules

import (
	"log"
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Spoil",
			Function:            cmdSpoil,
			Description:         "Marks the previous message as a spoiler",
			RequiredPermissions: discordgo.PermissionManageMessages,
			GuildOnly:           true,
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name:        "Spoil",
			Description: "Marks the previous message as a spoiler",
		},
	)
}

func cmdSpoil(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	messages, err := session.ChannelMessages(event.Message.ChannelID, 2, event.Message.ID, "", event.Message.ID)
	if err != nil {
		return functions.NewErrorEmbed("Unable to find any messages")
	}
	if len(messages[1].Embeds) > 0 {
		return functions.NewErrorEmbed("Unable to mark an embed as spoiler")
	}

	/**if len(args) == 0 {
		return functions.NewErrorEmbed("You must specify a user")
	}

	re := regexp.MustCompile(`[<][@](\d*)[>]`)
	substring := re.FindStringSubmatch(args[0])

	if len(substring) == 0 {
		return functions.NewErrorEmbed("You must specify a user")
	}

	userID := substring[1]

	user, err := session.GuildMember(event.Message.GuildID, userID)
	if err != nil {
		return functions.NewErrorEmbed("Invalid user")
	}

	err = session.GuildMemberDelete(event.Message.GuildID, userID)
	if err != nil {
		return functions.NewErrorEmbed("Unable to kick user")
	}**/

	content := "||" + strings.Replace(messages[1].Content, "|||", "", 0) + "||"
	user := messages[1].Author

	messagesToDelete := make([]string, 1)
	messagesToDelete[0] = event.Message.ID
	messagesToDelete = append(messagesToDelete, messages[1].ID)

	err = session.ChannelMessagesBulkDelete(event.Message.ChannelID, messagesToDelete)
	if err != nil {
		log.Println(err)
		return functions.NewErrorEmbed("Unable to delete original message")
	}

	return functions.NewRepostEmbed(content, user)
}
