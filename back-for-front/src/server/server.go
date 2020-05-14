package server

import "github.com/gorilla/mux"

// InitRouter maps the defined routes
func InitRouter() *mux.Router {
	muxRouter := mux.NewRouter()

	for _, route := range routes {
		muxRouter.Name(route.Name).Methods(route.Method).Path(route.Path).Handler(route.HandlerFunc)
	}

	return muxRouter
}
