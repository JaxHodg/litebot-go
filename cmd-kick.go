package main

import (
	"fmt"
	"regexp"

	"github.com/bwmarrin/discordgo"
)

func cmdKick(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a user")
	}

	re := regexp.MustCompile(`<@!(\d*)>`)
	userID := re.FindStringSubmatch(args[0])[1]
	fmt.Println(userID)

	if userID == "" {
		return NewErrorEmbed("You must specify a user")
	}

	user, err := env.session.GuildMember(env.Guild.ID, userID)
	if err != nil {
		return NewErrorEmbed("Invalid user")
	}

	err = env.session.GuildMemberDelete(env.Guild.ID, userID)
	if err != nil {
		return NewErrorEmbed("Unable to kick user")
	}

	return NewGenericEmbed("Kick", "Kicked "+user.Mention())
}
