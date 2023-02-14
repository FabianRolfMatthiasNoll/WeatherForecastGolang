package main

import (
	"os"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		arg := os.Args[1]
		Longitude, Latitude := getCoordinates(arg)
		getWeatherData(Longitude, Latitude)
	} else {
		Longitude, Latitude := getCoordinates("Tuttlingen")
		getWeatherData(Longitude, Latitude)
	}
}
