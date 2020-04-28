package main

import (
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/fsantiag/track-progress/backend/src/controller"
	"github.com/stretchr/testify/assert"
)

func TestMain(t *testing.T) {
	testTable := []struct {
		Method      string
		HandlerFunc http.HandlerFunc
	}{
		{"GET", controller.Health},
	}

	for _, testData := range testTable {
		httpTest := httptest.NewServer(http.HandlerFunc(testData.HandlerFunc))
		defer httpTest.Close()

		client := &http.Client{}
		request, _ := http.NewRequest(testData.Method, httpTest.URL, nil)
		response, err := client.Do(request)
		assert.Nil(t, err, "Error performing request.")
		defer response.Body.Close()

		responseBody, _ := ioutil.ReadAll(response.Body)
		assert.NotEmpty(t, responseBody)
		assert.Equal(t, http.StatusOK, response.StatusCode)
	}
}
