package internal

import (
	"WeatherForecastGolang/assets"
	"strconv"
	"strings"
)

type CityData struct {
	CityName  string
	Longitude float64
	Latitude  float64
}

func (CityData) GetCoordinates(cityName string) (float64, float64, error) {
	var Longitude, Latitude float64
	var err error
	for _, city := range assets.WorldCities {
		if strings.Contains(city[0], cityName) {
			if Longitude, err = strconv.ParseFloat(city[2], 64); err != nil {
				return 0, 0, err
			}
			if Latitude, err = strconv.ParseFloat(city[1], 64); err != nil {
				return 0, 0, err
			}
			break
		}
	}
	return Longitude, Latitude, nil
}
