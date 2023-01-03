package spot_finder

type SortingOptions struct {
	Key        Key
	Descending bool
}

type Key uint

const (
	Unspecified Key = iota
	SpotIdentifier
)
