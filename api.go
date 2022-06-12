package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/JakubKopecky/PIDVisualizer/model"
)

var url string = "https://api.golemio.cz/v2/vehiclepositions?routeShortName=19"

func callApi() {
	client := &http.Client{}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Panic(err.Error())
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("X-Access-Token", os.Getenv("TOKEN"))

	response, err := client.Do(req)
	if err != nil {
		log.Panic(err.Error())
	}

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Panic(err.Error())
	}

	var parsed model.VehiclePositionsStruct
	err = json.Unmarshal(responseData, &parsed)
	if err != nil {
		log.Panic(err.Error())
	}

	for index := range parsed.Features {
		trip_id := parsed.Features[index].Properties.Trip.Gtfs.TripID
		long := parsed.Features[index].Geometry.Coordinates[0]
		lati := parsed.Features[index].Geometry.Coordinates[1]
		log.Printf("trip_id: %s\n", trip_id)
		log.Printf("\tlati: %f\n", lati)
		log.Printf("\tlong: %f\n", long)
	}
}
