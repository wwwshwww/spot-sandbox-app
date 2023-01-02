package address

type Resolver interface {
	ReverseGeocode(lat, lng float64) (postalCode, addressRepresentation string, err error)
}
