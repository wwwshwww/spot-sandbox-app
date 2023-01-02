package spots_profile

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"

type SpotProfileUsecase interface {
	Get(spots_profile.Identifier) (spots_profile.SpotsProfile, error)
	Save(spots_profile.Identifier, spots_profile.SpotsProfilePreferences) error
}
