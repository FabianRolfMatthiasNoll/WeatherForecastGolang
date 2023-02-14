package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type cityData struct {
	CityName  string
	Longitude float64
	Latitude  float64
}

func getCoordinates(cityName string) (float64, float64) {
	csvFile, err := os.Open("assets/world-cities.csv")
	if err != nil {
		fmt.Println("File not Found")
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			fmt.Println("File couldn't be closed")
		}
	}(csvFile)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		fmt.Println("Input not readable")
	}
	for _, line := range csvLines {
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
