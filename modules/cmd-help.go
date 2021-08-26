package modules

import (
	"strings"

	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name: "Help",

			Function:    cmdHelp,
			Description: "Lists all the available commands",
			HelpText:    "`{PREFIX}help`\nYou can also get more help on specific commands:\n`{PREFIX}help kick`",
		},
	)
}

func cmdHelp(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	helpEmbed := &discordgo.MessageEmbed{}
	helpEmbed.Color = 0xEBCB8B

	commandName := ""

	if len(args) >= 1 {
		commandName = strings.ToLower(args[0])
	}

	if command, exists := manager.Commands[commandName]; exists {
		helpEmbed.Title = strings.ToUpper(commandName)

		helpEmbed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Description",
				Value: command.Description,
			},
		}
		if command.HelpText != "" {
			HelpText := command.HelpText
			prefix, _ := state.GetData(event.GuildID, "Prefix", "Prefix")
			HelpText = strings.ReplaceAll(HelpText, "{PREFIX}", prefix)
			helpEmbed.Fields = append(helpEmbed.Fields, &discordgo.MessageEmbedField{
				Name:  "Examples",
				Value: HelpText,
			})
		}
	} else {
		helpEmbed.Title = "Help"

		helpEmbed.Fields = make([]*discordgo.MessageEmbedField, len(manager.Commands))

		i := 0
		for _, commandID := range []string{"help", "kick", "ban", "purge", "spoiler", "ping", "prefix", "joinmessage", "leavemessage", "block", "unblock", "enable", "disable"} /**manager.ListCommands()**/ {
			command, _ := manager.GetCommand(commandID)
			helpEmbed.Fields[i] = &discordgo.MessageEmbedField{Name: command.Name, Value: command.Description}
			i++
		}
		helpEmbed.Fields = append(helpEmbed.Fields, &discordgo.MessageEmbedField{
			Name:  "---------Get Support---------",
			Value: "Join the official support server [here](https://discord.gg/DSuy3CB)",
		})
	}
	return helpEmbed
}
