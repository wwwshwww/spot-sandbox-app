package model

type DbscanProfile struct {
	ID               string       `json:"id"`
	DistanceType     DistanceType `json:"distanceType"`
	MinCount         int          `json:"minCount"`
	MaxCount         *int         `json:"maxCount"`
	MeterThreshold   *int         `json:"meterThreshold"`
	MinutesThreshold *int         `json:"minutesThreshold"`
}
