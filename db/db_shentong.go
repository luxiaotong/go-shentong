//go:build shentong
// +build shentong

package db

import (
	"database/sql"
	"log"

	"github.com/luxiaotong/go-shentong/shentong"
)

func OpenDB(dsn string) (*sql.DB, error) {
	shentong.Init()
	db, err := sql.Open("aci", dsn)
	if err != nil {
		log.Panicf("open error: %v", err)
		return nil, err
	}
	return db, nil
}
