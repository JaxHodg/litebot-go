package state

import (
	"strings"

	"cloud.google.com/go/firestore"

	"context"
	"errors"
	"log"

	"github.com/JaxHodg/litebot-go/manager"
)

var client *firestore.Client
var ctx context.Context

func createClient(ctx context.Context) *firestore.Client {
	// Sets your Google Cloud Platform project ID.
	projectID := "pure-feat-278023"

	client, err := firestore.NewClient(ctx, projectID)
	if err != nil {
		log.Print("Failed to create client: %v", err)
	}
	// Close client when done with
	// defer client.Close()
	return client
}

func InitDB() {
	ctx = context.Background()
	client = createClient(ctx)
}

func StopDB() {
	client.Close()
}

func GetEnabled(guildID string, moduleID string) (bool, error) {
	if !manager.IsValidModule(moduleID) {
		return false, errors.New("invalid module")
	}
	variable, err := manager.GetVariable(moduleID, "enabled")
	if err != nil {
		return false, errors.New("invalid module")
	}
	defaultValue, ok := variable.DefaultValue.(bool)
	if !ok {
		return false, errors.New("invalid module")
	}
	dsnap, err := client.Collection("guilds").Doc(guildID).Get(ctx)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}
	n, err := dsnap.DataAt(moduleID)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}
	enabledInterface, ok := n.(map[string]interface{})["enabled"]
	if !ok {
		return defaultValue, errors.New("firestore error")
	}

	enabled, ok := enabledInterface.(bool)
	if !ok {
		return defaultValue, errors.New("firestore error")
	}

	return enabled, nil
}

func EnableModule(guildID string, moduleID string) {
	if !manager.IsValidModule(moduleID) {
		return
	}
	_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
		moduleID: map[string]interface{}{
			"enabled": true,
		},
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func DisableModule(guildID string, moduleID string) {
	if !manager.IsValidModule(moduleID) {
		return
	}
	_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
		moduleID: map[string]interface{}{
			"enabled": false,
		},
	}, firestore.MergeAll)
	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func GetData(guildID string, moduleID string, variableID string) (string, error) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	variable, err := manager.GetVariable(moduleID, variableID)
	if err != nil {
		return "", errors.New("invalid module")
	}
	defaultValue, ok := variable.DefaultValue.(string)
	if !ok {
		return "", errors.New("can't get default value")
	}
	dsnap, err := client.Collection("guilds").Doc(guildID).Get(ctx)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}

	dataInterface, err := dsnap.DataAt(moduleID + "." + variableID)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}
	data, ok := dataInterface.(string)
	if !ok {
		return defaultValue, errors.New("firestore error")
	}
	return data, nil
}

func SetData(guildID string, moduleID string, variableID string, value string) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	if !manager.IsValidVariable(moduleID, variableID) {
		return
	}

	_, err := client.Collection("guilds").Doc(guildID).Set(ctx, map[string]interface{}{
		moduleID: map[string]interface{}{
			variableID: value,
		},
	}, firestore.MergeAll)

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func GetList(guildID string, moduleID string, variableID string) ([]string, error) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	if !manager.IsValidModule(moduleID) {
		return []string{}, errors.New("invalid module")
	}
	variable, err := manager.GetVariable(moduleID, variableID)
	if err != nil {
		return []string{}, errors.New("invalid module")
	}
	defaultValue, ok := variable.DefaultValue.([]string)
	if !ok {
		return []string{}, errors.New("invalid module")
	}
	dsnap, err := client.Collection("guilds").Doc(guildID).Get(ctx)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}
	dataInterface, err := dsnap.DataAt(moduleID + "." + variableID)
	if err != nil {
		return defaultValue, errors.New("firestore error")
	}
	dataInterface2, ok := dataInterface.([]interface{})
	if !ok {
		return defaultValue, errors.New("firestore error")
	}
	data := make([]string, len(dataInterface2))
	for i := range dataInterface2 {
		data[i] = dataInterface2[i].(string)
	}
	return data, nil
}

func AddToList(guildID string, moduleID string, variableID string, value string) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	if !manager.IsValidModule(moduleID) {
		return
	}

	_, err := client.Collection("guilds").Doc(guildID).Update(ctx, []firestore.Update{
		{Path: moduleID + "." + variableID, Value: firestore.ArrayUnion(value)},
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func RemoveToList(guildID string, moduleID string, variableID string, value string) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	if !manager.IsValidModule(moduleID) {
		return
	}

	_, err := client.Collection("guilds").Doc(guildID).Update(ctx, []firestore.Update{
		{Path: moduleID + "." + variableID, Value: firestore.ArrayRemove(value)},
	})

	if err != nil {
		// Handle any errors in an appropriate way, such as returning them.
		log.Printf("An error has occurred: %s", err)
	}
}

func RemoveGuild(guildID string) {
	_, err := client.Collection("guilds").Doc(guildID).Delete(ctx)
	if err != nil {
		log.Printf("An error has occurred: %s", err)
	}
}
