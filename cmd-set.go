package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdSet(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewGenericEmbed("Set values", "You may set the following values: "+strings.Join(DataValues, ", "))
	} else if len(args) < 2 {
		return NewErrorEmbed("You must specify what to set " + args[0] + " to")
	}

	value := strings.ToLower(args[0])

	if !Contains(DataValues, value) {
		return NewErrorEmbed(args[0] + " is an invalid value")
	}

	data := strings.Join(args[1:], " ")

	SetData(env.Guild, args[0], data)

	return NewGenericEmbed("Set", "Successfully set `"+args[0]+"` to `"+data+"`")
}
