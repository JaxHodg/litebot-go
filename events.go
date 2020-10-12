package main

import (
	"log"

	"github.com/bwmarrin/discordgo"

	"./functions"
	"./modules"
	"./state"
)

func DiscordMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if !functions.VerifyMessage(session, event.Message) {
		return //Error with getting message data
	}

	if event.Message.Content == "<@!405829095054770187>" {
		if functions.CanSpeak(session, event.Message.ChannelID) {
			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, functions.NewGenericEmbed("Litebot", "Hi, I'm litebot. My prefix is `"+state.CheckData(event.GuildID, "prefix")+"`"))
			if err != nil {
				log.Println(err)
			}
		}
	}
	modules.BlockMessage(session, event.Message)

	if functions.CanSpeak(session, event.Message.ChannelID) {
		CallCommand(session, event)
	}

}

func DiscordMessageUpdate(session *discordgo.Session, event *discordgo.MessageUpdate) {
	if !functions.VerifyMessage(session, event.Message) {
		return //Error with message details
	}
	modules.BlockMessage(session, event.Message)
}

func DiscordConnect(session *discordgo.Session, event *discordgo.Connect) {
	functions.UpdateStatus(session)
}

// DiscordGuildMemberAdd handles a user joining the server
func DiscordGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	modules.JoinMessage(session, event)
	functions.UpdateStatus(session)
}

// DiscordGuildMemberRemove handles a user leaving the server
func DiscordGuildMemberRemove(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	modules.LeaveMessage(session, event)
	functions.UpdateStatus(session)
}
