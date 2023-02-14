package main

import (
	"encoding/csv"
	"os"
	"strconv"
	"strings"
)

type cityData struct {
	CityName  string
	Longitude float64
	Latitude  float64
}

func getCoordinates(cityName string) (float64, float64) {

	cityName = strings.ToLower(cityName)

	csvFile, err := os.Open("assets/world-cities.csv")
	if err != nil {
		return 0, 0
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			return
		}
	}(csvFile)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return 1, 1
	}

	for _, line := range csvLines {
		line[0] = strings.ToLower(line[0])
		if cityName == line[0] {
			city := cityData{}
			if long, err := strconv.ParseFloat(line[2], 64); err == nil {
				city.Longitude = long
			}
			if lat, err := strconv.ParseFloat(line[3], 64); err == nil {
				city.Latitude = lat
			}
			return city.Longitude, city.Latitude
		}
	}
	return 0, 0
}
