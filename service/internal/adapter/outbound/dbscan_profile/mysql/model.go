package dbscan_profile_mysql

import "time"

type DistanceType string

const (
	Hubeny      DistanceType = "Hubeny"
	RouteLength DistanceType = "RootLength"
	TravelTime  DistanceType = "TravelTime"
)

type DbscanProfile struct {
	ID                uint `gorm:"primaryKey;columns:id"`
	DistanceType      DistanceType
	MinCount          uint
	MaxCount          *uint
	MeterThreshold    *int
	DurationThreshold *time.Duration
}
