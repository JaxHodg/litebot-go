package main

import (
	"github.com/bwmarrin/discordgo"

	"./functions"
	"./modules"
	"./state"
)

func DiscordMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	functions.UpdateStatus(session)
	content := event.Message.Content

	if content == "" {
		return //The message was empty
	} else if content == "<@!405829095054770187>" {
		session.ChannelMessageSendEmbed(event.Message.ChannelID, functions.NewGenericEmbed("Litebot", "Hi, I'm litebot. My prefix is `"+state.CheckData(event.GuildID, "prefix")+"`"))
	}
	modules.BlockMessage(session, event)

	CallCommand(session, event)
}

// DiscordGuildMemberAdd handles a user joining the server
func DiscordGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	functions.UpdateStatus(session)
	modules.JoinMessage(session, event)
}

// DiscordGuildMemberRemove handles a user leaving the server
func DiscordGuildMemberRemove(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	functions.UpdateStatus(session)
	modules.LeaveMessage(session, event)
}
