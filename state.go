package main

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"github.com/bwmarrin/discordgo"
)

var GuildEnabled = make(map[string]map[string]bool)
var GuildData = make(map[string]map[string]string)
var DataValues = []string{"joinmessage", "joinchannel", "leavemessage", "leavechannel"}

func InitState() {
	file, err := os.Open("./GuildEnabled.json")
	if err != nil {
		os.Create("./GuildEnabled.json")
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildEnabled)
	file, err = os.Open("./GuildData.json")
	if err != nil {
		os.Create("./GuildEnabled.json")
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildData)
}

func DumpEnabled() {
	jsonData, _ := json.Marshal(GuildEnabled)
	jsonFile, _ := os.Create("./GuildEnabled.json")
	jsonFile.Write(jsonData)
}

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

func CheckEnabled(guild *discordgo.Guild, command string) bool {
	VerifyState(guild)
	return GuildEnabled[guild.ID][command]
}

func EnableCommand(guild *discordgo.Guild, command string) {
	VerifyState(guild)
	GuildEnabled[guild.ID][command] = true
	DumpEnabled()
}

func DisableCommand(guild *discordgo.Guild, command string) { //Add more error detection
	VerifyState(guild)
	GuildEnabled[guild.ID][command] = false
	DumpEnabled()
}

func CheckData(guild *discordgo.Guild, data string) string {
	VerifyState(guild)
	return GuildData[guild.ID][data]
}

func SetData(guild *discordgo.Guild, data string, value string) {
	VerifyState(guild)
	GuildData[guild.ID][data] = value
	DumpData()
}
