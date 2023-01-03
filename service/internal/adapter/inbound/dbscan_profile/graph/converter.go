package dbscan_profile_graph

import (
	"github.com/wwwwshwww/spot-sandbox/graph/model"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
)

func Marshal(sp dbscan_profile.DbscanProfile) model.DbscanProfile {
	return model.DbscanProfile{}
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
