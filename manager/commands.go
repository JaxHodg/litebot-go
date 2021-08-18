package manager

import (
	"errors"
	"sort"
	"strings"

	"github.com/bwmarrin/discordgo"
)

// Commands is a list of all possible Commands
var Commands map[string]*Command

// Command is a struct that contains all the data about a command
type Command struct {
	Name       string
	ModuleName string

	Function    func([]string, *discordgo.Session, *discordgo.MessageCreate) *discordgo.MessageEmbed
	Description string

	GuildOnly           bool
	RequiredPermissions int64
}

func RegisterCommand(command *Command) {
	commandID := strings.ToLower(command.Name)

	if Commands == nil {
		Commands = make(map[string]*Command)
	}
	Commands[commandID] = command
}

func GetCommand(commandID string) (*Command, error) {
	commandID = strings.ToLower(commandID)

	if _, ok := Commands[commandID]; ok {
		return Commands[commandID], nil
	}
	return nil, errors.New("invalid command")
}

func IsValidCommand(commandID string) bool {
	commandID = strings.ToLower(commandID)

	_, exists := Commands[commandID]
	return exists
}

func ListCommands() []string {
	keys := make([]string, 0, len(Commands))
	for k := range Commands {

		keys = append(keys, k)

	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	return keys
}
