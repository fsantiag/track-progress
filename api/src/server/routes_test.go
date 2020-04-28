package server

import (
	"net/http"
	"testing"

	"github.com/fsantiag/track-progress/api/src/controller"
	"github.com/fsantiag/track-progress/api/src/util"
	"github.com/stretchr/testify/assert"
)

func TestValidRoutes(t *testing.T) {
	testTable := []Route{
		{"Health", http.MethodGet, "/health", controller.Health},
	}

	for _, route := range routes {
		for _, testData := range testTable {
			expectedFunction := util.GetFunctionName(testData.HandlerFunc)
			actualFunction := util.GetFunctionName(route.HandlerFunc)

			assert.Equal(t, expectedFunction, actualFunction)
			assert.Equal(t, testData.Name, route.Name)
			assert.Equal(t, testData.Method, route.Method)
			assert.Equal(t, testData.Path, route.Path)
		}
	}
}
