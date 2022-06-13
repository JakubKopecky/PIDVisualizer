package model

var TBLTRIPS_sql = `CREATE TABLE TBLTRIPS (
	id INTEGER PRIMARY KEY,
	trip_id TEXT
	);`

type TBLTRIPSStruct struct {
	Id      int    `db:"id"`
	Trip_id string `db:"trip_id"`
}

var TBLTRIPUPDATES_sql = `CREATE TABLE TBLTRIPUPDATES (
	id INTEGER PRIMARY KEY,
	trip_id INTEGER,
	data TEXT,
	FOREIGN KEY(trip_id) REFERENCES TBLTRIPS(id)
	);`

type TBLTRIPUPDATESStruct struct {
	Id      int    `db:"id"`
	Trip_id int    `db:"trip_id"`
	Data    string `db:"data"`
}
