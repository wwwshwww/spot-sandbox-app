package dbscan_profile

type DistanceType string

const (
	Hubeny      DistanceType = "Hubeny"
	RouteLength DistanceType = "RootLength"
	TravelTime  DistanceType = "TravelTime"
)
