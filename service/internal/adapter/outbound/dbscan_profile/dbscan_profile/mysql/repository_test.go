package dbscan_profile_mysql_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/rdb"
	dbscan_profile_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/dbscan_profile/dbscan_profile/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/dbscan_profile/dbscan_profile"
)

const testDB = "dbscan_profile_test"

func TestGet(t *testing.T) {
	db, closeDB, err := rdb.NewMySQLInstance(testDB, &dbscan_profile_mysql.DbscanProfile{})
	defer func() {
		_ = closeDB()
	}()
	assert.NoError(t, err)

	db = prepareDB(t, db)
	repo := dbscan_profile_mysql.New(db)

	tests := []struct {
		i        dbscan_profile.Identifier
		expected dbscan_profile.DbscanProfile
		isErr    bool
	}{
		{
			dbscan_profile.Identifier(1),
			dbscan_profile.Restore(
				1,
				dbscan_profile.DbscanProfilePreferences{
					DistanceType:   dbscan_profile.RouteLength,
					MinCount:       1,
					MeterThreshold: func(n int) *int { return &n }(10),
				},
			),
			false,
		},
		{
			dbscan_profile.Identifier(2),
			dbscan_profile.Restore(
				2,
				dbscan_profile.DbscanProfilePreferences{
					DistanceType:      dbscan_profile.TravelTime,
					MinCount:          1,
					DurationThreshold: func(n time.Duration) *time.Duration { return &n }(time.Hour * 6),
				},
			),
			false,
		},
		{
			dbscan_profile.Identifier(999),
			nil,
			false,
		},
	}

	for _, tt := range tests {
		s, err := repo.Get(tt.i)
		if tt.isErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		assert.Equal(t, tt.expected, s)
	}
}

func TestSave(t *testing.T) {
	db, closeDB, err := rdb.NewMySQLInstance(testDB, &dbscan_profile_mysql.DbscanProfile{})
	defer func() {
		_ = closeDB()
	}()
	assert.NoError(t, err)

	repo := dbscan_profile_mysql.New(db)

	data := dbscan_profile.Restore(
		1,
		dbscan_profile.DbscanProfilePreferences{
			DistanceType:   dbscan_profile.RouteLength,
			MinCount:       1,
			MeterThreshold: func(n int) *int { return &n }(10),
		},
	)
	err = repo.Save(data)
	assert.NoError(t, err)

	actual, err := repo.Get(data.Identifier())
	assert.NoError(t, err)
	assert.Equal(t, data, actual)
}

func prepareDB(t *testing.T, db *gorm.DB) *gorm.DB {

	defaultSpot := []dbscan_profile_mysql.DbscanProfile{
		{
			ID:             1,
			DistanceType:   dbscan_profile_mysql.RouteLength,
			MinCount:       1,
			MeterThreshold: func(n int) *int { return &n }(10),
		},
		{
			ID:                2,
			DistanceType:      dbscan_profile_mysql.TravelTime,
			MinCount:          1,
			DurationThreshold: func(n time.Duration) *time.Duration { return &n }(time.Hour * 6),
		},
	}
	err := db.Save(defaultSpot).Error
	assert.NoError(t, err)

	return db
}
