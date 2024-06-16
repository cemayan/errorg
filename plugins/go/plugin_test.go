package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestStdGoError_GetErrorMap(t *testing.T) {
	stdGoError := stdGoError{}
	errorMap := stdGoError.GetErrorMap()

	assert.NotNil(t, errorMap["100"])
}
