package main

import (
	"strconv"

	"github.com/bwmarrin/discordgo"
)

func cmdPurge(args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args) == 0 {
		return NewErrorEmbed("You must specify a number of messages to delete")
	}
	num, err := strconv.Atoi(args[0])
	if err != nil {
		return NewErrorEmbed("You must specify a number of messages to delete")
	}

	if num < 1 || num > 99 {
		return NewErrorEmbed("You can delete between 1 and 99 messages")
	}

	messagesToDelete := make([]string, 1)
	messagesToDelete[0] = env.Message.ID

	messages, _ := env.session.ChannelMessages(env.Channel.ID, num, env.Message.ID, "", "")

	for _, m := range messages {
		messagesToDelete = append(messagesToDelete, m.ID)
	}

	env.session.ChannelMessagesBulkDelete(env.Channel.ID, messagesToDelete)

	return nil
}
