package configuration

// Migrate creates the KeySpace and Table for the project
func Migrate() {
	session, _ := NewSession()
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

	session.Close()
}
