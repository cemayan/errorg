package errorg

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"os/exec"
)

// config is used to reach config  whenever what you want
var config GeneralConfig

var (
	ConfigName  = "config"
	ConfigPaths = "configs"
)

// directoryInitializer executes make init command
func directoryInitializer() {
	cmd := exec.Command("make", "init")
	cmd.Dir, _ = getModuleFolder()
	err := cmd.Start()
	if err != nil {
		fmt.Printf("make init command err: %v", err)
		return
	}
}

// configInitializer unmarshalls the yaml
func configInitializer() {
	yamlFile := readYaml(ConfigName, ConfigPaths)
	err := yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		fmt.Printf("Unmarshal: %v", err)
	}
}

// getConfig returns the config
func getConfig() GeneralConfig {
	return config
}
