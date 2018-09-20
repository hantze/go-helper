package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// Queryable ...
type Queryable interface {
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRowx(query string, args ...interface{}) *sqlx.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
	GetDB() *sqlx.DB
}
