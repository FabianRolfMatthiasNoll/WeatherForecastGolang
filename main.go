package main

import (
	"fmt"
	"os"
)

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		arg := os.Args[1]
		Longitude, Latitude := getCoordinates(arg)
		fmt.Println(Longitude, Latitude) //Please remove later on
		getWeatherData(Longitude, Latitude)
	} else {
		fmt.Println("Hello World")
	}
}
