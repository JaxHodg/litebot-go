package main

import "github.com/bwmarrin/discordgo"

var GuildEnabled = make(map[string]map[string]bool)

//InitState creates save data for all guilds
func InitState(s *discordgo.Session) {
	for _, guild := range s.State.Guilds {
		if _, ok := GuildEnabled[guild.ID]; ok {
			GuildEnabled[guild.ID] = make(map[string]bool)
			for cmdName, cmd := range Commands {
				if cmd.CanDisable {
					GuildEnabled[guild.ID][cmdName] = true
				}
			}
		}
	}
}

// func disableCommand(guild)
