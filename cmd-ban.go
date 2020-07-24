package main

import (
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func cmdBan(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a user")
	}

	re := regexp.MustCompile(`<@!(\d*)>`)
	userID := re.FindStringSubmatch(args[0])[1]

	if userID == "" {
		return NewErrorEmbed("You must specify a user")
	}

	user, err := env.session.GuildMember(env.Guild.ID, userID)
	if err != nil {
		return NewErrorEmbed("Invalid user")
	}

	err = env.session.GuildBanCreate(env.Guild.ID, userID, 0)
	if err != nil {
		return NewErrorEmbed("Unable to ban user")
	}

	return NewGenericEmbed("Ban", "Banned  "+user.Mention())
}
