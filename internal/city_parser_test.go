package internal

import "testing"

type cityParseTest struct {
	wantLong, wantLat float64
	cityName          string
	errWant           error
}

var parseTests = []cityParseTest{
	{8.8233, 47.9850, "tutTLingEn", nil},
	{8.8233, 47.9850, "Tuttlingen", nil},
	{7.3661, 46.2304, "Sion", nil},
	{7.3661, 46.2304, "sion", nil},
	{7.3661, 46.2304, "SION", nil},
}

func TestCityParse(t *testing.T) {

	for _, test := range parseTests {
		gotLong, gotLat, err := GetCoordinates(test.cityName)
		if gotLong != test.wantLong || gotLat != test.wantLat || err != test.errWant {
			t.Errorf("got Long: %.4f, wanted: %.4f // got Lat: %.4f, wanted: %.4f // got err: %d, wanted: %d", gotLong, test.wantLong, gotLat, test.wantLat, err, test.errWant)
		}
	}
}
