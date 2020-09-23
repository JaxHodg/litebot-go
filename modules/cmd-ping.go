package modules

import (
	"../functions"
	"../manager"
	"github.com/bwmarrin/discordgo"
)

func init() {
	manager.RegisterCommand(
		&manager.Command{
			Name:        "Ping",
			Function:    cmdPing,
			Description: "Displays the current ping",
		},
	)
}

func cmdPing(args []string, session *discordgo.Session, event *discordgo.MessageCreate) *discordgo.MessageEmbed {
	return functions.NewGenericEmbed("Ping", "Pong")
}
