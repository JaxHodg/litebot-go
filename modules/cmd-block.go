package modules

import (
	"strings"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"

	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:                "Block",
			Function:            cmdBlock,
			Description:         "Blocks the specified term",
			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdBlock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	data := strings.ToLower(strings.Join(args, " "))
	if data == "" {
		return functions.NewErrorEmbed("You must specify a term to block")
	}
	list, _ := state.GetList(event.Message.GuildID, "BlockTerm", "BlockedTerms")
	/**if err != nil {
		return functions.NewErrorEmbed("Error blocking term")
	}**/
	if functions.Find(list, data) != -1 {
		return functions.NewErrorEmbed("`" + data + "` is already blocked")
	}
	state.AddToList(event.Message.GuildID, "BlockTerm", "BlockedTerms", data)
	return functions.NewGenericEmbed("BlockedTerms", "Successfully blocked `"+data+"`")
}
