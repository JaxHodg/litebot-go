package main

import (
	"github.com/bwmarrin/discordgo"
)

func cmdEnable(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a command")
	}

	if !isValidCmd(args[0]) {
		return NewErrorEmbed(args[0] + " is not a valid command")
	}

	EnableCommand(env.Guild, args[0])
	return NewGenericEmbed("Enabled", "Enabled "+args[0])
}
