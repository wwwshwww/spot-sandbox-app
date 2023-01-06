package dbscan_profile_finder_mysql

import (
	dbscan_profile_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/dbscan_profile/dbscan_profile/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
	"gorm.io/gorm"
)

type Finder struct {
	db *gorm.DB
}

func New(db *gorm.DB) Finder {
	return Finder{db: db}
}

func (f Finder) Find() (
	[]dbscan_profile.Identifier,
	error,
) {
	query := f.db.Model(&dbscan_profile_mysql.DbscanProfile{}).Order("id")

	var dpis []dbscan_profile.Identifier
	if err := query.Select("id").Find(&dpis).Error; err != nil {
		return nil, err
	}

	return dpis, nil
}
