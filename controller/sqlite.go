package controller

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sql.DB
}

func NewDbController(file string) *Sqlite {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}

	cleanDB(db)

	return &Sqlite{db: db}
}

func cleanDB(db *sql.DB) {
	rows, err := db.Query("SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		log.Print("Droping table: " + name)
		dropTableByName(db, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}

func dropTableByName(db *sql.DB, name string) {
	sql := fmt.Sprintf("DROP TABLE %s;", name)
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func (sqlite *Sqlite) Exec(sql string) (sql.Result, error) {
	return sqlite.db.Exec(sql)
}
