package app

import (
	"WeatherForecastGolang/internal"
	"fmt"
	"strings"
	"time"
)

type Manager struct {
	api         APIHandler
	cityParser  CSVParser
	weatherData DataParser
	timer       time.Time
}

func NewManager(a APIHandler, c CSVParser, w DataParser, t time.Time) *Manager {
	return &Manager{
		api:         a,
		cityParser:  c,
		weatherData: w,
		timer:       t,
	}
}

func (m *Manager) Execute(cityName string) {
	var Longitude, Latitude float64
	var err error

	for Longitude == 0 && Latitude == 0 {
		if len(cityName) > 0 {
			Longitude, Latitude, err = m.cityParser.GetCoordinates(cityName)
		} else {
			cityName = internal.GetCity()
			Longitude, Latitude, err = m.cityParser.GetCoordinates(cityName)
		}
	}
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	apiUrl := m.api.CreateAPILink(Longitude, Latitude)
	weatherData, err := m.api.GetData(apiUrl)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	weather, firstEntry, lastEntry, err := m.weatherData.ParseWeather(weatherData, m.timer)
	if err != nil {
		fmt.Println(err.Error())
		time.Sleep(30 * time.Second)
	}
	internal.DisplayWeather(firstEntry, lastEntry, weather, strings.ToUpper(cityName))
}
