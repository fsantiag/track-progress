package health

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthController(t *testing.T) {
	request, err := http.NewRequest("GET", "/health", nil)
	assert.Nil(t, err)

	recorder := httptest.NewRecorder()
	handler := http.HandlerFunc(Health)

	handler.ServeHTTP(recorder, request)

	assert.Equal(t, http.StatusOK, recorder.Code)
	assert.Equal(t, "{\"status\":\"Up!\"}", recorder.Body.String())
}
