package database

import "github.com/fsantiag/track-progress/api/src/repository"

// Migrate creates the KeySpace and Table for the project
func Migrate(session repository.SessionInterface) {
	session.Query(`
	CREATE KEYSPACE IF NOT EXISTS tp
	WITH REPLICATION = {
	'class' : 'SimpleStrategy',
	'replication_factor' : 1 }`).Exec()

	session.Query(`
	CREATE TABLE IF NOT EXISTS tp.task (
		id UUID PRIMARY KEY,
		title text,
		description text,
		status text)`).Exec()
}
