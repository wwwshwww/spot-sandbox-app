package dbscan_profile_mysql

import "time"

type DistanceType string

const (
	Hubeny      DistanceType = "Hubeny"
	RouteLength DistanceType = "RootLength"
	TravelTime  DistanceType = "TravelTime"
)

type DbscanProfile struct {
	ID                uint `gorm:"primaryKey"`
	DistanceType      DistanceType
	MinCount          int
	MaxCount          *int
	MeterThreshold    *int
	DurationThreshold *time.Duration
}
