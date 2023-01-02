package dbscan_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) dbscan_profile.Repository {
	return Repository{db: db}
}

func (r Repository) Get(i dbscan_profile.Identifier) (dbscan_profile.DbscanProfile, error) {
	var row DbscanProfile
	if err := r.db.
		Model(&row).
		Where("id = ?", i).
		First(&row).
		Error; err != nil {
		return nil, err
	}
	return unmarshal(row), nil
}

func (r Repository) Save(dp dbscan_profile.DbscanProfile) error {
	if err := r.db.
		Model(&DbscanProfile{}).
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
	if err := r.db.Save(row).Error; err != nil {
		return 0, err
	}

	return dbscan_profile.Identifier(row.ID), nil
}
