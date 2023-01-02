package dbscan_profile

import (
	"time"
)

var (
	defaultThresholdMeter int          = 200
	defaultDistanceType   DistanceType = RouteLength
	defaultMinCount       uint         = 1
	defaultMaxCount       uint         = 20
)

// DistanceTypeに応じて保持するThresholdが変わる
type DbscanProfile interface {
	Identifier() Identifier
	DistanceType() DistanceType
	MinCount() uint
	MaxCount() *uint
	MeterThreshold() *int
	DurationThreshold() *time.Duration

	OverWrite(DbscanProfilePreferences)
}

type DbscanProfilePreferences struct {
	DistanceType      DistanceType
	MinCount          uint
	MaxCount          *uint
	MeterThreshold    *int
	DurationThreshold *time.Duration
}

func New(i Identifier) DbscanProfile {
	dp := &dbscanProfile{
		identifier:     i,
		distanceType:   defaultDistanceType,
		minCount:       defaultMinCount,
		maxCount:       &defaultMaxCount,
		meterThreshold: &defaultThresholdMeter,
	}
	return dp
}

func Restore(i Identifier, p DbscanProfilePreferences) DbscanProfile {
	dp := New(i)
	dp.OverWrite(p)
	return dp
}

type dbscanProfile struct {
	identifier        Identifier
	distanceType      DistanceType
	minCount          uint
	maxCount          *uint
	meterThreshold    *int
	durationThreshold *time.Duration
}

func (dp *dbscanProfile) OverWrite(p DbscanProfilePreferences) {
	if p.MaxCount != nil && (*p.MaxCount < 1 || *p.MaxCount < p.MinCount) {
		return
	}
	dp.maxCount = p.MaxCount
	dp.minCount = p.MinCount

	dp.distanceType = p.DistanceType
	switch p.DistanceType {
	case Hubeny, RouteLength:
		dp.meterThreshold = p.MeterThreshold
		dp.durationThreshold = nil
	case TravelTime:
		dp.durationThreshold = p.DurationThreshold
		dp.meterThreshold = nil
	}
}

func (dp *dbscanProfile) Identifier() Identifier            { return dp.identifier }
func (dp *dbscanProfile) DistanceType() DistanceType        { return dp.distanceType }
func (dp *dbscanProfile) MinCount() uint                    { return dp.minCount }
func (dp *dbscanProfile) MaxCount() *uint                   { return dp.maxCount }
func (dp *dbscanProfile) MeterThreshold() *int              { return dp.meterThreshold }
func (dp *dbscanProfile) DurationThreshold() *time.Duration { return dp.durationThreshold }
