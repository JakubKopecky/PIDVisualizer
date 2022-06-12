package main

import (
	"database/sql"
	"log"
)

func main() {
	callApi()

	db, err := sql.Open("sqlite", "./names.db")
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
