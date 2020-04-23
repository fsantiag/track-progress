package repository

import (
	"log"

	"github.com/fsantiag/track-progress/src/model"
	"github.com/gocql/gocql"
)

const insert = "INSERT INTO tp.task (id, title, description, status) VALUES (?, ?, ?, ?)"

// TaskRepository represents a repository of task
type TaskRepository struct{}

// Save method used to save a new task
func (repository TaskRepository) Save(session *gocql.Session, task model.Task) (err error) {
	id, _ := gocql.RandomUUID()
	err = session.Query(insert, id, task.Title, task.Description, task.Status).Exec()
	// defer session.Close()
	return
}

// GetAll returns all tasks in the table
func (repository TaskRepository) GetAll(session *gocql.Session) (tasks []model.Task) {
	iter := session.Query("SELECT * FROM tp.task").Iter()
	var id gocql.UUID
	var title, description, status string
	for iter.Scan(&id, &title, &description, &status) {
		tasks = append(tasks, model.Task{ID: id, Title: title, Description: description, Status: status})
	}
	err := iter.Close()
	// defer session.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
