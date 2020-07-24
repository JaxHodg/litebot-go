package main

import (
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var Events map[string]*Event

type Event struct {
	Description string
	CanDisable  bool
}

func InitEvents() {
	Events = make(map[string]*Event)

	Events["JoinMessage"] = &Event{Description: "", CanDisable: true}
	Events["LeaveMessage"] = &Event{Description: "", CanDisable: true}
}

func DiscordGuildMemberAdd(session *discordgo.Session, event *discordgo.GuildMemberAdd) {
	guild, err := session.Guild(event.GuildID)
	if err != nil {
		return
	}

	if !CheckEnabled(guild, "JoinMessage") {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)
	channelID := re.FindStringSubmatch(CheckData(guild, "joinchannel"))[1]
	_, err = session.Channel(channelID)
	if err != nil {
		return
	}

	message := CheckData(guild, "joinmessage")
	if message == "" {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := NewGenericEmbed("New Member", message)

	if response != nil {
		session.ChannelMessageSendEmbed(channelID, response)
	}
}

func DiscordGuildMemberRemove(session *discordgo.Session, event *discordgo.GuildMemberRemove) {
	guild, err := session.Guild(event.GuildID)
	if err != nil {
		return
	}

	if !CheckEnabled(guild, "LeaveMessage") {
		return
	}

	re := regexp.MustCompile(`<#(\d*)>`)
	channelID := re.FindStringSubmatch(CheckData(guild, "leavechannel"))[1]
	_, err = session.Channel(channelID)
	if err != nil {
		return
	}

	message := CheckData(guild, "leavemessage")
	if message == "" {
		return
	}

	message = strings.ReplaceAll(message, "{user}", event.Mention())

	response := NewGenericEmbed("Member left", message)

	if response != nil {
		session.ChannelMessageSendEmbed(channelID, response)
	}
}
