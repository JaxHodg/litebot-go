package manager

import (
	"errors"
	"sort"
	"strings"
)

var Modules map[string]*Module

//TODO: decide between ID or Name
type Module struct {
	Name string

	Description string
}

func RegisterModule(module *Module) {
	moduleID := strings.ToLower(module.Name)

	if Modules == nil {
		Modules = make(map[string]*Module)
	}
	Modules[moduleID] = module
}

func GetModule(moduleID string) (*Module, error) {
	moduleID = strings.ToLower(moduleID)

	if _, ok := Modules[moduleID]; ok {
		return Modules[moduleID], nil
	}
	return nil, errors.New("invalid module")
}

func IsValidModule(moduleID string) bool {
	moduleID = strings.ToLower(moduleID)

	_, exists := Modules[moduleID]
	return exists
}

func ListModules() []string {
	keys := make([]string, 0, len(Modules))
	for k := range Modules {

		keys = append(keys, k)

	}
	sort.Slice(keys, func(i, j int) bool {
		return len(keys[i]) < len(keys[j])
	})
	return keys
}
