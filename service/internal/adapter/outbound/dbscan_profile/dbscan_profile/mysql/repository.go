package dbscan_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) dbscan_profile.Repository {
	return Repository{db: db}
}

func (r Repository) Get(i dbscan_profile.Identifier) (dbscan_profile.DbscanProfile, error) {
	res, err := r.BulkGet([]dbscan_profile.Identifier{i})
	if err != nil {
		return nil, err
	}
	return res[i], nil
}

func (r Repository) BulkGet(is []dbscan_profile.Identifier) (map[dbscan_profile.Identifier]dbscan_profile.DbscanProfile, error) {
	result := make(map[dbscan_profile.Identifier]dbscan_profile.DbscanProfile)

	if len(is) == 0 {
		return result, nil
	}

	var rows []DbscanProfile
	if err := r.db.
		Model(&DbscanProfile{}).
		Where("id in ?", is).
		Find(&rows).
		Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		dp := unmarshal(row)
		result[dp.Identifier()] = dp
	}
	return result, nil
}

func (r Repository) Save(dp dbscan_profile.DbscanProfile) error {
	if err := r.db.
		Model(&DbscanProfile{}).
		Where("id = ?", dp.Identifier()).
		Save(marshal(dp)).
		Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) Delete(i dbscan_profile.Identifier) error {
	if err := r.db.
		Model(&DbscanProfile{}).
		Where("id = ?").
		Delete(&DbscanProfile{}).
		Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) NextIdentifier() (dbscan_profile.Identifier, error) {
	row := DbscanProfile{}
	if err := r.db.Save(&row).Error; err != nil {
		return 0, err
	}
	if err := r.db.Delete(row).Error; err != nil {
		return 0, err
	}

	return dbscan_profile.Identifier(row.ID), nil
}
