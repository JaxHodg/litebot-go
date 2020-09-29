package main

import (
	"log"
	"regexp"
	"strings"

	"./functions"
	"./manager"
	"./state"

	"github.com/bwmarrin/discordgo"
)

// CallCommand calls the command and returns the embed it generates
func CallCommand(session *discordgo.Session, event *discordgo.MessageCreate) {
	var response *discordgo.MessageEmbed

	prefix := state.CheckData(event.GuildID, "prefix")
	if prefix == "" {
		prefix = "!"
	}

	re := regexp.MustCompile("[" + prefix + "](\\w*)")

	if !re.MatchString(event.Message.Content) {
		return
	}

	args := strings.Split(event.Message.Content, " ")
	commandName := strings.ToLower(strings.TrimPrefix(args[0], prefix))

	if command, exists := manager.Commands[commandName]; exists {
		if !state.CheckEnabled(event.GuildID, commandName) {
			response = functions.NewErrorEmbed(commandName + " is disabled")
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
	}
	_, err := session.ChannelMessageSendEmbed(event.ChannelID, response)
	if err != nil {
		log.Println(err)
	}
}
