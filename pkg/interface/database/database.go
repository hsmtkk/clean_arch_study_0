package database

type SQLHandler interface {
	Execute(statement string, args ...interface{}) (Result, error)
	Query(statement string, args ...interface{}) (Row, error)
}

type Result interface {
	LastInsertID() (int64, error)
	RowsAffected() (int64, error)
}

type Row interface {
	Scan(dest ...interface{}) error
	Next() bool
	Close() error
}
