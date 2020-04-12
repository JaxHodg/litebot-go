package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func cmdKick(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	re := regexp.MustCompile(`[<][@][!](\d*)[>]`)
	user := re.FindString(args[0])

	if user == "" {
		return NewErrorEmbed("You must specify a user")
	}

	err := env.session.GuildMemberDelete(env.Guild.ID, user)

	if err != nil {
		return NewErrorEmbed("Unable to kick user")
	}

	return NewGenericEmbed("Kick", "Kicked "+user)
}
