package server

import (
	"github.com/gorilla/mux"
)

// InitRouter maps the defined routes
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	spa := SpaHandler{staticPath: "static", indexPath: "index.html"}
	router.PathPrefix("/").Handler(spa)
	for _, r := range routes {
		router.Methods(r.Method).Name(r.Name).Path(r.Path).Handler(r.HandlerFunc)
	}
	return router
}
