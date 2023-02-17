package app

import (
	"WeatherForecastGolang/internal"
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
	ParseWeather([]byte, time.Time) (*internal.Weather, int, int, error)
}

func Run(cityName string) {

	manager := NewManager(internal.API{}, internal.CityData{}, internal.Weather{}, time.Now())
	manager.Execute(cityName)
}
