package spot_mysql

import (
	"github.com/wwwwshwww/spot-sandbox/internal/common"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func New(db *gorm.DB) spot.Repository {
	return Repository{db: db}
}

func (r Repository) Get(i spot.Identifier) (spot.Spot, error) {
	result, err := r.BulkGet([]spot.Identifier{i})
	if err != nil {
		return nil, err
	}

	s, ok := result[i]
	if !ok {
		return nil, spot.ErrNotFound
	}
	return s, nil
}

func (r Repository) BulkGet(is []spot.Identifier) (map[spot.Identifier]spot.Spot, error) {
	result := make(map[spot.Identifier]spot.Spot)

	if len(is) == 0 {
		return result, nil
	}

	var rows []Spot
	if err := r.db.
		Model(&Spot{}).
		Where("id in ?", is).
		Find(&rows).
		Error; err != nil {
		return nil, err
	}

	if len(rows) == 0 {
		return result, nil
	}

	for _, row := range rows {
		s := unmarshal(row)
		result[s.Identifier()] = s
	}

	return result, nil
}

func (r Repository) Save(s spot.Spot) error {
	if err := r.BulkSave([]spot.Spot{s}); err != nil {
		return err
	}
	return nil
}

func (r Repository) BulkSave(ss []spot.Spot) error {
	if len(ss) == 0 {
		return nil
	}

	spots := common.Map(marshal, ss)
	if err := r.db.
		Model(&Spot{}).
		Save(&spots).
		Error; err != nil {
		return err
	}

	return nil
}

func (r Repository) Delete(i spot.Identifier) error {
	if err := r.BulkDelete([]spot.Identifier{i}); err != nil {
		return err
	}
	return nil
}

func (r Repository) BulkDelete(is []spot.Identifier) error {
	if len(is) == 0 {
		return nil
	}

	if err := r.db.
		Model(&Spot{}).
		Where("id in ?", is).
		Delete(&Spot{}).
		Error; err != nil {
		return err
	}

	return nil
}
