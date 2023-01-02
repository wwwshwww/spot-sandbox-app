package dbscan_profile

type Identifier uint

func (i Identifier) Value() uint {
	return uint(i)
}
