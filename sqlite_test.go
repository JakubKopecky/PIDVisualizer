package main_test

import (
	"io/ioutil"
	"log"
	"os"
	"testing"

	"github.com/JakubKopecky/PIDVisualizer/controller"
	"github.com/JakubKopecky/PIDVisualizer/model"
)

var sqlite *controller.Sqlite

func TestMain(m *testing.M) {
	log.SetOutput(ioutil.Discard)
	log.SetFlags(0)
	sqlite = controller.NewDbController("./pid.db")
	os.Exit(m.Run())
}

func TestSqliteInsert(t *testing.T) {

	data := model.TBLTRIPSStruct{Trip_id: "19_2342_1233"}
	err := sqlite.InsertIntoTBLTRIPS(data)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(sqlite.GetTBLTRIPS("WHERE trip_id='19_2342_1233'")) != 1 {
		t.Fatalf("Returned wrong number of rows")
	}

	data2 := model.TBLTRIPUPDATESStruct{Trip_id: 1, Data: "{DATA}"}
	err = sqlite.InsertIntoTBLTRIPUPDATES(data2)
	if err != nil {
		t.Fatalf(err.Error())
	}
	if len(sqlite.GetTBLTRIPUPDATES("WHERE trip_id='1'")) != 1 {
		t.Fatalf("Returned wrong number of rows")
	}
}

func TestSqliteGet(t *testing.T) {
	rowsTBLTRIPS := sqlite.GetTBLTRIPS("WHERE trip_id='19_2342_1233'")
	if rowsTBLTRIPS == nil {
		t.Fatalf("GetTBLTRIPS() returned nothing")
	}
	if len(rowsTBLTRIPS) != 1 {
		t.Fatalf("GetTBLTRIPS() returned more then expected")
	}

	rowsTBLTRIPUPDATES := sqlite.GetTBLTRIPUPDATES("WHERE trip_id='1'")
	if rowsTBLTRIPUPDATES == nil {
		t.Fatalf("GetTBLTRIPUPDATES() returned nothing")
	}
	if len(rowsTBLTRIPUPDATES) != 1 {
		t.Fatalf("GetTBLTRIPUPDATES() returned more then expected")
	}
}

func TestSqliteError(t *testing.T) {
	_, err := sqlite.Exec("DELETE FROM TBLTRIPS WHERE id='1';")
	if err == nil {
		t.Fatalf("there should be a foreign key error but is none")
	}
	_, err = sqlite.Exec("INSERT INTO TBLTRIPUPDATES VALUES (null, 3, \"{DATA}\");")
	if err == nil {
		t.Fatalf("there should be a foreign key error but is none")
	}
}
