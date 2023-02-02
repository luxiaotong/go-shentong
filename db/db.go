//go:build !shentong
// +build !shentong

package db

import (
	"database/sql"
	"errors"
)

func OpenDB(dsn string) (*sql.DB, error) {
	return nil, errors.New("cannot support [ShenTong] database.")
}
