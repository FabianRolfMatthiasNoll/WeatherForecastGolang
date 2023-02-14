package main

import (
	"encoding/json"
	"fmt"
)

type Weather struct {
	Latitude             float64 `json:"latitude"`
	Longitude            float64 `json:"longitude"`
	GenerationtimeMs     float64 `json:"generationtime_ms"`
	UtcOffsetSeconds     int     `json:"utc_offset_seconds"`
	Timezone             string  `json:"timezone"`
	TimezoneAbbreviation string  `json:"timezone_abbreviation"`
	Elevation            float64 `json:"elevation"`
	HourlyUnits          struct {
		Time               string `json:"time"`
		Temperature2M      string `json:"temperature_2m"`
		Relativehumidity2M string `json:"relativehumidity_2m"`
		Rain               string `json:"rain"`
		Showers            string `json:"showers"`
		Cloudcover         string `json:"cloudcover"`
	} `json:"hourly_units"`
	Hourly struct {
		Time               []string  `json:"time"`
		Temperature2M      []float64 `json:"temperature_2m"`
		Relativehumidity2M []int     `json:"relativehumidity_2m"`
		Rain               []float64 `json:"rain"`
		Showers            []int     `json:"showers"`
		Cloudcover         []int     `json:"cloudcover"`
	} `json:"hourly"`
}

func parseWeather(data []byte) {
	var weather Weather
	err := json.Unmarshal(data, &weather)
	if err != nil {
	}
	for _, value := range weather.Hourly.Temperature2M {
		fmt.Println(value)
	}

}
