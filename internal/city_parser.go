package internal

import (
	"encoding/csv"
	"errors"
	"os"
	"strconv"
	"strings"
)

type CityData struct {
	CityName  string
	Longitude float64
	Latitude  float64
}

func GetCoordinates(cityName string) (float64, float64, error) {

	cityName = strings.ToLower(cityName)
	city := &CityData{}

	csvFile, err := os.Open("../assets/world-cities.csv")
	if err != nil {
		return 0, 0, err
	}
	defer func(csvFile *os.File) {
		err := csvFile.Close()
		if err != nil {
			return
		}
	}(csvFile)

	csvLines, err := csv.NewReader(csvFile).ReadAll()
	if err != nil {
		return 0, 0, err
	}

	for _, line := range csvLines {
		line[0] = strings.ToLower(line[0])
		if cityName == line[0] {

			if long, err := strconv.ParseFloat(line[3], 64); err == nil {
				city.Longitude = long
			}
			if lat, err := strconv.ParseFloat(line[2], 64); err == nil {
				city.Latitude = lat
			}
			city.CityName = line[0]
			return city.Longitude, city.Latitude, nil
		}
	}
	err404 := errors.New("city not found in Database")
	return 0, 0, err404
}
