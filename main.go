package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	callApi()

	db, err := sql.Open("sqlite3", "./names.db")
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
