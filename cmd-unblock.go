package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdUnblock(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	data := strings.Join(args, " ")
	if data == "" {
		return NewErrorEmbed("You must specify a term to unblock")
	}
	pos := Find(CheckList(env.Guild, "blocked"), data)
	if pos < 0 {
		return NewErrorEmbed("`" + data + "` is not currently blocked")
	}

	RemoveFromList(env.Guild, "blocked", pos)
	return NewGenericEmbed("Blocked", "Successfully unblocked `"+data+"`")
}
