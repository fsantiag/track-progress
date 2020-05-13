package main

import (
	"net/http"

	"github.com/fsantiag/track-progress/back-for-front/src/controller/task"
)

func main() {
	http.HandleFunc("/", task.SaveTask)

	http.ListenAndServe(":8081", nil)
}
