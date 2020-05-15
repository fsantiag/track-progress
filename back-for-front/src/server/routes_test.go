package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const (
	firstRoute  = 0
	secondRoute = 1
)

func TestRoutes(t *testing.T) {
	t.Run("Should have Health route", func(t *testing.T) {
		route := routes[firstRoute]
		assert.Equal(t, "Health", route.Name)
		assert.Equal(t, "GET", route.Method)
		assert.Equal(t, "/health", route.Path)
	})

	t.Run("Should have SaveTask route", func(t *testing.T) {
		route := routes[secondRoute]
		assert.Equal(t, "SaveTask", route.Name)
		assert.Equal(t, "POST", route.Method)
		assert.Equal(t, "/task/save", route.Path)
	})
}
