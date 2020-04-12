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

	Arguments           []CommandArgument
	RequiredPermissions int
}
type CommandArgument struct {
	//Used for help text
	Name        string //The name of the argument
	ArgType     string //The argument's type
	Description string //A description of the argument

	//Used for command argument parsing
	Value string //The value supplied with the argument
}

func InitCommands() {
	Commands = make(map[string]*Command)

	Commands["ping"] = &Command{Function: cmdPing}
	Commands["kick"] = &Command{
		Function:            cmdKick,
		RequiredPermissions: discordgo.PermissionKickMembers,
		Arguments: []CommandArgument{
			{Name: "user", ArgType: "mention"}},
	}
}
func CallCommand(commandName string, args []string, env *CommandEnvironment) string {
	if command, exists := Commands[commandName]; exists {
		if command.RequiredPermissions != 0 {
			if permissionsAllowed, _ := MemberHasPermission(env, command.RequiredPermissions); !permissionsAllowed {
				return "Error, you do not have the required permissions to use " + commandName
			}
		}

		return command.Function(args, env)
	}
	return ""
}
