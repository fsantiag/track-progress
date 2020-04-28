package model

import "github.com/gocql/gocql"

// Task represents data about each task
type Task struct {
	ID          gocql.UUID
	Title       string
	Description string
	Status      string
}
