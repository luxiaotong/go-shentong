//go:build darwin

package db

import (
	"database/sql"
	"errors"
)

func OpenDB(dsn string) (*sql.DB, error) {
	return nil, errors.New("darwin cannot support [ShenTong] database.")
}
