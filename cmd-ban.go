package main

import (
	"regexp"
)

func cmdBan(args []string, env *CommandEnvironment) string {
	re := regexp.MustCompile(`[<][@][!](\d*)[>]`)
	user := re.FindString(args[0])

	if user == "" {
		return "Error: You must specify a user"
	}

	err := env.session.GuildBanCreate(env.Guild.ID, user, 0)

	if err != nil {
		return "Error banning " + user
	}

	return "Banned " + user
}
