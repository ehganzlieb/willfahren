package wlclient

import "github.com/ehganzlieb/2025-10-26_willfahren/dto"

// AggregateStops() combines the list of stops with the list of lines from the parse functions into a dto map with unique stop and line objects that correctly reference each other.
func AggregateStops(stops []*Stop, lines []dto.Line) map[dto.Line][]*dto.Stop {
	stopMap := make(map[dto.Line][]*dto.Stop)

	for _, line := range lines {
		stopMap[line] = make([]*dto.Stop, 0)
	}

	for _, stop := range stops {
		for _, line := range stop.Lines {
			// compare line dto line names with stop line names
			for _, l := range lines {
				if l.Name == line {
					stopMap[l] = append(stopMap[l], &dto.Stop{
						Name:     stop.Name,
						Location: stop.Location,
					})

				}
			}

		}
	}

	return stopMap
}
