package main

import (
		"github.com/bwmarrin/discordgo"
		"fmt"
		"strings"
	)

func cmdSet (args []string, env *CommandEnvironment) *discordgo.MessageEmbed {
	if len(args)<1{
		return NewErrorEmbed("You must specify which value to set")// TODO: List the possible values
	} else if len(args)<2{
		return NewErrorEmbed("You must specify what to set "+args[0]+" to")		
	}
	
	value := strings.ToLower(args[0])

	if !Contains(DataValues,value) {
		return NewErrorEmbed(args[0]+" is an invalid value")
	}

	data := strings.Join(args[1:], " ")

	SetData(env.Guild,args[0], data)
	fmt.Println(GuildData)
	return NewGenericEmbed("Set","Successfully set "+args[0]+" to "+data)
}
