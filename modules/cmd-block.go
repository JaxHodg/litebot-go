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
			Name: "Block",

			Function:    cmdBlock,
			Description: "Blocks the specified term",
			HelpText:    "`{PREFIX}block frick`\nTip: Admin messages won't be blocked",

			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdBlock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	term := ""
	if len(args) >= 1 {
		term = strings.Join(args, " ")
		term = strings.TrimSpace(term)
		term = functions.NormaliseString(term)
		term = strings.ToLower(term)
	}

	if term == "" {
		return functions.NewErrorEmbed("You must specify a term to block")
	}

	list, _ := state.GetList(event.Message.GuildID, "BlockTerm", "BlockedTerms")
	if functions.Find(list, term) != -1 {
		return functions.NewErrorEmbed("`" + term + "` is already blocked")
	}

	state.AddToList(event.Message.GuildID, "BlockTerm", "BlockedTerms", term)
	return functions.NewGenericEmbed("BlockedTerms", "Successfully blocked `"+term+"`")
}
