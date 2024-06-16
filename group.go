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
		// load module
		// 1. open the so file to load the symbols

		dirname, err := getModuleFolder()

		path := fmt.Sprintf("%s/%s", dirname, plg.Location)
		plug, err := plugin.Open(path)
		if err != nil {
			fmt.Errorf("an error occurred while opening plugin: %v", err)
			return
		}

		//
		sym, err := plug.Lookup(symName)
		if err != nil {
			fmt.Errorf("an error occurred while looking plugin: %v", err)
			return
		}

		// 3. Assert that loaded symbol is of a desired type
		// in this case interface type Greeter (defined above)
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
