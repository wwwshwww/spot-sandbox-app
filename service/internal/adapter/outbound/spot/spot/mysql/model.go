package spot_mysql

type Spot struct {
	ID                    uint
	PostalCode            string
	AddressRepresentation string
	Lat                   float64
	Lng                   float64
}
