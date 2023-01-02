package spots_profile_mysql

type SpotsProfile struct {
	ID                uint
	SpotsProfileSpots []SpotsProfileSpot
}

type SpotsProfileSpot struct {
	SpotsProfileID uint `gorm:"primaryKey;autoIncrement:false"`
	SpotsID        uint `gorm:"primaryKey;autoIncrement:false"`
}
