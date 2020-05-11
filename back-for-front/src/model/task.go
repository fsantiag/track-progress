package model

import "github.com/gocql/gocql"

// Task represents data about each task
type Task struct {
	ID          gocql.UUID `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Status      string     `json:"status"`
}
