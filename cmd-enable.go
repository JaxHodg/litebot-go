package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdEnable(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a command")
	}

	cmd := strings.ToLower(args[0])

	if !isValidCmd(cmd) {
		return NewErrorEmbed(cmd + " is not a valid command")
	} else if CheckEnabled(env.Guild, cmd) {
		return NewGenericEmbed("Enabled", cmd+" is already enabled")
	}

	EnableCommand(env.Guild, args[0])
	return NewGenericEmbed("Enabled", "Enabled "+args[0])
}
