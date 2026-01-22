package dto

import "fmt"

type Stop struct {
	Name     string
	Location Coordinates
	Lines    *[]Line
}

func (s *Stop) String() string {
	str := fmt.Sprintf("%s (%f, %f) [", s.Name, s.Location.X, s.Location.Y)
	for i, l := range *s.Lines {
		if i > 0 {
			str += ", "
		}
		str += l.String()
	}
	str += "]"
	return str
}
