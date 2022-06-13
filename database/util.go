package database

import (
	"fmt"
	"log"

	"github.com/JakubKopecky/PIDVisualizer/model"
	"github.com/jmoiron/sqlx"
)

var getTablesSql = "SELECT name FROM sqlite_schema WHERE type ='table' AND name NOT LIKE 'sqlite_%';"

func ShowTBLTRIPS(sqlite *Sqlite) {
	table := sqlite.GetTBLTRIPS("")
	for _, item := range table {
		log.Printf("item: %v\n", item)
	}
}

func ShowTBLTRIPUPDATES(sqlite *Sqlite) {
	table := sqlite.GetTBLTRIPUPDATES("")
	for _, item := range table {
		log.Printf("item: %v\n", item)
	}
}

func cleanDB(db *sqlx.DB) {
	tables := getTables(db)
	if tables != nil {
		log.Print("Found old tables, cleaning...")
	}
	for _, item := range tables {
		dropTableByName(db, item)

	}
	initTables(db)
}

func initTables(db *sqlx.DB) {
	log.Print("Creating table: TBLTRIPS")
	_, err := db.Exec(model.TBLTRIPS_sql)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Creating table: TBLTRIPUPDATES")
	_, err = db.Exec(model.TBLTRIPUPDATES_sql)
	if err != nil {
		log.Fatal(err)
	}

	log.Print("Enabling foreign_keys")
	_, err = db.Exec("PRAGMA foreign_keys = ON;")
	if err != nil {
		log.Fatal(err)
	}
}

func dropTableByName(db *sqlx.DB, name string) {
	log.Print("Droping table: " + name)

	sql := fmt.Sprintf("DROP TABLE %s;", name)
	_, err := db.Exec(sql)
	if err != nil {
		log.Fatal(err)
	}
}

func getTables(db *sqlx.DB) []string {
	rows, err := db.Query(getTablesSql)
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
