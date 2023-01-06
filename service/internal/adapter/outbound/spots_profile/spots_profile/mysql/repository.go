package spots_profile_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile/spots_profile"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) spots_profile.Repository {
	return Repository{db: db}
}

func (r Repository) Get(i spots_profile.Identifier) (spots_profile.SpotsProfile, error) {
	res, err := r.BulkGet([]spots_profile.Identifier{i})
	if err != nil {
		return nil, err
	}
	return res[i], nil
}

func (r Repository) BulkGet(is []spots_profile.Identifier) (map[spots_profile.Identifier]spots_profile.SpotsProfile, error) {
	result := make(map[spots_profile.Identifier]spots_profile.SpotsProfile)

	if len(is) == 0 {
		return result, nil
	}

	var rows []SpotsProfile
	if err := r.db.
		Model(&SpotsProfile{}).
		Where("id in ?", is).
		Preload("SpotsProfileSpots", func(db *gorm.DB) *gorm.DB {
			return db.Order("spots_profile_spot.spot_id ASC")
		}).
		Find(&rows).
		Error; err != nil {
		return nil, err
	}

	for _, row := range rows {
		sp := unmarshal(row)
		result[sp.Identifier()] = sp
	}
	return result, nil
}

func (r Repository) Save(sp spots_profile.SpotsProfile) error {
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

func (r Repository) NextIdentifier() (spots_profile.Identifier, error) {
	row := SpotsProfile{}
	if err := r.db.Save(&row).Error; err != nil {
		return 0, err
	}
	if err := r.db.Delete(row).Error; err != nil {
		return 0, err
	}

	return spots_profile.Identifier(row.ID), nil
}
