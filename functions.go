package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// MemberHasPermission returns whether a member has the requested permission and whether they have admin
func MemberHasPermission(env *CommandEnvironment, permission int) (bool, bool, error) { // Perm, Admin, Error
	// Iterate through the role IDs stored in member.Roles
	// to check permissions
	for _, roleID := range env.Member.Roles {
		role, err := env.session.State.Role(env.Guild.ID, roleID)
		if err != nil {
			return false, false, err
		}
		if role.Permissions&permission != 0 {
			if role.Permissions&discordgo.PermissionAdministrator != 0 {
				return true, true, nil
			}
			return true, false, nil
		}
	}

	return false, false, nil
}

//CheckIfDm returns true if the message came from a Dm
func CheckIfDm(s *discordgo.Session, m *discordgo.MessageCreate) bool {
	channel, err := s.State.Channel(m.ChannelID)

	if err != nil {
		if channel, err = s.Channel(m.ChannelID); err != nil {
			return false
		}
	}
	return channel.Type == discordgo.ChannelTypeDM
}

//NewGenericEmbed returns a generic embed
func NewGenericEmbed(embedTitle, embedMsg string) *discordgo.MessageEmbed {
	genericEmbed := &discordgo.MessageEmbed{}
	genericEmbed.Title = embedTitle
	genericEmbed.Description = embedMsg
	genericEmbed.Color = 0xD8DEE9
	return genericEmbed
}

//NewErrorEmbed returns an error embed
func NewErrorEmbed(embedMsg string) *discordgo.MessageEmbed {
	errorEmbed := &discordgo.MessageEmbed{}
	errorEmbed.Title = "ERROR"
	errorEmbed.Description = embedMsg
	errorEmbed.Color = 0xBF616A
	return errorEmbed
}

// Contains checks if an array contains a string
func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func Find(s []string, e string) int {

	for i := 0; i < len(s); i++ {
		if s[i] == e {
			return i
		}
	}
	return -1
}

func Remove(s []string, i int) []string {
	s[len(s)-1], s[i] = s[i], s[len(s)-1]
	return s[:len(s)-1]
}

func UpdateStatus(session *discordgo.Session) {
	session.UpdateStatus(0, "@lite-bot | "+strconv.Itoa(len(session.State.Guilds))+" Guilds")
}
