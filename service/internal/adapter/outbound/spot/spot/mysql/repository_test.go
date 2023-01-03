package spot_mysql_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"

	"github.com/wwwwshwww/spot-sandbox/internal/adapter/gateway/rdb"
	spot_mysql "github.com/wwwwshwww/spot-sandbox/internal/adapter/outbound/spot/spot/mysql"
	"github.com/wwwwshwww/spot-sandbox/internal/domain/spot/spot"
)

const testDB = "spot_test"

func TestGet(t *testing.T) {
	db, closeDB, err := rdb.NewMySQLInstance(testDB, &spot_mysql.Spot{})
	defer func() {
		_ = closeDB()
	}()
	assert.NoError(t, err)

	db = prepareDB(t, db)
	repo := spot_mysql.New(db)

	tests := []struct {
		i        spot.Identifier
		expected spot.Spot
		isErr    bool
	}{
		{
			spot.Identifier(1),
			spot.Restore(
				spot.Identifier(1),
				spot.SpotPreferences{
					PostalCode:            "1000000",
					AddressRepresentation: "とうきょ",
					Lat:                   0.11,
					Lng:                   0.11,
				},
			),
			false,
		},
		{
			spot.Identifier(999),
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
	db, closeDB, err := rdb.NewMySQLInstance(testDB, &spot_mysql.Spot{})
	defer func() {
		_ = closeDB()
	}()
	assert.NoError(t, err)

	db = prepareDB(t, db)
	repo := spot_mysql.New(db)

	tests := []struct {
		s        spot.Spot
		expected spot.Spot
		isErr    bool
	}{
		{
			spot.Restore(
				spot.Identifier(4),
				spot.SpotPreferences{
					PostalCode:            "4000000",
					AddressRepresentation: "とうきょ",
					Lat:                   0.44,
					Lng:                   0.44,
				},
			),
			spot.Restore(
				spot.Identifier(4),
				spot.SpotPreferences{
					PostalCode:            "4000000",
					AddressRepresentation: "ほっかいど",
					Lat:                   0.45,
					Lng:                   0.45,
				},
			),
			false,
		},
	}

	for _, tt := range tests {
		err := repo.Save(tt.s)
		if tt.isErr {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
		actual, err := repo.Get(tt.s.Identifier())
		assert.Nil(t, err)
		assert.Equal(t, tt.s, actual)
	}
}

func prepareDB(t *testing.T, db *gorm.DB) *gorm.DB {

	defaultSpot := []spot_mysql.Spot{
		{
			ID:                    1,
			PostalCode:            "1000000",
			AddressRepresentation: "とうきょ",
			Lat:                   0.11,
			Lng:                   0.11,
		},
		{
			ID:                    2,
			PostalCode:            "2000000",
			AddressRepresentation: "きょうと",
			Lat:                   0.22,
			Lng:                   0.22,
		},
	}
	err := db.Save(defaultSpot).Error
	assert.NoError(t, err)

	return db
}
