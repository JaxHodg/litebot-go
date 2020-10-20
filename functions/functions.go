package functions

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// MemberHasPermission returns whether a member has the requested permission and whether they have admin
func MemberHasPermission(session *discordgo.Session, message *discordgo.Message, permission int) (bool, bool, error) { // Perm, Admin, Error
	if message.Member == nil {
		return false, false, errors.New("Nil member")
	}
	userPerm, err := session.UserChannelPermissions(message.Author.ID, message.ChannelID)
	if err != nil {
		fmt.Println(err)
		//return false, false, err
	}
	return userPerm&permission != 0, userPerm&discordgo.PermissionAdministrator != 0, nil
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
	err := session.UpdateStatus(0, "@lite-bot | "+strconv.Itoa(len(session.State.Guilds))+" Guilds")
	if err != nil {
		log.Println(err)
	}
}

func VerifyMessage(session *discordgo.Session, oldmessage *discordgo.Message) bool {
	message, err := session.ChannelMessage(oldmessage.ChannelID, oldmessage.ID)
	if err != nil {
		return false
	}
	if message.Author.ID == session.State.User.ID {
		return false
	}
	if message.Author.Bot {
		return false
	}

	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		return false
	}
	guild, err := session.State.Guild(channel.GuildID)
	if err != nil {
		return false
	}
	content := message.Content
	if content == "" {
		return false
	}
	_, err = session.GuildMember(guild.ID, message.Author.ID)
	if err != nil {
		return false
	}
	return true
}

func CanSpeak(session *discordgo.Session, channelID string) bool {
	userPerm, err := session.UserChannelPermissions(session.State.User.ID, channelID)
	if err != nil {
		log.Println(err)
		return false
	}
	if userPerm&discordgo.PermissionSendMessages != -1 || userPerm&discordgo.PermissionAdministrator != -1 {
		return true
	}
	return false
}
