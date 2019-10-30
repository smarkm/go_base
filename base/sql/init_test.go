package sql_test

import (
	"database/sql"
	"log"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", "root:ossera@tcp(172.16.2.2)/test")
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		log.Println(err)
		return
	}
}
