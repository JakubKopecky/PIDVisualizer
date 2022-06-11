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

func main() {
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

	log.Printf("Api call returned %d bytes.", len(responseData))
	log.Printf("long: %f", parsed.Features[0].Geometry.Coordinates[0])
	log.Printf("lati: %f", parsed.Features[0].Geometry.Coordinates[1])
}
