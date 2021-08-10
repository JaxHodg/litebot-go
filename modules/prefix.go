package modules

import "github.com/JaxHodg/litebot-go/manager"

func init() {
	manager.RegisterVariable(
		&manager.Variable{
			Name:         "Prefix",
			ModuleName:   "Prefix",
			DefaultValue: "!",
		},
	)
	manager.RegisterModule(
		&manager.Module{
			Name: "Prefix",

			Description: "",
		},
	)
}
