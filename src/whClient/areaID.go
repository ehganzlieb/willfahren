package whclient

import (
	"fmt"

	"github.com/ehganzlieb/willwohnen/src/dto"
)

func AreaID(d *dto.District) (uint64, error) {
	id, ok := areaIDs[d.PostCode()]
	if !ok {
		return 0, fmt.Errorf("no area id for postcode %s", d.Name)
	}
	return id, nil
}

var areaIDs = map[int]uint64{ // postcode to WH area id
	1000: 900,
	1010: 117223,
	1020: 117224,
	1030: 117225,
	1040: 117226,
	1050: 117227,
	1060: 117228,
	1070: 117229,
	1080: 117230,
	1090: 117231,
	1100: 117232,
	1110: 117233,
	1120: 117234,
	1130: 117235,
	1140: 117236,
	1150: 117237,
	1160: 117238,
	1170: 117239,
	1180: 117240,
	1190: 117241,
	1200: 117242,
	1210: 117243,
	1220: 117244,
	1230: 117245,
}
