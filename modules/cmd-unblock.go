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
			Name: "Unblock",

			Function:    cmdUnblock,
			Description: "Unblocks the specified term",
			HelpText:    "`{PREFIX}unblock frick`\nYou can also get a list of blocked terms\n`{PREFIX}unblock`",

			RequiredPermissions: discordgo.PermissionAdministrator,
			GuildOnly:           true,
		},
	)
}

func cmdUnblock(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	term := ""
	if len(args) >= 1 {
		term = strings.Join(args, " ")
		term = strings.TrimSpace(term)
		term = functions.NormaliseString(term)
		term = strings.ToLower(term)
	}

	if term == "" {
		pm, err := session.UserChannelCreate(event.Message.Author.ID)
		if err != nil {
			log.Println(err)
			return functions.NewErrorEmbed("Unable to send a DM containing blocked terms")
		}
		blockedList, err := state.GetList(event.Message.GuildID, "BlockTerm", "BlockedTerms")
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
	list, _ := state.GetList(event.Message.GuildID, "BlockTerm", "BlockedTerms")

	pos := functions.Find(list, term)
	if pos < 0 {
		return functions.NewErrorEmbed("`" + term + "` is not currently blocked")
	}

	state.RemoveToList(event.Message.GuildID, "BlockTerm", "BlockedTerms", term)
	return functions.NewGenericEmbed("BlockedTerms", "Successfully unblocked `"+term+"`")
}
