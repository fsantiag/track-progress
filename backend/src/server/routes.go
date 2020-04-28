package server

import (
	"net/http"

	"github.com/fsantiag/track-progress/backend/src/controller"
)

// Route a data structure for routes
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

var routes = []Route{
	{"Health", http.MethodGet, "/health", controller.Health},
}
