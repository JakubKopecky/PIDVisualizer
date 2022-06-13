package controller

import (
	"log"
)

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
