package app

import (
	"WeatherForecastGolang/internal"
	"fmt"
	"strings"
	"time"
)

type APIHandler interface {
	GetData(string) ([]byte, error)
	CreateAPILink(float64, float64) string
}

type CSVParser interface {
	GetCoordinates(string) (float64, float64, error)
}

type DataParser interface {
	ParseWeather(data []byte) (internal.Weather, int, int, error)
}

// Build: go build -o ..\bin\Weather_Forecast_App
func Run(cityName string) {
	var Longitude, Latitude float64
	var err error

	var weatherAPI internal.API
	var api APIHandler = weatherAPI

	var cityParser internal.CityData
	var city CSVParser = cityParser

	var weatherParser internal.Weather
	var weatherParsed DataParser = weatherParser

	for Longitude == 0 && Latitude == 0 {
		if len(cityName) > 0 {
			Longitude, Latitude, err = city.GetCoordinates(cityName)
		} else {
			cityName = internal.GetCity()
			Longitude, Latitude, err = city.GetCoordinates(cityName)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	apiUrl := api.CreateAPILink(Longitude, Latitude)
	weatherData, err := api.GetData(apiUrl)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	weather, firstEntry, lastEntry, err := weatherParsed.ParseWeather(weatherData)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	internal.DisplayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
