package server

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMappingRoute(t *testing.T) {
	testTable := []struct {
		route string
	}{
		{"Health"},
	}

	router := InitRouter()
	for _, testData := range testTable {
		assert.NotNil(t, router.Get(testData.route))
	}
}
