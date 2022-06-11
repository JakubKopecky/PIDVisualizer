package model

import "time"

type VehiclePositionsStruct struct {
	Features []struct {
		Geometry struct {
			Coordinates []float64 `json:"coordinates"`
			Type        string    `json:"type"`
		} `json:"geometry"`
		Properties struct {
			LastPosition struct {
				Bearing int `json:"bearing"`
				Delay   struct {
					Actual            int         `json:"actual"`
					LastStopArrival   interface{} `json:"last_stop_arrival"`
					LastStopDeparture interface{} `json:"last_stop_departure"`
				} `json:"delay"`
				IsCanceled interface{} `json:"is_canceled"`
				LastStop   struct {
					ArrivalTime   time.Time `json:"arrival_time"`
					DepartureTime time.Time `json:"departure_time"`
					ID            string    `json:"id"`
					Sequence      int       `json:"sequence"`
				} `json:"last_stop"`
				NextStop struct {
					ArrivalTime   time.Time `json:"arrival_time"`
					DepartureTime time.Time `json:"departure_time"`
					ID            string    `json:"id"`
					Sequence      int       `json:"sequence"`
				} `json:"next_stop"`
				OriginTimestamp   time.Time   `json:"origin_timestamp"`
				ShapeDistTraveled string      `json:"shape_dist_traveled"`
				Speed             interface{} `json:"speed"`
				StatePosition     string      `json:"state_position"`
				Tracking          bool        `json:"tracking"`
				ValidTo           string      `json:"valid_to"`
			} `json:"last_position"`
			Trip struct {
				AgencyName struct {
					Real      string `json:"real"`
					Scheduled string `json:"scheduled"`
				} `json:"agency_name"`
				Cis struct {
					LineID     interface{} `json:"line_id"`
					TripNumber interface{} `json:"trip_number"`
				} `json:"cis"`
				Gtfs struct {
					RouteID        string      `json:"route_id"`
					RouteShortName string      `json:"route_short_name"`
					RouteType      int         `json:"route_type"`
					TripHeadsign   string      `json:"trip_headsign"`
					TripID         string      `json:"trip_id"`
					TripShortName  interface{} `json:"trip_short_name"`
				} `json:"gtfs"`
				OriginRouteName           string    `json:"origin_route_name"`
				SequenceID                int       `json:"sequence_id"`
				StartTimestamp            time.Time `json:"start_timestamp"`
				VehicleRegistrationNumber int       `json:"vehicle_registration_number"`
				VehicleType               struct {
					DescriptionCs string `json:"description_cs"`
					DescriptionEn string `json:"description_en"`
					ID            int    `json:"id"`
				} `json:"vehicle_type"`
				WheelchairAccessible bool        `json:"wheelchair_accessible"`
				AirConditioned       interface{} `json:"air_conditioned"`
			} `json:"trip"`
		} `json:"properties"`
		Type string `json:"type"`
	} `json:"features"`
	Type string `json:"type"`
}
