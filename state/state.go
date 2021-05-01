package state

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
)

// GuildEnabled store data about which commands/events are enabled
var GuildEnabled = make(map[string]map[string]bool)

// GuildData stores strings for each guilds
var GuildData = make(map[string]map[string]string)

// GuildLists stores []strings for each guilds
var GuildLists = make(map[string]map[string][]string)

// DataValues contains the data that is stored in GuildData
var DataValues = []string{"prefix", "joinmessage", "joinchannel", "leavemessage", "leavechannel"}

// ListValues contains the lists that are stored in GuildLists
var ListValues = []string{"blocked"}

// InitState loads GuildEnabled & GuildData from files if they exist, creates them if they don't
func InitState() {
	// Loads GuildEnabled
	file, err := os.Open("./GuildEnabled.json")
	if err != nil {
		os.Create("./GuildEnabled.json")
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildEnabled)
	// Loads GuildData
	file, err = os.Open("./GuildData.json")
	if err != nil {
		os.Create("./GuildData.json")
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildData)
	// Loads GuildLists
	file, err = os.Open("./GuildLists.json")
	if err != nil {
		os.Create("./GuildLists.json")
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildLists)
}

// DumpEnabled dumps the data for which commands are enabled to json
func DumpEnabled() {
	jsonData, err := json.Marshal(GuildEnabled)
	if err != nil {
		log.Println(err)
	}
	jsonFile, err := os.Create("./GuildEnabled.json")
	if err != nil {
		log.Println(err)
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

// DumpData dumps the strings to json
func DumpData() {
	jsonData, err := json.Marshal(GuildData)
	if err != nil {
		log.Println(err)
	}
	jsonFile, err := os.Create("./GuildData.json")
	if err != nil {
		log.Println(err)
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

// DumpLists dumps the []strings to json
func DumpLists() {
	jsonData, err := json.Marshal(GuildLists)
	if err != nil {
		log.Println(err)
	}
	jsonFile, err := os.Create("./GuildLists.json")
	if err != nil {
		log.Println(err)
	}
	_, err = jsonFile.Write(jsonData)
	if err != nil {
		log.Println(err)
	}
}

//VerifyState creates state if missing for selected Guild
func VerifyState(guildID string) {
	// Verifys guild has a section under GuildEnabled
	if _, ok := GuildEnabled[guildID]; !ok {
		GuildEnabled[guildID] = make(map[string]bool)
	}
	// Verifys each module that can be disabled is listed
	for modName := range manager.Modules {
		if _, ok := GuildEnabled[guildID][modName]; !ok {
			GuildEnabled[guildID][modName] = true
		}
	}
	// Verifys guild has a section under GuildData
	if _, ok := GuildData[guildID]; !ok {
		GuildData[guildID] = make(map[string]string)
	}
	// Verifys every data value is listed
	for _, val := range DataValues {
		if _, ok := GuildData[guildID][val]; !ok {
			GuildData[guildID][val] = ""
		}
	}
	// Verifys guild has a section under GuildLists
	if _, ok := GuildLists[guildID]; !ok {
		GuildLists[guildID] = make(map[string][]string)
	}
	// Verifys every data value is listed
	for _, val := range ListValues {
		if _, ok := GuildLists[guildID][val]; !ok {
			GuildLists[guildID][val] = []string{}
		}
	}
}

// CheckEnabled checks if a command is enabled
func CheckEnabled(guildID string, module string) bool {
	VerifyState(guildID)
	err, ok := manager.Modules[module]
	if ok {
		return GuildEnabled[guildID][module]
	} else if err != nil {
		log.Println(err)
	}
	return true
}

// EnableModule sets commmand to enabled
func EnableModule(guildID string, module string) {
	VerifyState(guildID)
	if !manager.IsValidModule(module) {
		return
	}
	GuildEnabled[guildID][module] = true
	DumpEnabled()
}

// DisableModule sets commmand to disabled
func DisableModule(guildID string, Module string) { //Add more error detection
	VerifyState(guildID)
	if !manager.IsValidModule(Module) {
		return
	}
	GuildEnabled[guildID][Module] = false
	DumpEnabled()
}

// CheckData gets string data
func CheckData(guildID string, data string) string {
	VerifyState(guildID)
	return GuildData[guildID][data]
}

// SetData sets string data
func SetData(guildID string, data string, value string) {
	VerifyState(guildID)
	GuildData[guildID][data] = value
	DumpData()
}

// AddToList adds data to the specified list
func AddToList(guildID string, data string, value string) {
	VerifyState(guildID)
	GuildLists[guildID][data] = append(GuildLists[guildID][data], value)
	DumpLists()
}

// RemoveFromList removes data from the specified list
func RemoveFromList(guildID string, data string, index int) {
	VerifyState(guildID)
	GuildLists[guildID][data] = functions.Remove(GuildLists[guildID][data], index)
	DumpLists()
}

// CheckList gets the list
func CheckList(guildID string, data string) []string {
	VerifyState(guildID)
	return GuildLists[guildID][data]
}
