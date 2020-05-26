package main

import (
	"github.com/bwmarrin/discordgo"
)

// Commands is a list of all possible Commands
var Commands map[string]*Command

// CommandEnvironment is a struct that contains all the info a command needs to run
type CommandEnvironment struct {
	Dm bool

	session *discordgo.Session
	event   *discordgo.MessageCreate

	Channel *discordgo.Channel //The channel the command was executed in
	Guild   *discordgo.Guild   //The guild the command was executed in
	Message *discordgo.Message //The message that triggered the command execution
	User    *discordgo.User    //The user that executed the command
	Member  *discordgo.Member  //The guild member that executed the command
}

// Command is a struct that contains all the data about a command
type Command struct {
	Function    func([]string, *CommandEnvironment) *discordgo.MessageEmbed
	Description string

	GuildOnly           bool
	RequiredPermissions int
}

// InitCommands creates all the commands and adds the to the slice
func InitCommands() {
	Commands = make(map[string]*Command)

	Commands["help"] = &Command{Function: cmdHelp, Description: "Displays this message"}
	Commands["ping"] = &Command{Function: cmdPing, Description: "Displays the ping"}
	Commands["kick"] = &Command{
		Function:            cmdKick,
		Description:         "Kicks the mentioned user",
		RequiredPermissions: discordgo.PermissionKickMembers,
		GuildOnly:           true}
	Commands["ban"] = &Command{
		Function:            cmdBan,
		Description:         "Bans the mentioned user",
		RequiredPermissions: discordgo.PermissionBanMembers,
		GuildOnly:           true}
}

// CallCommand calls the command and returns the embed it generates
func CallCommand(commandName string, args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if command, exists := Commands[commandName]; exists {
		if command.GuildOnly && env.Dm {
			return NewErrorEmbed("This command is for servers only")
		}
		if command.RequiredPermissions != 0 {
			if permissionsAllowed, isAdmin, _ := MemberHasPermission(env, command.RequiredPermissions); !permissionsAllowed && !isAdmin {
				return NewErrorEmbed("You do not have the required permissions to use " + commandName)
			}
		}
		return command.Function(args, env)
	}
	return nil
}
