package modules

import (
	"strings"

	"../manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:        "Help",
			Function:    cmdHelp,
			Description: "Displays this message",
		},
	)
}

func cmdHelp(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	helpEmbed := &discordgo.MessageEmbed{}
	helpEmbed.Color = 0xEBCB8B

	var commandName string

	if len(args) != 0 {
		commandName = strings.ToLower(args[0])
	} else {
		commandName = ""
	}

	if command, exists := manager.Commands[commandName]; exists {
		helpEmbed.Title = strings.ToUpper(commandName)

		helpEmbed.Fields = []*discordgo.MessageEmbedField{
			{
				Name:  "Description",
				Value: command.Description},
		}
	} else {
		helpEmbed.Title = "Help"

		helpEmbed.Fields = make([]*discordgo.MessageEmbedField, len(manager.Commands))

		i := 0
		for j := range manager.Commands {
			helpEmbed.Fields[i] = &discordgo.MessageEmbedField{Name: j, Value: manager.Commands[j].Description}
			i++
		}
	}
	return helpEmbed
}
