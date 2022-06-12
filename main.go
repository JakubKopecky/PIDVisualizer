package main

import (
	"database/sql"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	callApi()

	db, err := sql.Open("sqlite3", "./pid.db")
	db.Exec("create table memos(text, priority INTEGER);")
	db.Exec("insert into memos values('deliver project description', 10);")
	db.Exec("insert into memos values('lunch with Christine', 100);")
	checkErr(err)
	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
