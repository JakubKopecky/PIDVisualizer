package controller

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/JakubKopecky/PIDVisualizer/database"
	"github.com/JakubKopecky/PIDVisualizer/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Sqlite struct {
	db *sqlx.DB
}

func NewDbController(file string) *Sqlite {
	log.Print("Opening db file: " + file)
	db, err := sqlx.Open("sqlite3", file)
	if err != nil {
		log.Fatal(err)
	}
	database.CleanDB(db)

	return &Sqlite{db: db}
}

func (sqlite *Sqlite) Exec(sql string) (sql.Result, error) {
	log.Print("Executing: " + sql)
	return sqlite.db.Exec(sql)
}

func (sqlite *Sqlite) Query(sql string) (*sqlx.Rows, error) {
	log.Print("Executing: " + sql)
	return sqlite.db.Queryx(sql)
}

func (sqlite *Sqlite) GetTBLTRIPS(where string) []model.TBLTRIPSStruct {
	rows, err := sqlite.Query(fmt.Sprintf("SELECT * FROM TBLTRIPS %s;", where))
	if err != nil {
		log.Print(err)
		return nil
	}
	var toReturn []model.TBLTRIPSStruct
	var line model.TBLTRIPSStruct
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&line)
		if err != nil {
			log.Panic(err)
		}
		toReturn = append(toReturn, line)
	}
	return toReturn
}

func (sqlite *Sqlite) GetTBLTRIPUPDATES(where string) []model.TBLTRIPUPDATESStruct {
	rows, err := sqlite.Query(fmt.Sprintf("SELECT * FROM TBLTRIPUPDATES %s;", where))
	if err != nil {
		log.Print(err)
		return nil
	}
	var toReturn []model.TBLTRIPUPDATESStruct
	var line model.TBLTRIPUPDATESStruct
	defer rows.Close()
	for rows.Next() {
		err := rows.StructScan(&line)
		if err != nil {
			log.Panic(err)
		}
		toReturn = append(toReturn, line)
	}
	return toReturn
}

func (sqlite *Sqlite) InsertIntoTBLTRIPS(data model.TBLTRIPSStruct) error {
	var query = "INSERT INTO TBLTRIPS VALUES (null, :trip_id);"
	_, err := sqlite.db.NamedExec(query, data)
	return err
}

func (sqlite *Sqlite) InsertIntoTBLTRIPUPDATES(data model.TBLTRIPUPDATESStruct) error {
	var query = "INSERT INTO TBLTRIPUPDATES VALUES (null, :trip_id, :data);"
	_, err := sqlite.db.NamedExec(query, data)
	return err
}

func (sqlite *Sqlite) Close() {
	sqlite.db.Close()
}
