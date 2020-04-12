package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func cmdKick(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	re := regexp.MustCompile(`[<][@][!](\d*)[>]`)
	user := re.FindString(args[0])

	if user == "" {
		return NewGenericEmbed("Error:", "You must specify a user")
	}

	err := env.session.GuildMemberDelete(env.Guild.ID, user)

	if err != nil {
		return NewGenericEmbed("Error:", "Unable to kick user")
	}

	return NewGenericEmbed("Kick", "Kicked "+user)
}
