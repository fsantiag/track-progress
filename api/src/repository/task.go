package repository

import (
	"log"

	"github.com/fsantiag/track-progress/api/src/model"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

const insertTask = "INSERT INTO tp.task (id, title, description, status) VALUES (?, ?, ?, ?)"
const selectTasks = "SELECT * FROM tp.task"
const updateTask = "UPDATE tp.task SET title = ?, description = ?, status = ? WHERE id = ?"
const deleteTask = "DELETE FROM tp.task WHERE id = ?"

var (
	logger  *logrus.Logger
	session SessionInterface
)

// Repository representes all method of tasks
type Repository interface {
	Save(task model.Task) (err error)
	GetAll() (tasks []model.Task)
	Update(task model.Task) (err error)
	Delete(id gocql.UUID) (err error)
}

type taskRepository struct{}

// NewTaskRepository create a new instance of Repository with your arguments
func NewTaskRepository(loggerLogrus *logrus.Logger, sessionInterface SessionInterface) Repository {
	logger = loggerLogrus
	session = sessionInterface
	return taskRepository{}
}

// Save method used to save a new task
func (repository taskRepository) Save(task model.Task) (err error) {
	id, _ := gocql.RandomUUID()
	err = session.Query(insertTask, id, task.Title, task.Description, task.Status).Exec()
	return
}

// GetAll returns all tasks in the table
func (repository taskRepository) GetAll() (tasks []model.Task) {
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

// Update changes task values
func (repository taskRepository) Update(task model.Task) (err error) {
	err = session.Query(updateTask, task.Title, task.Description, task.Status, task.ID).Exec()
	return
}

// Delete remove task on database by id
func (repository taskRepository) Delete(id gocql.UUID) (err error) {
	err = session.Query(deleteTask, id).Exec()
	return
}
