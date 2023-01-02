package spots_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) spots_profile.Repository {
	return Repository{db: db}
}

func (r Repository) Get(i spots_profile.Identifier) (spots_profile.SpotProfile, error) {
	var row SpotsProfile

	if err := r.db.
		Preload("SpotsProfileSpots").
		Model(&SpotsProfile{}).
		Where("id = ?", i).
		First(&row).
		Error; err != nil {
		return nil, err
	}

	return unmarshal(row), nil
}

func (r Repository) Save(sp spots_profile.SpotProfile) error {
	if err := r.db.
		Model(&SpotsProfileSpot{}).
		Where("spots_profile_id = ?", sp.Identifier()).
		Delete(&SpotsProfileSpot{}).
		Error; err != nil {
		return err
	}
	if err := r.db.
		Where("id = ?", sp.Identifier()).
		Save(marshal(sp)).
		Error; err != nil {
		return err
	}
	return nil
}

func (r Repository) Delete(i spots_profile.Identifier) error {
	if err := r.db.
		Model(&SpotsProfileSpot{}).
		Where("spots_profile_id = ?", i).
		Delete(&SpotsProfileSpot{}).
		Error; err != nil {
		return err
	}
	if err := r.db.
		Model(&SpotsProfile{}).
		Where("id = ?", i).
		Delete(&SpotsProfile{}).
		Error; err != nil {
		return err
	}
	return nil
}
