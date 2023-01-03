package spots_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
)

func unmarshal(row SpotsProfile) spots_profile.SpotsProfile {
	return spots_profile.Restore(
		spots_profile.Identifier(row.ID),
		spots_profile.SpotsProfilePreferences{
			Spots: func(sps []SpotsProfileSpot) []spot.Identifier {
				spots := make([]spot.Identifier, len(sps))
				for i := range sps {
					spots[i] = spot.Identifier(sps[i].SpotID)
				}
				return spots
			}(row.SpotsProfileSpots),
		},
	)
}

func marshal(sp spots_profile.SpotsProfile) SpotsProfile {
	return SpotsProfile{
		ID: sp.Identifier().Value(),
		SpotsProfileSpots: func(spots []spot.Identifier) []SpotsProfileSpot {
			sps := make([]SpotsProfileSpot, len(spots))
			for i := range spots {
				sps[i] = SpotsProfileSpot{
					SpotsProfileID: sp.Identifier().Value(),
					SpotID:         spots[i].Value(),
				}
			}
			return sps
		}(sp.Spots()),
	}
}
