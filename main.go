package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"

	"./functions"
	"./state"

	"github.com/bwmarrin/discordgo"
)

func main() {
	file, err := os.Open("./key.config")
	if err != nil {
		os.Create("./key.config")
		fmt.Println("Place the key in key.config")
		os.Exit(0)
	}
	key, _ := ioutil.ReadAll(file)
	dg, err := discordgo.New("Bot " + string(key))

	dg.AddHandler(DiscordMessageCreate)
	dg.AddHandler(DiscordMessageUpdate)
	dg.AddHandler(DiscordGuildMemberAdd)
	dg.AddHandler(DiscordGuildMemberRemove)

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}
	state.InitState()
	functions.UpdateStatus(dg)

	fmt.Println("Lite-bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
