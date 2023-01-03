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

	return result[i], nil
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

func (r Repository) NextIdentifier() (spot.Identifier, error) {
	ids, err := r.NextIdentifiers(1)
	if err != nil {
		return 0, err
	}

	if len(ids) != 1 {
		panic("NextIdentifiers didn't provide exactly 1 identifier")
	}

	return ids[0], nil
}

func (r Repository) NextIdentifiers(n uint) ([]spot.Identifier, error) {
	if n == 0 {
		return nil, nil
	}

	rows := make([]Spot, n)
	if err := r.db.Save(&rows).Error; err != nil {
		return nil, err
	}
	if err := r.db.Save(rows).Error; err != nil {
		return nil, err
	}

	ids := make([]spot.Identifier, n)
	for i, row := range rows {
		ids[i] = spot.Identifier(row.ID)
	}

	return ids, nil
}
