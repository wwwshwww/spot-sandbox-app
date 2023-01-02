package spots_profile

import "github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"

type SpotProfileUsecase interface {
	Get(spots_profile.Identifier) (spots_profile.SpotsProfile, error)
	Save(spots_profile.Identifier, spots_profile.SpotsProfilePreferences) error
}

func New(spr spots_profile.Repository) SpotProfileUsecase {
	return spotsProfileUsecase{spr: spr}
}

type spotsProfileUsecase struct {
	spr spots_profile.Repository
}

func (u spotsProfileUsecase) Get(si spots_profile.Identifier) (
	spots_profile.SpotsProfile,
	error,
) {
	return nil, nil
}

func (u spotsProfileUsecase) Save(
	si spots_profile.Identifier,
	sp spots_profile.SpotsProfilePreferences,
) error {
	return nil
}
