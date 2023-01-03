package dbscan_profile

import "github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"

type DbscanProfileUsecase interface {
	Get(dbscan_profile.Identifier) (dbscan_profile.DbscanProfile, error)
	Save(dbscan_profile.Identifier, dbscan_profile.DbscanProfilePreferences) error
}

func New(dpr dbscan_profile.Repository) DbscanProfileUsecase {
	return dbscanProfileUsecase{dpr: dpr}
}

type dbscanProfileUsecase struct {
	dpr dbscan_profile.Repository
}

func (u dbscanProfileUsecase) Get(dpi dbscan_profile.Identifier) (
	dbscan_profile.DbscanProfile,
	error,
) {
	dp, err := u.dpr.Get(dpi)
	if err != nil {
		return nil, err
	}
	return dp, nil
}

func (u dbscanProfileUsecase) Save(
	dpi dbscan_profile.Identifier,
	dpp dbscan_profile.DbscanProfilePreferences,
) error {
	// dp, err := u.dpr.Get(dpi)
	return nil
}
