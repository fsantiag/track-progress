package repository

import "github.com/gocql/gocql"

// SessionInterface allows gomock mock of gocql.Session
type SessionInterface interface {
	Query(string, ...interface{}) QueryInterface
	Close()
}

// QueryInterface allows gomock mock of gocql.Query
type QueryInterface interface {
	Exec() error
	Iter() IterInterface
}

// IterInterface allows gomock mock of gocql.Iter
type IterInterface interface {
	MapScan(map[string]interface{}) bool
	Close() error
}

// Session is a wrapper for a session for mockability.
type Session struct {
	session *gocql.Session
}

// Query is a wrapper for a query for mockability.
type Query struct {
	query *gocql.Query
}

// Iter is a wrapper for an iter for mockability.
type Iter struct {
	iter *gocql.Iter
}

// NewSession instantiates a new Session
func NewSession(session *gocql.Session) SessionInterface {
	return &Session{
		session,
	}
}

// NewQuery instantiates a new Query
func NewQuery(query *gocql.Query) QueryInterface {
	return &Query{
		query,
	}
}

// NewIter instantiates a new Iter
func NewIter(iter *gocql.Iter) IterInterface {
	return &Iter{
		iter,
	}
}

// Query wraps the session's query method
func (s *Session) Query(stmt string, values ...interface{}) QueryInterface {
	return NewQuery(s.session.Query(stmt, values...))
}

//Close sssion
func (s *Session) Close() {
	s.session.Close()
}

// Exec wraps the query's Exec method
func (q *Query) Exec() error {
	return q.query.Exec()
}

// Iter wraps the query's Iter method
func (q *Query) Iter() IterInterface {
	return NewIter(q.query.Iter())
}

// Scan is a wrapper for the iter's Scan method
func (i *Iter) Scan(dest ...interface{}) bool {
	return i.iter.Scan(dest...)
}

//MapScan does stuff
func (i *Iter) MapScan(m map[string]interface{}) bool {
	return i.iter.MapScan(m)
}

//Close stuff
func (i *Iter) Close() error {
	return i.iter.Close()
}
