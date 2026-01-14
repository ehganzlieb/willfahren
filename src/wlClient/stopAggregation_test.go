package wlclient

import (
	_ "embed"
	"testing"
)

func TestAggregateStops(t *testing.T) {
	// get stops from csv
	stops, err := ParseStopsCSV(testStops)
	if err != nil {
		t.Fatal(err)
	}

	//get lines from csv
	lines, err := ParseLinesCSV(testLines)
	if err != nil {
		t.Fatal(err)
	}
}
