package wlclient

import (
	"encoding/csv"
	"strings"

	"github.com/ehganzlieb/willwohnen/src/dto"
)

func parseLinesCSV(input string) ([]dto.Line, error) {
	r := csv.NewReader(strings.NewReader(input))

	//get field names from first line
	fieldNames, err := r.Read()
	if err != nil {
		return nil, err
	}

	indexMap := make(map[string]int)

	for i, v := range fieldNames {
		indexMap[v] = i
	}

	lines := make([]dto.Line, 0)

}
