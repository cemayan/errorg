package errorg

import (
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"testing"
)

func Test_groupsInitializer(t *testing.T) {
	errorsMap = map[string]map[string]interface{}{}
	yamlFile := readYaml("config", "configs")
	_ = yaml.Unmarshal(yamlFile, &config)

	groupsInitializer()
	assert.Equal(t, len(errorsMap), 2, "two lengths should be the same.")
}

func Test_groupsInitializer_WrongLocation(t *testing.T) {
	errorsMap = map[string]map[string]interface{}{}
	yamlFile := readYaml("config", "configs")
	_ = yaml.Unmarshal(yamlFile, &config)

	if len(config.Plugins) > 0 {
		config.Plugins[0].Location = ""
	}

	groupsInitializer()
	assert.Equal(t, len(errorsMap), 0, "two lengths should be the same.")
}

func Test_groupsInitializer_WrongGroup(t *testing.T) {
	errorsMap = map[string]map[string]interface{}{}
	yamlFile := readYaml("config", "configs")
	_ = yaml.Unmarshal(yamlFile, &config)

	symName = "wrongGroup"

	groupsInitializer()
	assert.Equal(t, len(errorsMap), 0, "two lengths should be the same.")
}

func Test_groupsInitializer_WrongCasting(t *testing.T) {
	errorsMap = map[string]map[string]interface{}{}
	yamlFile := readYaml("config", "configs")
	_ = yaml.Unmarshal(yamlFile, &config)

	symName = "wrongGroup"

	groupsInitializer()
	assert.Equal(t, len(errorsMap), 0, "two lengths should be the same.")
}
