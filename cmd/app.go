package cmd

import (
	"WeatherForecastGolang/internal"
	"fmt"
	"strings"
)

// Build: go build -o ..\bin\Weather_Forecast_App
// IMplement interfaces that are needed
func Run(cityName string) {
	var Longitude, Latitude float64
	var err error

	for Longitude == 0 && Latitude == 0 {
		if len(cityName) > 0 {
			Longitude, Latitude, err = internal.GetCoordinates(cityName)
		} else {
			cityName = internal.GetCity()
			Longitude, Latitude, err = internal.GetCoordinates(cityName)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	weatherData, err := internal.GetWeatherData(Longitude, Latitude)

	weather, firstEntry, lastEntry, err := internal.ParseWeather(weatherData)
	if err != nil {
		fmt.Println(err.Error())
	}

	internal.DisplayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
