package spot_finder_mysql

import (
	spot_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spot/spot/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot_finder"
	"gorm.io/gorm"
)

type Finder struct {
	db *gorm.DB
}

func New(db *gorm.DB) Finder {
	return Finder{db: db}
}

func (f Finder) Find(
	fo spot_finder.FilteringOptions,
	so spot_finder.SortingOptions,
) (
	[]spot.Identifier,
	error,
) {
	query := f.db.Model(&spot_mysql.Spot{})
	if fo.PostalCode != "" {
		query = query.Where("postal_code = ?", fo.PostalCode)
	}
	if fo.SpotIdentifiers != nil && len(fo.SpotIdentifiers) > 0 {
		query = query.Where("id in ?", fo.SpotIdentifiers)
	}

	switch so.Key {
	case spot_finder.SpotIdentifier:
		if so.Descending {
			query = query.Order("id desc")
		} else {
			query = query.Order("id")
		}
	}

	var sis []spot.Identifier
	if err := query.Select("id").Find(&sis).Error; err != nil {
		return nil, err
	}
	return sis, nil
}
