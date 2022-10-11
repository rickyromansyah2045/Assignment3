package main

import (
	"log"
	"net/http"
)

const PORT = ":8000"

type Status struct {
	Water int `json:"water"`
	Wind  int `json:"wind"`
}

type DataJSON struct {
	Status `json:"status"`
}

type Weather struct {
	WaterValue  int    `json:"water_value"`
	WaterStatus string `json:"water_status"`
	WindValue   int    `json:"wind_value"`
	WindStatus  int    `json:"wind_status"`
}

func main() {
	http.HandleFunc("/", action)

	log.Println("Server started at port", PORT)
	http.ListenAndServe(PORT, nil)
}

func action(w http.ResponseWriter, r *http.Request) {
	// var randomValue DataJSON

	// min :=
	// max :=
}
