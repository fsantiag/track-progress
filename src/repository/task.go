package repository

import (
	"log"

	"github.com/fsantiag/track-progress/src/model"
	"github.com/gocql/gocql"
)

const insertTask = "INSERT INTO tp.task (id, title, description, status) VALUES (?, ?, ?, ?)"

const selectTasks = "SELECT * FROM tp.task"

// TaskRepository represents a repository of task
type TaskRepository struct{}

// Save method used to save a new task
func (repository TaskRepository) Save(session SessionInterface, task model.Task) (err error) {
	id, _ := gocql.RandomUUID()
	err = session.Query(insertTask, id, task.Title, task.Description, task.Status).Exec()
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}

// GetAll returns all tasks in the table
func (repository TaskRepository) GetAll(session SessionInterface) (tasks []model.Task) {
	iter := session.Query(selectTasks).Iter()

	m := map[string]interface{}{}

	for iter.MapScan(m) {
		tasks = append(tasks, model.Task{
			ID:          m["id"].(gocql.UUID),
			Title:       m["title"].(string),
			Description: m["description"].(string),
			Status:      m["status"].(string),
		})
		m = map[string]interface{}{}
	}
	err := iter.Close()
	if err != nil {
		log.Fatal(err.Error())
	}
	return
}
