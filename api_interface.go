package main

import (
	"fmt"
	"io"
	"net/http"
)

// adding error handling => return error
func getWeatherData(long, lat float64) ([]byte, error) {
	url := fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,relativehumidity_2m,rain,showers,cloudcover", lat, long)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(res.Body)

	body, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	return body, nil
}
