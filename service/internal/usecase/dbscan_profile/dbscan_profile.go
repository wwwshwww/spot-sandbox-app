package dbscan_profile

import "github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"

type DbscanProfile interface {
	Get(dbscan_profile.Identifier) (dbscan_profile.DbscanProfile, error)
	Save(dbscan_profile.Identifier, dbscan_profile.DbscanProfilePreferences) error
}
