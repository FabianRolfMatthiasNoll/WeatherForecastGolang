package main

import (
	"WeatherForecastGolang/cmd/app"
	"os"
)

//go:generate go run generate-csv.go

func main() {
	argLength := len(os.Args[1:])
	if argLength > 0 {
		cityName := os.Args[1]
		app.Run(cityName)
	} else {
		app.Run("")
	}
}
