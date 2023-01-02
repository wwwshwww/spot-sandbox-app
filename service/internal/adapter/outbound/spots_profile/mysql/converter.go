package spots_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

func unmarshal(row SpotsProfile) spots_profile.SpotProfile {
	return spots_profile.Restore(
		spots_profile.Identifier(row.ID),
		func(sps []SpotsProfileSpot) []spot.Identifier {
			spots := make([]spot.Identifier, len(sps))
			for i := range sps {
				spots[i] = spot.Identifier(sps[i].SpotsID)
			}
			return spots
		}(row.SpotsProfileSpots),
	)
}

func marshal(sp spots_profile.SpotProfile) SpotsProfile {
	return SpotsProfile{
		ID: sp.Identifier().Value(),
		SpotsProfileSpots: func(spots []spot.Identifier) []SpotsProfileSpot {
			sps := make([]SpotsProfileSpot, len(spots))
			for i := range spots {
				sps[i] = SpotsProfileSpot{
					SpotsProfileID: sp.Identifier().Value(),
					SpotsID:        spots[i].Value(),
				}
			}
			return sps
		}(sp.Spots()),
	}
}
