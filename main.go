package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"

	"github.com/bwmarrin/discordgo"

	"github.com/JaxHodg/litebot-go/state"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	godotenv.Load()

	discordToken := os.Getenv("DISCORD_TOKEN")

	/**
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
	**/
	dg, err := discordgo.New("Bot " + string(discordToken))
	if err != nil {
		log.Fatal(err)
	}

	dg.AddHandler(DiscordMessageCreate)
	dg.AddHandler(DiscordMessageUpdate)
	dg.AddHandler(DiscordGuildMemberAdd)
	dg.AddHandler(DiscordGuildMemberRemove)
	dg.AddHandler(DiscordConnect)

	dg.Identify.Intents = discordgo.MakeIntent(discordgo.IntentsAllWithoutPrivileged | discordgo.IntentsGuildMembers)

	err = dg.Open()
	if err != nil {
		log.Fatal(err)
		return
	}
	state.InitDB()
	state.MigrateDB()

	fmt.Println("Lite-bot is now running.  Press CTRL-C to exit.")

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	state.StopDB()
	// Cleanly close down the Discord session.
	dg.Close()
}
