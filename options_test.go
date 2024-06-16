package errorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_optionsInitializer(t *testing.T) {

	jsonLog := JsonLog()
	options := []Options{}
	options = append(options, jsonLog)

	optionsInitializer(options)

	logConf := defaultConfigs[log]

	assert.Equal(t, logConf, "json", "two configs should be the same.")
}
