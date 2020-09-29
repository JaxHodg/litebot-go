package manager

import (
	"strings"
)

var Modules map[string]*Module

type Module struct {
	Name string

	Description string
}

func RegisterModule(module *Module) {
	if Modules == nil {
		Modules = make(map[string]*Module)
	}
	Modules[strings.ToLower(module.Name)] = module
}

func GetModule(module string) *Module {
	if _, ok := Modules[module]; ok {
		return Modules[module]
	}
	return nil
}

func IsValidModule(module string) bool {
	if _, ok := Modules[module]; ok {
		return true
	}
	return false
}
