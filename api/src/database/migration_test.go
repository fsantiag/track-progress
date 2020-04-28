package database

import (
	"testing"

	"github.com/fsantiag/track-progress/api/src/repository"
	"github.com/stretchr/testify/assert"
)

var parameters = make([]string, 0, 2)

type stubSession struct {
	repository.SessionInterface
}

type stubQuery struct {
	repository.QueryInterface
}

func (s stubSession) Query(query string, variables ...interface{}) repository.QueryInterface {
	parameters = append(parameters, query)
	return stubQuery{}
}

func (q stubQuery) Exec() error {
	return nil
}

func TestMigration(t *testing.T) {
	session := stubSession{}
	Migrate(session)

	assert.Equal(t, `
	CREATE KEYSPACE IF NOT EXISTS tp
	WITH REPLICATION = {
	'class' : 'SimpleStrategy',
	'replication_factor' : 1 }`, parameters[0])

	assert.Equal(t, `
	CREATE TABLE IF NOT EXISTS tp.task (
		id UUID PRIMARY KEY,
		title text,
		description text,
		status text)`, parameters[1])
}
