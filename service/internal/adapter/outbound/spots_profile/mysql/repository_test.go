package spots_profile_mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/rdb"
	spots_profile_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spots_profile/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spots_profile"
	"gorm.io/gorm"
)

const testDB = "spots_profile_test"

func TestGet(t *testing.T) {
	db, closeDB, err := rdb.NewMySQLInstance(testDB, &spots_profile_mysql.SpotsProfile{}, &spots_profile_mysql.SpotsProfileSpot{})
	defer func() {
		_ = closeDB()
	}()
	assert.NoError(t, err)

	db = prepareDB(t, db)
	repo := spots_profile_mysql.New(db)

	var tests = []struct {
		i     spots_profile.Identifier
		sp    spots_profile.SpotsProfile
		isErr bool
	}{
		{
			spots_profile.Identifier(1),
			spots_profile.Restore(
				spots_profile.Identifier(1),
				spots_profile.SpotsProfilePreferences{
					Spots: []spot.Identifier{11, 12, 13},
				},
			),
			false,
		},
		{
			spots_profile.Identifier(999),
			nil,
			false,
		},
	}

	for _, tt := range tests {
		sp, err := repo.Get(tt.i)
		if tt.isErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, tt.sp, sp)
	}
}

func prepareDB(t *testing.T, db *gorm.DB) *gorm.DB {
	defaultSpotsProfileDB := []spots_profile_mysql.SpotsProfile{
		{
			ID: 1,
			SpotsProfileSpots: []spots_profile_mysql.SpotsProfileSpot{
				{
					SpotsProfileID: 1,
					SpotID:         12,
				},
				{
					SpotsProfileID: 1,
					SpotID:         13,
				},
				{
					SpotsProfileID: 1,
					SpotID:         11,
				},
			},
		},
	}
	err := db.Save(defaultSpotsProfileDB).Error
	assert.NoError(t, err)

	return db
}
