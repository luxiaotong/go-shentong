package main

import (
	"fmt"
	"log"

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

	q := "SELECT AREA_CODE, AREA_NAME FROM AREA LIMIT 0, 10"
	fmt.Println("query: ", q)
	rows, err := db.Query(q)
	if err != nil {
		log.Panicf("query error: %v\n", err)
	}
	defer rows.Close()

	for rows.Next() {
		var code uint32
		var name string
		err = rows.Scan(&code, &name)
		if err != nil {
			log.Panicf("scan error: %v\n", err)
		}
		log.Printf("area code: %v, name: %v", code, name)
	}
}
