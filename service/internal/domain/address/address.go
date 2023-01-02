package address

import (
	"errors"
	"regexp"
)

var (
	RegexpPostalCode = regexp.MustCompile(`^[0-9]{3}-?[0-9]{4}$`) // 日本限定

	ErraddressInvalid        = errors.New("invalid address")
	ErrPostalCodeInvalid     = errors.New("invalid postal code")
	ErrPrefectureInvalid     = errors.New("invalid prefecture")
	ErrCityLengthExceeded    = errors.New("city length exceeds the limit")
	ErrStreet1LengthExceeded = errors.New("street1 length exceeds the limit")
	ErrStreet2LengthExceeded = errors.New("street2 length exceeds the limit")
	ErrLatLngMustBoth        = errors.New("both lat and lng must be set")
)

type Address interface {
	PostalCode() string
	AddressRepresentation() string
	Lat() float64
	Lng() float64

	UpdatePostalCode(string) error
	UpdateAddressReporesentation(string) error
	UpdateLatLng(lat, lng float64) error

	CorrectAddressRepresentationByLatLng(func(lat, lng float64) (postalCode, addressRepresentation string, err error)) error

	Equals(Address) bool
	String() string
}

type address struct {
	postalCode            string
	addressRepresentation string
	lat                   float64
	lng                   float64
}

func New(
	postalCode, addressRepresentation string,
	lat, lng float64,
) (Address, error) {
	a := address{}

	if err := a.UpdatePostalCode(postalCode); err != nil {
		return nil, err
	}
	if err := a.UpdateAddressReporesentation(addressRepresentation); err != nil {
		return nil, err
	}
	if err := a.UpdateLatLng(lat, lng); err != nil {
		return nil, err
	}

	return &a, nil
}

func (t address) PostalCode() string            { return t.postalCode }
func (t address) AddressRepresentation() string { return t.addressRepresentation }
func (t address) Lat() float64                  { return t.lat }
func (t address) Lng() float64                  { return t.lng }

func (t address) String() string {
	return t.addressRepresentation
}

func (t *address) UpdatePostalCode(postalCode string) error {
	t.postalCode = postalCode
	return nil
}

func (t *address) UpdateAddressReporesentation(addressRepresentation string) error {
	t.addressRepresentation = addressRepresentation
	return nil
}

func (t *address) UpdateLatLng(lat, lng float64) error {
	t.lat = lat
	t.lng = lng
	return nil
}

func (t address) Equals(a Address) bool {
	return t.postalCode == a.PostalCode() &&
		t.addressRepresentation == a.AddressRepresentation() &&
		t.lat == a.Lat() &&
		t.lng == a.Lng()
}

func (t *address) CorrectAddressRepresentationByLatLng(
	rslvFn func(lat, lng float64) (postalCode, addressRepresentation string, err error),
) error {
	postalCode, addressRepresentation, err := rslvFn(t.lat, t.lng)
	if err != nil {
		return err
	}
	if err := t.UpdatePostalCode(postalCode); err != nil {
		return err
	}
	if err := t.UpdateAddressReporesentation(addressRepresentation); err != nil {
		return err
	}
	return nil
}
