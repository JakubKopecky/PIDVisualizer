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
	for _, item := range getTables(db) {
		dropTableByName(db, item)

	}
	initTables(db)
}

func getTables(db *sql.DB) []string {
	rows, err := db.Query("SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var tables []string
	var name string
	for rows.Next() {
		err := rows.Scan(&name)
		if err != nil {
			log.Fatal(err)
		}
		tables = append(tables, name)
	}

	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return tables
}

func dropTableByName(db *sql.DB, name string) {
	log.Print("Droping table: " + name)

	sql := fmt.Sprintf("DROP TABLE %s;", name)
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func initTables(db *sql.DB) {
	log.Print("Creating table: TBLTRIPS")
	_, err := db.Exec("CREATE TABLE `TBLTRIPS` (`ID` INT NOT NULL AUTO_INCREMENT, `TRIP_ID` TEXT NOT NULL, PRIMARY KEY (`ID`));")
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Creating table: TBLTRIPUPDATES")
	_, err = db.Exec("CREATE TABLE `TBLTRIPUPDATES` (`ID` INT NOT NULL AUTO_INCREMENT, `TRIP_ID` TEXT NOT NULL, `DATA` TEXT NOT NULL, PRIMARY KEY (`ID`));")
	if err != nil {
		log.Fatal(err)
	}
}

func (sqlite *Sqlite) Exec(sql string) (sql.Result, error) {
	log.Print("Executing: " + sql)
	return sqlite.db.Exec(sql)
}

func (sqlite *Sqlite) Close() {
	sqlite.db.Close()
}
