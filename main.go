package main

import (
	"os"
	"strings"
)

func main() {
	var Longitude, Latitude float64
	var cityName string

	argLength := len(os.Args[1:])
	if argLength > 0 {
		cityName = os.Args[1]
		Longitude, Latitude = getCoordinates(cityName)
	} else {
		cityName = getCity()
		Longitude, Latitude = getCoordinates(cityName)
	}
	weatherData := getWeatherData(Longitude, Latitude)
	weather, firstEntry, lastEntry := parseWeather(weatherData)
	displayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
