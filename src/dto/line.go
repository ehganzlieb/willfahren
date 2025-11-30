package dto

type Line struct {
	Name  string
	Type  LineType
	Stops []*Stop
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
