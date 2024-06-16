package errorg

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	Init()
	assert.Greater(t, len(config.Plugins), 0)
	assert.Greater(t, len(errorsMap), 0)
}

func TestInit_WithOptions(t *testing.T) {
	Init(JsonLog())
	defaultConfigs[log] = logTypeMap[jsonType]
}

func Test_prepare(t *testing.T) {

	msg400 := "The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing)."

	Init(TextLog())
	e := errorg{Group: "http", Code: "400"}
	prepare := e.prepare()
	assert.Equal(t, msg400, prepare)
}

func Test_prepare_Json(t *testing.T) {

	msg400 := "{\"msg\":\"The server cannot or will not process the request due to something that is perceived to be a client error (e.g., malformed request syntax, invalid request message framing, or deceptive request routing).\"}"

	Init(JsonLog())
	e := errorg{Group: "http", Code: "400"}
	prepare := e.prepare()

	assert.Equal(t, msg400, prepare)
}
