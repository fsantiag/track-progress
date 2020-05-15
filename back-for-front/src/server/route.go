package server

import (
	"net/http"

	"github.com/fsantiag/track-progress/back-for-front/src/controller/health"
	"github.com/fsantiag/track-progress/back-for-front/src/controller/task"
)

// Route a data structure for routes
type Route struct {
	Name, Method, Path string
	HandlerFunc        http.HandlerFunc
}

var routes = []Route{
	{Name: "Health", Method: "GET", Path: "/health", HandlerFunc: health.Health},
	{Name: "SaveTask", Method: "POST", Path: "/task/save", HandlerFunc: task.SaveTask},
}
