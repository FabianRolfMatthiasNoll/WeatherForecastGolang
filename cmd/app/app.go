package app

import (
	"WeatherForecastGolang/internal"
	"fmt"
	"strings"
	"time"
)

// Build: go build -o ..\bin\Weather_Forecast_App
// Implement interfaces that are needed
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
		time.Sleep(30 * time.Second)
	}
	weatherData, err := internal.GetWeatherData(Longitude, Latitude)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	weather, firstEntry, lastEntry, err := internal.ParseWeather(weatherData)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	internal.DisplayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
