package main

import (
	"log"

	"github.com/JakubKopecky/PIDVisualizer/controller"
)

func main() {
	//callApi()

	sqlite := controller.NewDbController("./pid.db")
	//sqlite.Exec("create table memos(text, priority INTEGER);")
	//sqlite.Exec("insert into memos values('deliver project description', 10);")
	//sqlite.Exec("insert into memos values('lunch with Christine', 100);")
	sqlite.Close()
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
