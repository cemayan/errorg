package errorg

import (
	"fmt"
	"plugin"
)

var symName = "Group"

// Group represents interface that will be implemented by plugins
type Group interface {
	GetErrorMap() map[string]interface{}
}

// groupsInitializer adds a plugins  according to given plugins
func groupsInitializer() {

	for _, plg := range getConfig().Plugins {

		dirname, err := getModuleFolder()
		path := fmt.Sprintf("%s/%s", dirname, plg.Location)

		// Open opens a Go plugin.
		plug, err := plugin.Open(path)
		if err != nil {
			fmt.Errorf("an error occurred while opening plugin: %v", err)
			return
		}

		//Lookup searches for a symbol named symName in plugin p.
		sym, err := plug.Lookup(symName)
		if err != nil {
			fmt.Errorf("an error occurred while looking plugin: %v", err)
			return
		}

		var group Group
		group, ok := sym.(Group)
		if !ok {
			fmt.Errorf("%s", "unexpected type from module symbol")
			return
		}

		// returned map will be added as a new key to global errors map
		errorsMap[plg.Name] = group.GetErrorMap()
	}
}
