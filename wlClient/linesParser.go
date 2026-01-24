package wlclient

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"

	"github.com/ehganzlieb/willfahren/dto"
)

const (
	LineNameField = "BEZEICHNUNG"
	LineTypeField = "VERKEHRSMITTEL"

	TramString           = "ptTram"
	UBahnString          = "ptMetro"
	SBahnString          = "ptTrainS"
	BusString            = "ptBusCity"
	NightBusString       = "ptBusNight"
	GroupTaxiString      = "pt_RufbusTag"
	BadnerBahnString     = "ptBadner_Bahn"
	NightGroupTaxiString = "Pt_RufbusNacht"
)

/*
ParseLinesCSV parses a CSV string and returns a slice of dto.Lines.

The CSV string is expected to have the columns BEZEICHNUNG and VERKEHRSMITTEL, and to contain lines with the values ptTram, ptMetro, ptTrainS, ptBusCity, ptBusNight, pt_RufbusTag, ptBadner_Bahn, and pt_RufbusNacht for VERKEHRSMITTEL.
The function will return an error if the CSV string is malformed, or if a line contains an unknown VERKEHRSMITTEL value.
The returned slice of dto.Lines will contain the parsed lines, and will not contain the stops as they are not contained in this file.
*/
func ParseLinesCSV(input string) ([]dto.Line, error) {
	r := csv.NewReader(strings.NewReader(input))
	r.Comma = ';'

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

	//parse all lines
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Print(err)
			}
		}
		l := dto.Line{
			Name: record[indexMap[LineNameField]],
		}
		switch record[indexMap[LineTypeField]] {
		case TramString:
			l.Type = dto.LineTypeTram
		case UBahnString:
			l.Type = dto.LineTypeUBahn
		case SBahnString:
			l.Type = dto.LineTypeSBahn
		case BusString:
			l.Type = dto.LineTypeBus
		case NightBusString:
			l.Type = dto.LineTypeNightBus
		case GroupTaxiString:
			l.Type = dto.LineTypeGroupTaxi
		case BadnerBahnString:
			l.Type = dto.LineTypeBadnerBahn
		case NightGroupTaxiString:
			l.Type = dto.LineTypeNightGroupTaxi
		default:
			return nil, fmt.Errorf("unknown line type %s", record[indexMap[LineTypeField]])
		}
		lines = append(lines, l)
	}
	return lines, nil

}
