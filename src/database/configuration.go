package database

import (
	"fmt"

	"github.com/fsantiag/track-progress/src/repository"
	"github.com/fsantiag/track-progress/src/util"
	"github.com/gocql/gocql"
)

// NewSession return a new session to database
func NewSession() (repository.SessionInterface, error) {
	cluster := gocql.NewCluster(util.Getenv("CASSANDRA_HOST", "localhost"))

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: util.Getenv("CASSANDRA_USERNAME", "cassandra"),
		Password: util.Getenv("CASSANDRA_PASSWORD", "cassandra"),
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return repository.NewSession(session), nil
}
