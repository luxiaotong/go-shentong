//go:build windows

package db

import (
	"database/sql"
	"errors"
)

func OpenDB(dsn string) (*sql.DB, error) {
	return nil, errors.New("windows cannot support [ShenTong] database.")
}
