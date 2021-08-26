package manager

import (
	"errors"
	"sort"
	"strings"
)

var Variables map[string]map[string]*Variable

type Variable struct {
	Name string

	ModuleName string

	DefaultValue interface{}
}

func RegisterEnable(moduleName string, defaultEnabled bool) {
	RegisterVariable(
		&Variable{
			Name:         "Enabled",
			ModuleName:   moduleName,
			DefaultValue: defaultEnabled,
		},
	)
}

func RegisterVariable(variable *Variable) {
	moduleID := strings.ToLower(variable.ModuleName)
	variableID := strings.ToLower(variable.Name)

	if Variables == nil {
		Variables = make(map[string]map[string]*Variable)
	}
	if Variables[moduleID] == nil {
		Variables[moduleID] = make(map[string]*Variable)
	}
	Variables[moduleID][variableID] = variable
}

func GetVariable(moduleID string, variableID string) (*Variable, error) {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	if _, ok := Variables[moduleID][variableID]; ok {
		return Variables[moduleID][variableID], nil
	}
	return nil, errors.New("invalid variable")
}

func IsValidVariable(moduleID string, variableID string) bool {
	moduleID = strings.ToLower(moduleID)
	variableID = strings.ToLower(variableID)

	_, exists := Variables[moduleID][variableID]
	return exists
}

func ListVariables(moduleID string) []string {
	moduleID = strings.ToLower(moduleID)

	keys := make([]string, 0, len(Variables[moduleID]))
	for k := range Variables[moduleID] {
		if k != "enabled" {
			keys = append(keys, k)
		}
	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	return keys
}
