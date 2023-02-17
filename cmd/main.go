package main

import (
	"WeatherForecastGolang/cmd/app"
	"os"
)

//Build: go build -o ..\bin\Weather_Forecast_App
//Running the Linter: golangci-lint run
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
