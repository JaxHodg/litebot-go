package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

// GuildEnabled store data about which commands/events are enabled
var GuildEnabled = make(map[string]map[string]bool)

// GuildData stores strings for each guilds
var GuildData = make(map[string]map[string]string)

// DataValues contains the data that is stored in GuildData
var DataValues = []string{"prefix", "joinmessage", "joinchannel", "leavemessage", "leavechannel"}

// InitState loads GuildEnabled & GuildData from files if they exist, creates them if they don't
func InitState() {
	file, err := os.Open("./GuildEnabled.json")
	if err != nil {
		os.Create("./GuildEnabled.json")
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildEnabled)
	file, err = os.Open("./GuildData.json")
	if err != nil {
		os.Create("./GuildData.json")
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildData)
}

// DumpEnabled dumps the data for which commands are enabled to json
func DumpEnabled() {
	jsonData, _ := json.Marshal(GuildEnabled)
	jsonFile, _ := os.Create("./GuildEnabled.json")
	jsonFile.Write(jsonData)
}

// DumpData dumps the strings to json
func DumpData() {
	jsonData, _ := json.Marshal(GuildData)
	jsonFile, _ := os.Create("./GuildData.json")
	jsonFile.Write(jsonData)
}

//VerifyState creates state if missing for selected Guild
func VerifyState(guild *discordgo.Guild) {
	// Verifys guild has a section under GuildEnabled
	if _, ok := GuildEnabled[guild.ID]; !ok {
		GuildEnabled[guild.ID] = make(map[string]bool)
	}
	// Verifys each command that can be disabled is listed
	for cmdName, cmd := range Commands {
		if _, ok := GuildEnabled[guild.ID][cmdName]; !ok && cmd.CanDisable {
			GuildEnabled[guild.ID][cmdName] = true
		}
	}
	// Verifys each event that can be disabled is listed
	for eventName, event := range Events {
		if _, ok := GuildEnabled[guild.ID][eventName]; !ok && event.CanDisable {
			GuildEnabled[guild.ID][eventName] = true
		}
	}
	// Verifys guild has a section under GuildData
	if _, ok := GuildData[guild.ID]; !ok {
		GuildData[guild.ID] = make(map[string]string)
	}
	// Verifys every data value is listed
	for _, val := range DataValues {
		if _, ok := GuildData[guild.ID][val]; !ok {
			GuildData[guild.ID][val] = ""
		}
	}
}

// CheckEnabled checks if a command is enabled
func CheckEnabled(guild *discordgo.Guild, command string) bool {
	VerifyState(guild)
	if !Commands[command].CanDisable {
		return true
	}
	return GuildEnabled[guild.ID][command]
}

// EnableCommand sets commmand to enabled
func EnableCommand(guild *discordgo.Guild, command string) {
	VerifyState(guild)
	if !Commands[command].CanDisable {
		return
	}
	GuildEnabled[guild.ID][command] = true
	DumpEnabled()
}

// DisableCommand sets commmand to disabled
func DisableCommand(guild *discordgo.Guild, command string) { //Add more error detection
	VerifyState(guild)
	if !Commands[command].CanDisable {
		return
	}
	GuildEnabled[guild.ID][command] = false
	DumpEnabled()
}

// CheckData gets string data
func CheckData(guild *discordgo.Guild, data string) string {
	VerifyState(guild)
	return GuildData[guild.ID][data]
}

// SetData sets string data
func SetData(guild *discordgo.Guild, data string, value string) {
	VerifyState(guild)
	GuildData[guild.ID][data] = value
	DumpData()
}
