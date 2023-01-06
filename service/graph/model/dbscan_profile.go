package model

type DbscanProfile struct {
	Key              int          `json:"key"`
	DistanceType     DistanceType `json:"distanceType"`
	MinCount         int          `json:"minCount"`
	MaxCount         *int         `json:"maxCount"`
	MeterThreshold   *int         `json:"meterThreshold"`
	MinutesThreshold *int         `json:"minutesThreshold"`
}
