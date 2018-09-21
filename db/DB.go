package db

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

// DB ...
type DB struct {
	Db *sqlx.DB
}

// Query ...
func (d *DB) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.Db.Queryx(query, args)
}

// Queryx ...
func (d *DB) Queryx(query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.Db.Queryx(query, args)
}

// QueryRowx ...
func (d *DB) QueryRowx(query string, args ...interface{}) *sqlx.Row {
	return d.Db.QueryRowx(query, args)
}

func (d *DB) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.Db.Exec(query, args)
}

// GetDB ...
func (d *DB) GetDB() *sqlx.DB {
	return d.Db
}

// NewDB ...
func NewDB(db *sqlx.DB) Queryable {
	return &DB{Db: db}
}
