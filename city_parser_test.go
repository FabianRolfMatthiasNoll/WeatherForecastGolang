package main

import "testing"

type cityParseTest struct {
	wantLong, wantLat float64
	cityName          string
}

var parseTests = []cityParseTest{
	{8.8233, 47.9850, "tutTLingEn"},
	{8.8233, 47.9850, "Tuttlingen"},
	{7.3661, 46.2304, "Sion"},
	{7.3661, 46.2304, "sion"},
	{7.3661, 46.2304, "SION"},
}

func TestCityParse(t *testing.T) {

	for _, test := range parseTests {
		gotLong, gotLat := getCoordinates(test.cityName)
		if gotLong != test.wantLong || gotLat != test.wantLat {
			t.Errorf("got Long: %.4f, wanted: %.4f // got Lat: %.4f, wanted: %.4f", gotLong, test.wantLong, gotLat, test.wantLat)
		}
	}
}
