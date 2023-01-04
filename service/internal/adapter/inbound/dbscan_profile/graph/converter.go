package dbscan_profile_graph

import (
	"strconv"
	"time"

	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
)

func Marshal(dp dbscan_profile.DbscanProfile) model.DbscanProfile {
	return model.DbscanProfile{
		ID:           strconv.Itoa(int(dp.Identifier())),
		DistanceType: MarshalDistanceType(dp.DistanceType()),
		MinCount:     int(dp.MinCount()),
		MaxCount: func(n *uint) *int {
			if n == nil {
				return nil
			} else {
				nn := int(*n)
				return &nn
			}
		}(dp.MaxCount()),
		MeterThreshold: dp.MeterThreshold(),
		MinutesThreshold: func(d *time.Duration) *int {
			if d == nil {
				return nil
			} else {
				m := int(d.Minutes())
				return &m
			}
		}(dp.DurationThreshold()),
	}
}

func MarshalDistanceType(dt dbscan_profile.DistanceType) model.DistanceType {
	switch dt {
	case dbscan_profile.Hubeny:
		return model.DistanceTypeHubeny
	case dbscan_profile.RouteLength:
		return model.DistanceTypeRouteLength
	case dbscan_profile.TravelTime:
		return model.DistanceTypeTravelTime
	default:
		panic("なに渡しとんねんワレ")
	}
}

func UnmarshalPreferences(m model.NewDbscanProfile) dbscan_profile.DbscanProfilePreferences {
	return dbscan_profile.DbscanProfilePreferences{
		DistanceType: UnmarshalDistanceType(m.DistanceType),
		MinCount:     uint(m.MinCount),
		MaxCount: func(n *int) *uint {
			if n == nil {
				return nil
			} else {
				nn := uint(*n)
				return &nn
			}
		}(m.MaxCount),
		MeterThreshold: m.MeterThreshold,
		DurationThreshold: func(n *int) *time.Duration {
			if n == nil {
				return nil
			} else {
				d := time.Duration(time.Minute * time.Duration(*n))
				return &d
			}
		}(m.MinutesThreshold),
	}
}

func UnmarshalDistanceType(m model.DistanceType) dbscan_profile.DistanceType {
	switch m {
	case model.DistanceTypeHubeny:
		return dbscan_profile.Hubeny
	case model.DistanceTypeRouteLength:
		return dbscan_profile.RouteLength
	case model.DistanceTypeTravelTime:
		return dbscan_profile.TravelTime
	default:
		panic("なに渡しとんねんワレ")
	}
}
