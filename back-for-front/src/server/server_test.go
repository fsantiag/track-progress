package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingRoute(t *testing.T) {
	router := InitRouter()

	t.Run("Should have mapped Health route", func(t *testing.T) {
		assert.NotNil(t, router.Get("Health"))
	})

	t.Run("Should have mapped SaveTask route", func(t *testing.T) {
		assert.NotNil(t, router.Get("SaveTask"))
	})
}
