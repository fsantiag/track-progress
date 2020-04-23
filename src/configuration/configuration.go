package configuration

import (
	"log"

	"github.com/gocql/gocql"
)

// NewConnection return a new connection to database
func NewConnection() *gocql.Session {
	cluster := gocql.NewCluster("localhost")
	cluster.Keyspace = "trackprogress"

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	return session
}
