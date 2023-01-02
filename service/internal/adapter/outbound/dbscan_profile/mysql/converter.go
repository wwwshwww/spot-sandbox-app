package dbscan_profile_mysql

import "github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"

func unmarshal(row DbscanProfile) dbscan_profile.DbscanProfile {
	return dbscan_profile.Restore(
		dbscan_profile.Identifier(row.ID),
		dbscan_profile.DbscanProfilePreferences{
			DistanceType:      unmarshalDistanceType(row.DistanceType),
			MinCount:          row.MinCount,
			MaxCount:          row.MaxCount,
			MeterThreshold:    row.MeterThreshold,
			DurationThreshold: row.DurationThreshold,
		},
	)
}

func marshal(dp dbscan_profile.DbscanProfile) DbscanProfile {
	return DbscanProfile{
		ID:                dp.Identifier().Value(),
		DistanceType:      marshalDistanceType(dp.DistanceType()),
		MinCount:          dp.MinCount(),
		MaxCount:          dp.MaxCount(),
		MeterThreshold:    dp.MeterThreshold(),
		DurationThreshold: dp.DurationThreshold(),
	}
}

func unmarshalDistanceType(dt DistanceType) dbscan_profile.DistanceType {
	switch dt {
	case Hubeny:
		return dbscan_profile.Hubeny
	case RouteLength:
		return dbscan_profile.RouteLength
	case TravelTime:
		return dbscan_profile.TravelTime
	default:
		panic("unknown distance type")
	}
}

func marshalDistanceType(dt dbscan_profile.DistanceType) DistanceType {
	switch dt {
	case dbscan_profile.Hubeny:
		return Hubeny
	case dbscan_profile.RouteLength:
		return RouteLength
	case dbscan_profile.TravelTime:
		return TravelTime
	default:
		panic("unknown distance type")
	}
}
