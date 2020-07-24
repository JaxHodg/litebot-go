package main

import (
	"github.com/bwmarrin/discordgo"
)

var GuildEnabled = make(map[string]map[string]bool)

//InitAllState creates save data for all guilds, won't be in final version of bot
func InitAllState(s *discordgo.Session) {
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

//VerifyState creates state if missing for selected Guild
func VerifyState(guild *discordgo.Guild) {
	if _, ok := GuildEnabled[guild.ID]; !ok {
		GuildEnabled[guild.ID] = make(map[string]bool)
	}
	for cmdName, cmd := range Commands {
		if _, ok := GuildEnabled[guild.ID][cmdName]; !ok && cmd.CanDisable {
			GuildEnabled[guild.ID][cmdName] = true
		}
	}
}

func CheckEnabled(guild *discordgo.Guild, command string) bool {
	VerifyState(guild)
	return GuildEnabled[guild.ID][command]
}

func EnableCommand(guild *discordgo.Guild, command string) {
	VerifyState(guild)
	GuildEnabled[guild.ID][command] = true
}

func DisableCommand(guild *discordgo.Guild, command string) { //Add more error detection
	VerifyState(guild)
	GuildEnabled[guild.ID][command] = false
}
