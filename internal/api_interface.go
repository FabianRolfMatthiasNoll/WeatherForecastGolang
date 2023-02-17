package internal

import (
	"fmt"
	"io"
	"net/http"
)

type API struct {
	url  string
	body []byte
}

func (api API) GetData(url string) ([]byte, error) {
	var readErr error
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

	api.body, readErr = io.ReadAll(res.Body)
	if readErr != nil {
		return nil, readErr
	}
	return api.body, nil
}

func (api API) CreateAPILink(long, lat float64) string {
	api.url = fmt.Sprintf("https://api.open-meteo.com/v1/forecast?latitude=%f&longitude=%f&hourly=temperature_2m,relativehumidity_2m,rain,showers,cloudcover", lat, long)
	return api.url
}
