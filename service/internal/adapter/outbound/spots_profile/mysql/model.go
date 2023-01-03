package spots_profile_mysql

type SpotsProfile struct {
	ID                uint
	SpotsProfileSpots []SpotsProfileSpot
}

type SpotsProfileSpot struct {
	SpotsProfileID uint `gorm:"primaryKey;autoIncrement:false"`
	SpotID         uint `gorm:"primaryKey;autoIncrement:false"`
}
