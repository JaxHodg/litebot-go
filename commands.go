package main

import (
	"github.com/bwmarrin/discordgo"
)

var Commands map[string]*Command

type CommandEnvironment struct {
	session *discordgo.Session
	event   *discordgo.MessageCreate
	message *discordgo.Message
}
type Command struct {
	Function func([]string, *CommandEnvironment) string

	//Arguments           []CommandArgument
	RequiredPermissions int
}

func InitCommands() {
	Commands = make(map[string]*Command)

	Commands["ping"] = &Command{Function: cmdPing}
	Commands["kick"] = &Command{Function: cmdPing, RequiredPermissions: discordgo.PermissionAdministrator}
}
func CallCommand(commandName string, args []string, env *CommandEnvironment) string {
	if command, exists := Commands[commandName]; exists {
		if command.RequiredPermissions != 0 {
			if permissionsAllowed, _ := MemberHasPermission(env.session, env.message.GuildID, env.message.Author.ID, command.RequiredPermissions); !permissionsAllowed {
				return "Error, you do not have the required permissions to use " + commandName
			}
		}

		return command.Function(args, env)
	}
	return ""
}
