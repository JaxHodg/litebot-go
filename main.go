package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"

	"./state"

	"github.com/bwmarrin/discordgo"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	file, err := os.Open("./key.config")
	if err != nil {
		os.Create("./key.config")
		fmt.Println("Place the key in key.config")
		os.Exit(0)
	}
	key, err := ioutil.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}
	key = bytes.TrimSuffix(key, []byte{'\n'})

	dg, err := discordgo.New("Bot " + string(key))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(DiscordMessageCreate)
	dg.AddHandler(DiscordMessageUpdate)
	dg.AddHandler(DiscordGuildMemberAdd)
	dg.AddHandler(DiscordGuildMemberRemove)
	dg.AddHandler(DiscordConnect)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	state.InitState()

	fmt.Println("Lite-bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
