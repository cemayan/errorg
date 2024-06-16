package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestHttError_GetErrorMap(t *testing.T) {
	httpErr := httpError{}
	errorMap := httpErr.GetErrorMap()
	assert.NotNil(t, errorMap["400"])
}
