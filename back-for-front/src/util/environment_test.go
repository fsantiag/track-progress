package util

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	defaultValue = "http://localhost:4576"
)

func TestAddress_ShouldReturnAddressOfEachEnvironment(t *testing.T) {
	t.Run("Should return address to dev environment", func(t *testing.T) {
		assert.Equal(t, "localhost:8081", Address["dev"])
	})

	t.Run("Should return address to prod environment", func(t *testing.T) {
		assert.Equal(t, ":8081", Address["prod"])
	})
}

func TestGetEnv_ShouldReturnDefaultValueWhenVariableNotFound(t *testing.T) {
	value := Getenv("SQS_HOST", defaultValue)

	assert.Equal(t, defaultValue, value)
}
