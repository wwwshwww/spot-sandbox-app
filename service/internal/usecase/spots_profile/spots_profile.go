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

func (u spotsProfileUsecase) Get(spi spots_profile.Identifier) (
	spots_profile.SpotsProfile,
	error,
) {
	sp, err := u.spr.Get(spi)
	if err != nil {
		return nil, err
	}
	return sp, nil
}

func (u spotsProfileUsecase) Save(
	spi spots_profile.Identifier,
	spp spots_profile.SpotsProfilePreferences,
) error {
	sp, err := u.spr.Get(spi)
	if err != nil {
		return err
	}
	if sp == nil {
		sp = spots_profile.New(spi)
	}
	if err := sp.Overwrite(spp); err != nil {
		return err
	}
	if err := u.spr.Save(sp); err != nil {
		return err
	}
	return nil
}
