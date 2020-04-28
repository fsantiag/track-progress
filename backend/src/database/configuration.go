package database

import (
	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/fsantiag/track-progress/backend/src/util"
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
		return nil, err
	}
	return repository.NewSession(session), nil
}
