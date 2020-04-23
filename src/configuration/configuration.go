package configuration

import (
	"log"

	"github.com/gocql/gocql"
)

// NewSession return a new session to database
func NewSession() *gocql.Session {
	cluster := gocql.NewCluster("localhost")

	cluster.Authenticator = gocql.PasswordAuthenticator{
		Username: "cassandra",
		Password: "cassandra",
	}

	session, err := cluster.CreateSession()
	if err != nil {
		log.Fatal(err.Error())
	}
	return session
}
