package main

import (
	"github.com/bwmarrin/discordgo"
)

func cmdDisable(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a command")
	}

	if !isValidCmd(args[0]) {
		return NewErrorEmbed(args[0] + " is not a valid command")
	}

	DisableCommand(env.Guild, args[0])
	return NewGenericEmbed("Disabled", "Disabled "+args[0])
}
