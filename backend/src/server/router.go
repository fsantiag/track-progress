package server

import (
	"net/http"

	"github.com/fsantiag/track-progress/backend/src/controller"
	"github.com/gorilla/mux"
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

// InitRouter maps the defined routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	for _, r := range routes {
		router.Methods(r.Method).Name(r.Name).Path(r.Path).Handler(r.HandlerFunc)
	}
	return router
}
