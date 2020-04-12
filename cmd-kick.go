package main

import (
	"regexp"
)

func cmdKick(args []string, env *CommandEnvironment) string {
	re := regexp.MustCompile(`[<][@](\d*)[>]`)
	user := re.FindString(args[0])

	env.session.GuildMemberDelete(env.message.GuildID, user)
	return "Kicked <@" + user + ">"
}
