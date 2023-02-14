package main

import (
	"os"
	"strings"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		arg := os.Args[1]
		Longitude, Latitude := getCoordinates(arg)
		getWeatherData(Longitude, Latitude, strings.ToUpper(arg))
	} else {
		Longitude, Latitude := getCoordinates("Tuttlingen")
		getWeatherData(Longitude, Latitude, "Tuttlingen")
	}
}
