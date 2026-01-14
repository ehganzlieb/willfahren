package dto

type Line struct {
	Name string
	Type LineType
}

type LineType uint

const (
	LineTypeUBahn LineType = iota
	LineTypeSBahn
	LineTypeBadnerBahn
	LineTypeTram
	LineTypeBus
	LineTypeGroupTaxi
	LineTypeNightBus
	LineTypeNightGroupTaxi
)

func (lt LineType) String() string {
	return []string{"U-Bahn", "S-Bahn", "Badner Bahn", "Tram", "Bus", "Group Taxi", "Night Bus", "Night Group Taxi"}[lt]
}

func (l *Line) String() string {
	return l.Name
}

func (l *Line) TypeString() string {
	return l.Type.String()
}
