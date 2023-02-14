package main

import (
	"fmt"
	"io"
	"net/http"
)

func getWeatherData(long, lat float64, cityName string) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,relativehumidity_2m,rain,showers,cloudcover", lat, long)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("couldn't request data")
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		fmt.Println(err.Error())
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err.Error())
		}
	}(res.Body)

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr.Error())
	}
	parseWeather(body, cityName)
}
