package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdHelp(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	helpEmbed := &discordgo.MessageEmbed{}
	var commandName string

	if len(args) != 0 {
		commandName = strings.ToLower(args[0])
	} else {
		commandName = ""
	}

	if Commands[commandName] != nil {
		helpEmbed.Title = strings.ToUpper(commandName)

		helpEmbed.Fields = []*discordgo.MessageEmbedField{
			&discordgo.MessageEmbedField{
				Name:  "Description",
				Value: Commands[commandName].Description}}

	} else {
		helpEmbed.Title = "Help"

		helpEmbed.Fields = make([]*discordgo.MessageEmbedField, len(Commands))

		i := 0
		for j := range Commands {
			helpEmbed.Fields[i] = &discordgo.MessageEmbedField{Name: j, Value: Commands[j].Description}
			i++
		}
	}
	return helpEmbed
}
