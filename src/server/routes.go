package server

import (
	"net/http"

	"github.com/fsantiag/track-progress/src/controller"
)

// Route a data structure for routes
type Route struct {
	Name        string
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

var routes = [...]Route{
	{"Status", http.MethodGet, "/health", controller.Health},
}
