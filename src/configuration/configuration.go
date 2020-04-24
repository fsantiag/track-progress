package configuration

import (
	"fmt"

	"github.com/fsantiag/track-progress/src/repository"
	"github.com/gocql/gocql"
)

// NewSession return a new session to database
func NewSession() (repository.SessionInterface, error) {
	cluster := gocql.NewCluster("localhost")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	session, err := cluster.CreateSession()
	if err != nil {
		return nil, fmt.Errorf("error")
	}
	return repository.NewSession(session), nil
}
