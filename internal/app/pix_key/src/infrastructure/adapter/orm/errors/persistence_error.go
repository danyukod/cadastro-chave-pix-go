package errors

import "errors"

type PersistenceError struct {
	table   string
	query   string
	message string
}

func NewPersistenceError(table, message, query string) *PersistenceError {
	return &PersistenceError{table: table, message: message, query: query}
}

func IsNoRecordError(err error) bool {
	var pe *PersistenceError
	if errors.As(err, &pe) {
		return pe.message == "record not found"
	}
	return false
}

func (b *PersistenceError) Tag() string {
	return "orm"
}

func (b *PersistenceError) Table() string {
	return b.table
}

func (b *PersistenceError) Query() string {
	return b.query
}

func (b *PersistenceError) Error() string {
	return b.message
}
