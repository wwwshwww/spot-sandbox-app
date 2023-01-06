package spots_profile_finder_mysql

import (
	spots_profile_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spots_profile/spots_profile/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile/spots_profile"
	"gorm.io/gorm"
)

type Finder struct {
	db *gorm.DB
}

func New(db *gorm.DB) Finder {
	return Finder{db: db}
}

func (f Finder) Find() (
	[]spots_profile.Identifier,
	error,
) {
	query := f.db.Model(&spots_profile_mysql.SpotsProfile{}).Order("id")

	var spis []spots_profile.Identifier
	if err := query.Select("id").Find(&spis).Error; err != nil {
		return nil, err
	}

	return spis, nil
}
