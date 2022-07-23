package entity

type DistanceFunction int

const (
	Hubeny DistanceFunction = iota * 10
	PathLength
	TravelTime
)

type DbscanPlofile struct {
	Epsilon  float64
	MinCount *uint64
	MaxCount *uint64
	Distance DistanceFunction
}
