package modules

import (
	"strings"

	"github.com/JaxHodg/litebot-go/manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:        "Help",
			Function:    cmdHelp,
			Description: "Lists all the available commands",
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
		for _, commandID := range []string{"help", "kick", "ban", "ping", "purge", "spoiler", "enable", "disable", "block", "unblock", "set"} {
			command, _ := manager.GetCommand(commandID)
			helpEmbed.Fields[i] = &discordgo.MessageEmbedField{Name: command.Name, Value: command.Description}
			i++
		}
	}
	return helpEmbed
}
