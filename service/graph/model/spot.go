package model

type Spot struct {
	Key         int     `json:"key"`
	PostalCode  string  `json:"postalCode"`
	AddressRepr string  `json:"addressRepr"`
	Lat         float64 `json:"lat"`
	Lng         float64 `json:"lng"`
}
