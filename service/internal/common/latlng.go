package common

import "fmt"

type LatLng struct {
	Lat float64
	Lng float64
}

func (l LatLng) String() string {
	return fmt.Sprintf("%f,%f", l.Lat, l.Lng)
}
