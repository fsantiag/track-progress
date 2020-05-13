package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultValue = "http://localhost:4576"
)

func TestGetEnv_ShouldReturnDefaultValueWhenVariableNotFound(t *testing.T) {
	value := Getenv("SQS_HOST", defaultValue)

	assert.Equal(t, defaultValue, value)
}
