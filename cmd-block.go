package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func cmdBlock(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	data := strings.ToLower(strings.Join(args, " "))
	if data == "" {
		return NewErrorEmbed("You must specify a term to block")
	}
	if Find(CheckList(env.Guild, "blocked"), data) != -1 {
		return NewErrorEmbed("`" + data + "` is already blocked")
	}

	AddToList(env.Guild, "blocked", data)
	return NewGenericEmbed("Blocked", "Successfully blocked `"+data+"`")
}
