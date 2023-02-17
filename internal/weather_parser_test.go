package internal

import (
	"encoding/json"
	"testing"
	"time"
)

type weatherParseTest struct {
}

func TestWeather_ParseWeather(t *testing.T) {
	data := map[string]interface{}{
		"latitude":              37.7749,
		"longitude":             -122.4194,
		"generationtime_ms":     1645082402000,
		"utc_offset_seconds":    -28800,
		"timezone":              "America/Los_Angeles",
		"timezone_abbreviation": "PST",
		"elevation":             47.2,
		"hourly_units": map[string]string{
			"time":                "America/Los_Angeles",
			"temperature_2m":      "C",
			"relativehumidity_2m": "%",
			"rain":                "mm",
			"showers":             "mm",
			"cloudcover":          "%",
		},
		"hourly": map[string]interface{}{
			"time":                []string{"2023-02-16T12:00", "2023-02-16T13:00", "2023-02-16T14:00", "2023-02-16T15:00", "2023-02-17T15:00", "2023-02-18T15:00"},
			"temperature_2m":      []float64{10.0, 12.5, 14.0, 13.5, 17.8},
			"relativehumidity_2m": []int{65, 60, 58, 62, 90},
			"rain":                []float64{0.0, 0.1, 0.2, 0.0, 8.9},
			"showers":             []float64{0.0, 0.2, 0.4, 0.1, 7.3},
			"cloudcover":          []int{10, 15, 20, 30, 88},
		},
	}
	jsonData, _ := json.Marshal(data)

	clock := time.Date(2023, 02, 16, 13, 05, 0, 0, time.UTC)
	weather, fEntry, lEntry, err := Weather.ParseWeather(Weather{}, jsonData, clock)
	if weather.Elevation != 47.2 {
		t.Errorf("Wanted: %.1f, Got: %.1f", 47.2, weather.Elevation)
	}
	if fEntry != 1 {
		t.Errorf("Wrong First Index. Wanted: 1, Got: %d", fEntry)
	}
	if lEntry != 3 {
		t.Errorf("Wrong Last Index. Wanted: 3, Got: %d and its; %s", lEntry, weather.Hourly.Time[lEntry])
	}
	if err != nil {
		t.Errorf(err.Error())
	}

}
