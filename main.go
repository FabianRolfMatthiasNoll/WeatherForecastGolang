package main

import (
	"fmt"
	"os"
	"strings"
)

// Implement Interfaces here that are needed from manager
func main() {
	var Longitude, Latitude float64
	var cityName string
	var err error

	for Longitude == 0 && Latitude == 0 {
		argLength := len(os.Args[1:])
		if argLength > 0 {
			cityName = os.Args[1]
			Longitude, Latitude, err = getCoordinates(cityName)
		} else {
			cityName = getCity()
			Longitude, Latitude, err = getCoordinates(cityName)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherData, err := getWeatherData(Longitude, Latitude)

	weather, firstEntry, lastEntry, err := parseWeather(weatherData)
	if err != nil {
		fmt.Println(err.Error())
	}

	displayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
