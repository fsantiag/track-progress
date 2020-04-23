package repository

import (
	"github.com/fsantiag/track-progress/src/model"
	"github.com/gocql/gocql"
)

const insert = "INSERT INTO track (id, title, description, status) VALUES (?, ?, ?, ?)"

// TaskRepository represents a repository of task
type TaskRepository struct{}

// Save method used to save a new task
func (repository TaskRepository) Save(session *gocql.Session, task model.Task) (err error) {
	id, _ := gocql.RandomUUID()
	err = session.Query(insert, id, task.Title, task.Description, task.Status).Exec()
	defer session.Close()
	return
}
