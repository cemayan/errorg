package errorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_configInitializer(t *testing.T) {
	configInitializer()
	assert.Equal(t, len(config.Plugins), 2, "two lengths should be the same.")
}
