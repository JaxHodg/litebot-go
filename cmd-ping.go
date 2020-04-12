package main

import "github.com/bwmarrin/discordgo"

func cmdPing(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	return NewGenericEmbed("Ping", "Pong")
}
