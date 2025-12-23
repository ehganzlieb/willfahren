package wlclient

import (
	_ "embed"
	"testing"

	"github.com/ehganzlieb/willwohnen/src/dto"

	"github.com/stretchr/testify/assert"
)

//go:embed wienerlinien-ogd-linien.csv
var testLines string

func TestParseLinesCSV(t *testing.T) {

	lines, err := parseLinesCSV(testLines)
	if err != nil {
		t.Fatal(err)
	}

	lineTypeMap := make(map[string]dto.LineType)
	for _, v := range lines {
		lineTypeMap[v.String()] = v.Type
	}

	assert.Equal(t, lineTypeMap["2"], dto.LineTypeTram)
	assert.Equal(t, lineTypeMap["U2"], dto.LineTypeUBahn)
	assert.Equal(t, lineTypeMap["S3"], dto.LineTypeSBahn)
	assert.Equal(t, lineTypeMap["13A"], dto.LineTypeBus)
	assert.Equal(t, lineTypeMap["BB"], dto.LineTypeBadnerBahn)
	assert.Equal(t, lineTypeMap["N6"], dto.LineTypeNightBus)

	nrOfType := make(map[dto.LineType]int)
	for k, v := range lineTypeMap {
		t.Logf("\t%s\t:\t%s\n", k, v)
		nrOfType[v]++
	}

	for k, v := range nrOfType {
		t.Logf("%dx %s", v, k)
	}
}
