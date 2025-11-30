package wlclient

import (
	"encoding/csv"
	"strings"

	"github.com/ehganzlieb/willwohnen/src/dto"
)

func parseLinesCSV(input string) ([]dto.Line, error) {
	r := csv.NewReader(strings.NewReader(input))

}
