package cassandra

import (
	"errors"
	"time"

	"github.com/fsantiag/track-progress/backend/src/repository"
	"github.com/fsantiag/track-progress/backend/src/util"
	"github.com/gocql/gocql"
	"github.com/sirupsen/logrus"
)

// NewSession return a new session to database
func NewSession(logger *logrus.Logger) (repository.SessionInterface, error) {
	cluster := gocql.NewCluster(util.Getenv("CASSANDRA_HOST", "localhost"))

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: util.Getenv("CASSANDRA_USERNAME", "cassandra"),
		Password: util.Getenv("CASSANDRA_PASSWORD", "cassandra"),
	}

	maxAttempts := 3
	var err error

	for currentAttempt := 0; currentAttempt < maxAttempts; currentAttempt++ {
		var session *gocql.Session
		session, err = cluster.CreateSession()
		if session != nil {
			return repository.NewSession(session), nil
		}
		if currentAttempt > maxAttempts {
			err = errors.New("Max attemps reached trying to connect to Cassandra")
		}
		logger.Warn("Error connecting to Cassandra, retrying...: ", err.Error())
		time.Sleep(5 * time.Second)
	}
	return nil, err
}
