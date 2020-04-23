package main

import (
	"log"
	"net/http"

	"github.com/fsantiag/track-progress/src/configuration"
	"github.com/fsantiag/track-progress/src/server"
)

func main() {
	configuration.Migrate()

	// session := configuration.NewSession()
	// defer session.Close()

	// repo := repository.TaskRepository{}
	// err := repo.Save(session, model.Task{Title: "1", Description: "1", Status: "1"})
	// if err != nil {
	// 	log.Fatal(err.Error())
	// }

	// tasks := repo.GetAll(session)
	// fmt.Print(tasks)

	s := server.InitRouter()
	log.Println("Server started...")
	log.Fatal(http.ListenAndServe(":8080", s))
}
