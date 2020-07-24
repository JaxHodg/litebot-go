package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdDisable(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a command")
	}

	cmd := strings.ToLower(args[0])

	if !isValidCmd(cmd) {
		return NewErrorEmbed(cmd + " is not a valid command")
	} else if !Commands[cmd].CanDisable {
		return NewErrorEmbed(cmd + " cannot be disabled")
	} else if !CheckEnabled(env.Guild, cmd) {
		return NewGenericEmbed("Disabled", cmd+" is already disabled")
	}

	DisableCommand(env.Guild, (cmd))
	return NewGenericEmbed("Disabled", "Disabled "+cmd)
}
