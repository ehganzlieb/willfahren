package wlclient

import (
	_ "embed"
)

var (
	//go:embed OEFFHALTESTOGD.csv
	testStops string
)

/*
func TestParseStopsCSV(t *testing.T) {
	stops, err := parseStopsCSV(testStops)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range stops {
		 t.Logf("%+v", v)
	}
}
*/
