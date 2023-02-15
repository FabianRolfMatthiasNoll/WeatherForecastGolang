package internal

import (
	"encoding/json"
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

func ParseWeather(data []byte) (Weather, int, int, error) {

	var weather Weather

	err := json.Unmarshal(data, &weather)
	if err != nil {
		return Weather{}, 0, 0, err
	}

	currentTime := time.Now()
	var day = currentTime.String()
	day = day[:10] //cutting off everything that doesn't include the date

	var lastEntry int
	var firstEntry int

	for i, s := range weather.Hourly.Time {
		if strings.Contains(s, day) {
		} else {
			lastEntry = i
			break
		}

		var currentHour = currentTime.Hour()
		//If its over 15 minutes we can start with the next hour
		if currentTime.Minute() > 15 {
			currentHour++
		}
		hour, err := strconv.Atoi(s[11:13])
		if err != nil {
			return Weather{}, 0, 0, err
		}
		if hour < currentHour {
			firstEntry = i + 1
		}
	}
	return weather, firstEntry, lastEntry, nil
}
