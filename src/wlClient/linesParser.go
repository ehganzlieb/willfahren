package wlclient

import (
	"encoding/csv"
	"io"
	"log"
	"strings"

	"github.com/ehganzlieb/willwohnen/src/dto"
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
		}
		lines = append(lines, l)
	}
	return lines, nil

}
