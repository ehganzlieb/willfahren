package dto

import "fmt"

type District struct {
	Name         string
	InsiderNames []string
	Number       int
}

func (d *District) PostCode() int {
	return 1000 + d.Number*10
}

func DistrictFromPostCode(postcode int) (*District, error) {
	d, ok := districts[(postcode-1000)/10]
	if !ok {
		return nil, fmt.Errorf("no district for postcode %d", postcode)
	}
	return &d, nil
}

var districts = map[int]District{
	0: {
		Name:         "Ganz Wien",
		InsiderNames: []string{"Iwaroi, Gåunze Stådt"},
		Number:       0,
	},
	1: {
		Name:         "Innere Stadt",
		InsiderNames: []string{"I-Town", "Disneyland"},
		Number:       1,
	},
	2: {
		Name:         "Leopoldstadt",
		InsiderNames: []string{"Leo"},
		Number:       2,
	},
	3: {
		Name:         "Landstraße",
		InsiderNames: []string{},
		Number:       3,
	},
	4: {
		Name:         "Wieden",
		InsiderNames: []string{},
		Number:       4,
	},
	5: {
		Name:         "Margareten",
		InsiderNames: []string{"Little Berlin"},
		Number:       5,
	},
	6: {
		Name:         "Mariahilf",
		InsiderNames: []string{"MaHüf"},
		Number:       6,
	},
	7: {
		Name:         "Neubau",
		InsiderNames: []string{"Bobostan"},
		Number:       7,
	},
	8: {
		Name:         "Josefstadt",
		InsiderNames: []string{"Moneytown", "Bürgistan"},
		Number:       8,
	},
	9: {
		Name:         "Alsergrund",
		InsiderNames: []string{},
		Number:       9,
	},
	10: {
		Name:         "Favoriten",
		InsiderNames: []string{"Little Istanbul", "X"},
		Number:       10,
	},
	11: {
		Name:         "Simmering",
		InsiderNames: []string{},
		Number:       11,
	},
	12: {
		Name:         "Meidling",
		InsiderNames: []string{"MeidLing", "Big L"},
		Number:       12,
	},
	13: {
		Name:         "Hietzing",
		InsiderNames: []string{},
		Number:       13,
	},
	14: {
		Name:         "Penzing",
		InsiderNames: []string{},
		Number:       14,
	},
	15: {
		Name:         "Rudolfsheim-Fünfhaus",
		InsiderNames: []string{"Rudolfscrime", "RH5H"},
		Number:       15,
	},
	16: {
		Name:         "Ottakring",
		InsiderNames: []string{"OTK"},
		Number:       16,
	},
	17: {
		Name:         "Hernals",
		InsiderNames: []string{},
		Number:       17,
	},
	18: {
		Name:         "Währing",
		InsiderNames: []string{},
		Number:       18,
	},
	19: {
		Name:         "Döbling",
		InsiderNames: []string{},
		Number:       19,
	},
	20: {
		Name:         "Brigittenau",
		InsiderNames: []string{"Brighettonau"},
		Number:       20,
	},
	21: {
		Name:         "Floridsdorf",
		InsiderNames: []string{"Flodorf", "Flowtown", "Flodo"},
		Number:       21,
	},
	22: {
		Name:         "Donaustadt",
		InsiderNames: []string{"DC", "Donau City"},
		Number:       22,
	},
	23: {
		Name:         "Liesing",
		InsiderNames: []string{"Mietkauf"},
		Number:       23,
	},
}

func DistrictByNumber(number int) (*District, error) {
	d, ok := districts[number]
	if !ok {
		return nil, fmt.Errorf("no district for number %d", number)
	}

	//return copy of district to keep map elements immutable
	return &District{
		Name:         d.Name,
		InsiderNames: d.InsiderNames,
		Number:       d.Number,
	}, nil

}
