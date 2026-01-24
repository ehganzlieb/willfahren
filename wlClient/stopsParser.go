package wlclient

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
	"sync"

	"github.com/ehganzlieb/willfahren/dto"
)

type Stop struct {
	Name      string
	ShortName string
	Location  dto.Coordinates
	Lines     []string
}

const (
	StopNameField      = "HTXT"
	StopShortNameField = "HTXTK"
	CoordsField        = "SHAPE"
	LinesField         = "HLINIEN"

	CoordsFormatString = "POINT (%f %f)"
)

/*
ParseStopsCSV parses a CSV string and returns a slice of Stop.

The CSV string is expected to have the columns HTXT, HTXTK, SHAPE, and HLINIEN. Other columns are ignored.
The function will return an error if the CSV string is malformed.
The returned slice of Stop will contain the parsed stops.
The stops contain lines with line names
*/
func ParseStopsCSV(input string) ([]*Stop, error) {
	//TODO: concurrency

	r := csv.NewReader(strings.NewReader(input))

	stops := make([]*Stop, 0)
	//get field names from first line
	fieldNames, err := r.Read()
	if err != nil {
		return nil, err
	}

	indexMap := make(map[string]int)
	for i, v := range fieldNames {
		indexMap[v] = i
	}

	//parse all lines
	ch := make(chan *Stop, 50)
	var wg sync.WaitGroup
	for {
		record, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Print(err)
			}
		}
		go parseStop(record, indexMap, ch, &wg)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()

	for stop := range ch {
		if stop != nil {
			stops = append(stops, stop)
		}
	}

	return stops, nil
}

func parseStop(record []string, indexMap map[string]int, ch chan *Stop, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	name := record[indexMap[StopNameField]]
	shortName := record[indexMap[StopShortNameField]]
	coordsString := record[indexMap[CoordsField]]
	linesString := record[indexMap[LinesField]]

	var xCoord, yCoord float64
	if _, err := fmt.Sscanf(coordsString, CoordsFormatString, &xCoord, &yCoord); err != nil {
		ch <- nil
		return
	}

	lines := strings.Split(linesString, ",")
	ch <- &Stop{
		Name:      name,
		ShortName: shortName,
		Location: dto.Coordinates{
			X: xCoord,
			Y: yCoord,
		},
		Lines: lines,
	}
}
