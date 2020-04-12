package main

import "github.com/bwmarrin/discordgo"

var Commands map[string]*Command

type CommandEnvironment struct {
	session *discordgo.Session
	event   *discordgo.MessageCreate
}
type Command struct {
	Function func([]string, *CommandEnvironment) string
}

func InitCommands() {
	Commands = make(map[string]*Command)

	Commands["ping"] = &Command{Function: cmdPing}
}
func CallCommand(commandName string, env *CommandEnvironment) string {
	command := Commands[commandName]

	var args []string

	return command.Function(args, env)
}
