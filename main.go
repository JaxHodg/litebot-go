package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	dg, err := discordgo.New("Bot " + "NTI4NDYwNDc4ODgwNjc3OTAz.XpJmiA.pIXI-9kFLN0KUIfEB-IiPIeoP3Q")
	dg.AddHandler(discordMessageCreate)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	InitCommands()

	fmt.Println("Bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
func discordMessageCreate(session *discordgo.Session, event *discordgo.MessageCreate) {
	message, err := session.ChannelMessage(event.ChannelID, event.ID) //Make it easier to keep track of what's happening
	if err != nil {
		return //Error finding message
	}

	if message.Author.ID == session.State.User.ID {
		return
	}
	if message.Author.Bot {
		return
	}
	if message.Content[0] == '!' {
		commandName := "ping" //TODO: Regex
		commandEnvironment := &CommandEnvironment{session, event}
		response := CallCommand(commandName, commandEnvironment)
		session.ChannelMessageSend(event.ChannelID, response)
	}

}
