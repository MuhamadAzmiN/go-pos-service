package iface

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

type (
	ISqlx interface {
		Exec(query string, args ...interface{}) (sql.Result, error)
		NamedExec(query string, arg interface{}) (sql.Result, error)
		Queryx(query string, args ...interface{}) (*sqlx.Rows, error)
		QueryRowx(query string, args ...interface{}) *sqlx.Row
		NamedQuery(query string, arg interface{}) (*sqlx.Rows, error)
		Get(dest interface{}, query string, args ...interface{}) error
		Select(dest interface{}, query string, args ...interface{}) error
		Rebind(query string) string
	}
)
