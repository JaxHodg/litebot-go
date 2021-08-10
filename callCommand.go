package main

import (
	"log"
	"regexp"
	"strings"

	"github.com/bwmarrin/discordgo"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
	"github.com/JaxHodg/litebot-go/state"
)

// CallCommand calls the command and returns the embed it generates
func CallCommand(session *discordgo.Session, event *discordgo.MessageCreate) {
	var response *discordgo.MessageEmbed

	prefix, _ := state.GetData(event.GuildID, "Prefix", "Prefix")

	re := regexp.MustCompile("[" + prefix + "](\\w*)")

	if !re.MatchString(event.Message.Content) {
		return
	}

	args := strings.Split(event.Message.Content, " ")
	commandName := strings.ToLower(strings.TrimPrefix(args[0], prefix))

	command, err := manager.GetCommand(commandName)
	if err != nil {
		return
	}
	if err == nil {
		enabled, err := state.GetEnabled(event.GuildID, command.ModuleName)
		if err == nil && !enabled {
			_, err := session.ChannelMessageSendEmbed(event.ChannelID, functions.NewErrorEmbed(commandName+" is disabled"))
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	if command.RequiredPermissions != 0 {
		if permissionsAllowed, isAdmin, err := functions.MemberHasPermission(session, event.Message, command.RequiredPermissions); !permissionsAllowed && !isAdmin || err != nil {
			_, err := session.ChannelMessageSendEmbed(event.ChannelID, functions.NewErrorEmbed("You do not have the required permissions to use "+commandName))
			if err != nil {
				log.Println(err)
			}
			return
		}
	}
	response = command.Function(args[1:], session, event)
	if response == nil {
		return
	}
	_, err = session.ChannelMessageSendEmbed(event.ChannelID, response)
	if err != nil {
		log.Println(err)
	}
}
