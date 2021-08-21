package state

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"

	"cloud.google.com/go/firestore"
	"github.com/JaxHodg/litebot-go/functions"
	"github.com/JaxHodg/litebot-go/manager"
)

func MigrateDB() {
	time.Sleep(5 * time.Second)
	// GuildEnabled store data about which commands/events are enabled
	var GuildEnabled = make(map[string]map[string]bool)

	// GuildData stores strings for each guilds
	var GuildData = make(map[string]map[string]string)

	// GuildLists stores []strings for each guilds
	var GuildLists = make(map[string]map[string][]string)

	// Loads GuildEnabled
	file, err := os.Open("./GuildEnabled.json")
	if err != nil {
		return
	}
	byteValue, _ := ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildEnabled)
	// Loads GuildData
	file, err = os.Open("./GuildData.json")
	if err != nil {
		return
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildData)
	// Loads GuildLists
	file, err = os.Open("./GuildLists.json")
	if err != nil {
		return
	}
	byteValue, _ = ioutil.ReadAll(file)
	json.Unmarshal(byteValue, &GuildLists)

	// Uploads GuildEnabled
	for guild, enabled := range GuildEnabled {
		guildID := strings.ToLower(guild)

		for command, value := range enabled {
			moduleID := strings.ToLower(command)

			switch moduleID {
			case "blockmessage":
				moduleID = "blockterm"
			}
			fmt.Println(guildID + " " + moduleID + " " + strconv.FormatBool(value))
			if !value && manager.IsValidVariable(moduleID, "enabled") {
				_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
					moduleID: map[string]interface{}{
						"enabled": value,
					},
				}, firestore.MergeAll)
				if err != nil {
					fmt.Print("Error updating db")
				}
			}
		}
	}
	// Uploads GuildData
	for guild, enabled := range GuildData {
		guildID := strings.ToLower(guild)

		for variable, value := range enabled {
			var moduleID, variableID string
			switch variable {
			case "joinmessage":
				moduleID = "joinmessage"
				variableID = "joinmessage"
			case "leavemessage":
				moduleID = "leavemessage"
				variableID = "leavemessage"
			case "leavechannel":
				moduleID = "leavemessage"
				variableID = "leavechannel"
			case "joinchannel":
				moduleID = "joinmessage"
				variableID = "joinchannel"
			case "prefix":
				moduleID = "prefix"
				variableID = "prefix"
			}
			fmt.Println(guildID + " " + moduleID + " " + value)
			if value != "" && manager.IsValidVariable(moduleID, variableID) {
				_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
					moduleID: map[string]interface{}{
						variableID: value,
					},
				}, firestore.MergeAll)
				if err != nil {
					fmt.Print("Error updating db")
				}
			}
		}
	}
	// Uploads GuildBlocked
	for guild, enabled := range GuildLists {
		guildID := strings.ToLower(guild)

		for variable, value := range enabled {
			var moduleID, variableID string
			switch variable {
			case "blocked":
				moduleID = "blockterm"
				variableID = "blockedterms"
			}
			for i, j := range value {
				term := j
				term = strings.TrimSpace(term)
				term = functions.NormaliseString(term)
				term = strings.ToLower(term)
				value[i] = term
			}

			fmt.Println(guildID + " " + moduleID)
			if len(value) > 0 && manager.IsValidVariable(moduleID, variableID) {
				_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
					moduleID: map[string]interface{}{
						variableID: value,
					},
				}, firestore.MergeAll)
				if err != nil {
					fmt.Print("Error updating db")
				}
			}
		}
	}
}
