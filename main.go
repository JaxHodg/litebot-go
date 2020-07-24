package main

import (
	"fmt"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + "NTI4NDYwNDc4ODgwNjc3OTAz.XpJmiA.pIXI-9kFLN0KUIfEB-IiPIeoP3Q")

	dg.AddHandler(discordMessageCreate)
	dg.AddHandler(DiscordGuildMemberAdd)
	dg.AddHandler(DiscordGuildMemberRemove)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	InitCommands()
	InitEvents()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}

func discordMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {

	message, err := session.ChannelMessage(event.ChannelID, event.ID)
	if err != nil {
		return //Error finding message
	}

	if message.Author.ID == session.State.User.ID {
		return //Don't respond to itself
	}
	if message.Author.Bot {
		return //Don't respond to bots
	}

	channel, err := session.State.Channel(message.ChannelID)
	if err != nil {
		return //Error finding the channel
	}

	content := message.Content
	if content == "" {
		return //The message was empty
	}

	Dm := CheckIfDm(session, event)

	var guild *discordgo.Guild = nil
	var member *discordgo.Member = nil

	if !Dm {
		guild, err = session.State.Guild(channel.GuildID)
		if err != nil {
			return
		}
		member, err = session.GuildMember(guild.ID, message.Author.ID)
		if err != nil {
			return //Error finding the guild member
		}
	}

	if content[0] == '!' {
		cmdMsg := strings.TrimPrefix(content, "!")
		cmd := strings.Split(cmdMsg, " ")

		commandEnvironment := &CommandEnvironment{Dm, session, event, channel, guild, message, member.User, member}

		response := CallCommand(cmd[0], cmd[1:], commandEnvironment)

		if response != nil {
			session.ChannelMessageSendEmbed(event.ChannelID, response)
		}
	}

}
