package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	"time"
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
	//weather.cleanWeather()
	currentTime := time.Now()

	var day = currentTime.String()
	day = day[:10]

	var lastEntry int
	var firstEntry int

	for i, s := range weather.Hourly.Time {
		if strings.Contains(s, day) {
		} else {
			lastEntry = i
			break
		}

		var currentHour = currentTime.Hour()
		hour, err := strconv.Atoi(s[11:13])
		if err != nil {
			fmt.Println("Int Conversion failed")
		}
		if hour < currentHour {
			firstEntry = i
		}
	}
	displayWeather(firstEntry, lastEntry)
}

/*
func (w *Weather) cleanWeather() {
	currentTime := time.Now()

	var day = currentTime.String()
	day = day[:10]

	fmt.Println(day)
	var result []string
	for _, s := range w.Hourly.Time {
		//fmt.Println(s)
		if strings.Contains(s, day) {
			result = append(result, s)
		}
	}
	w.Hourly.Time = result
	result = nil
	var currentHour = currentTime.Hour()

	for i, s := range w.Hourly.Time {
		hour,err := strconv.ParseInt(s[10:12], 2,64)
		if err != nil {
			fmt.Println("Int Conversion failed")
		}
		if hour >= currentHour {
			result = append(result, s)
		}
	}

	fmt.Println(result)
} */
