package main

import (
	"fmt"
	"log"
	"strings"

	db2 "github.com/luxiaotong/go-shentong/db"
)

func main() {
	fmt.Println("test shentong db")

	dsn := "test:datassets@139.9.119.21:52003/OSRDB"
	db, err := db2.OpenDB(dsn)
	if err != nil {
		log.Println(err)
		return
	}
	defer db.Close()

	// q := "SELECT AREA_CODE, AREA_NAME FROM AREA LIMIT 0, 10"
	ff := []string{
		"TEST_BIGINT",
		"TEST_BINARY",
		"TEST_BIT",
		"TEST_BLOB",
		"TEST_BOOL",
		"TEST_CHAR",
		"TEST_CLOB",
		"TEST_DATE",
		"TEST_DECIMAL",
		"TEST_DOUBLE",
		"TEST_FLOAT",
		"TEST_INT",
		"TEST_NUMERIC",
		"TEST_REAL",
		"TEST_SERIAL",
		"TEST_SMALLINT",
		"TEST_TEXT",
		// "TEST_TIME",
		"TEST_TIMESTAMP",
		"TEST_TINYINT",
		"TEST_VARBINARY",
		"TEST_VARCHAR",
	}
	fields := strings.Join(ff, ",")
	q := fmt.Sprintf("SELECT %s FROM TEST_TYPE  LIMIT 0, 50", fields)
	// q := `SELECT "TEST_BIGINT","TEST_VARCHAR" FROM "TEST_TYPE"  LIMIT 0, 50`
	fmt.Println("query: ", q)
	rows, err := db.Query(q)
	if err != nil {
		log.Panicf("query error: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		// var code uint32
		// var name string
		// err = rows.Scan(&code, &name)
		vals := make([]interface{}, len(ff))
		for i := range vals {
			vals[i] = &vals[i]
		}
		err = rows.Scan(vals...)
		if err != nil {
			log.Panicf("scan error: %v\n", err)
		}
		// log.Printf("area code: %v, name: %v", code, name)
		log.Printf("vals: %v", vals)
	}
}
