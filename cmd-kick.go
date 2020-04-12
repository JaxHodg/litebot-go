package main

import (
	"regexp"
)

func cmdKick(args []string, env *CommandEnvironment) string {
	re := regexp.MustCompile(`[<][@][!](\d*)[>]`)
	user := re.FindString(args[0])

	if user == "" {
		return "Error: You must specify a user"
	}

	err := env.session.GuildMemberDelete(env.Guild.ID, user)

	if err != nil {
		return "Error kicking " + user
	}

	return "Kicked " + user
}
