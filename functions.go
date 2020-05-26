package main

import (
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
