package main

import (
	"log"
	"time"

	"github.com/bwmarrin/discordgo"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/modules"
	"github.com/JaxHodg/litebot-go/state"
)

func DiscordMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	if !functions.VerifyMessage(session, event.Message) {
		return //Error with getting message data
	}

	if event.Message.Content == "<@!405829095054770187>" {
		if functions.CanSpeak(session, event.Message.ChannelID) {

			prefix, _ := state.GetData(event.GuildID, "Prefix", "Prefix")

			_, err := session.ChannelMessageSendEmbed(event.Message.ChannelID, functions.NewGenericEmbed("Litebot", "Hi, I'm litebot. My prefix is `"+prefix+"`"))
			if err != nil {
				log.Println(err)
			}
		}
	}
	modules.BlockTerm(session, event.Message)
	if functions.CanSpeak(session, event.Message.ChannelID) {
		CallCommand(session, event)
	}
}

func DiscordMessageUpdate(session *discordgo.Session, event *discordgo.MessageUpdate) {
	if !functions.VerifyMessage(session, event.Message) {
		return //Error with message details
	}
	modules.BlockTerm(session, event.Message)
}

func DiscordConnect(session *discordgo.Session, event *discordgo.Connect) {
	time.Sleep(1 * time.Second)
	functions.UpdateStatus(session)
}

// DiscordGuildMemberAdd handles a user joining the server
func DiscordGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	modules.JoinMessage(session, event)
	functions.UpdateStatus(session)
}

// DiscordGuildMemberRemove handles a user leaving the server
func DiscordGuildMemberRemove(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	if event.Member.User.ID == session.State.User.ID {
		state.RemoveGuild(event.GuildID)
	}
	modules.LeaveMessage(session, event)
	functions.UpdateStatus(session)
}
