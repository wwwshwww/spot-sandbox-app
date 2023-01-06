package dbscan_profile

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile_finder"
)

type DbscanProfileUsecase interface {
	Get(dbscan_profile.Identifier) (dbscan_profile.DbscanProfile, error)
	Save(dbscan_profile.Identifier, dbscan_profile.DbscanProfilePreferences) error

	ListAllDbscanProfiles() ([]dbscan_profile.Identifier, error)
}

func New(
	dpr dbscan_profile.Repository,
	dpf dbscan_profile_finder.Finder,
) DbscanProfileUsecase {
	return dbscanProfileUsecase{dpr: dpr, dpf: dpf}
}

type dbscanProfileUsecase struct {
	dpr dbscan_profile.Repository
	dpf dbscan_profile_finder.Finder
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
	dp, err := u.dpr.Get(dpi)
	if err != nil {
		return err
	}
	if dp == nil {
		dp = dbscan_profile.New(dpi)
	}
	if err := dp.Overwrite(dpp); err != nil {
		return err
	}
	if err := u.dpr.Save(dp); err != nil {
		return err
	}
	return nil
}

func (u dbscanProfileUsecase) ListAllDbscanProfiles() (
	[]dbscan_profile.Identifier,
	error,
) {
	dpis, err := u.dpf.Find()
	if err != nil {
		return nil, err
	}
	return dpis, nil
}
