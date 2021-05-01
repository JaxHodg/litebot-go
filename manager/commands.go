package manager

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Commands is a list of all possible Commands
var Commands map[string]*Command

// Command is a struct that contains all the data about a command
type Command struct {
	Name string

	Function    func([]string, *discordgo.Session, *discordgo.MessageCreate) *discordgo.MessageEmbed
	Description string

	GuildOnly           bool
	RequiredPermissions int64
}

func RegisterCommand(command *Command) {
	if Commands == nil {
		Commands = make(map[string]*Command)
	}
	Commands[strings.ToLower(command.Name)] = command
}

func IsValidCommand(command string) bool {
	_, exists := Commands[command]
	return exists
}
