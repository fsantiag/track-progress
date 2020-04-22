package server

import (
	"github.com/gorilla/mux"
)

// InitRoutes Map request into the defined routes
func InitRoutes() *mux.Router {
	route := mux.NewRouter()
	for _, r := range routes {
		route.Methods(r.Method).Name(r.Name).Path(r.Path).Handler(r.HandlerFunc)
	}
	return route
}
