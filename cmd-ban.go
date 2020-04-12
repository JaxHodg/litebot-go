package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func cmdBan(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	re := regexp.MustCompile(`[<][@][!](\d*)[>]`)
	user := re.FindString(args[0])

	if user == "" {
		return NewErrorEmbed("You must specify a user")
	}

	err := env.session.GuildBanCreate(env.Guild.ID, user, 0)

	if err != nil {
		return NewErrorEmbed("Unable to ban user")
	}

	return NewGenericEmbed("Ban", "Banned "+user)
}
