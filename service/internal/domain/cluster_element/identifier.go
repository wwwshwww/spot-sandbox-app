package cluster_element

type Identifier uint

func (i Identifier) Value() uint {
	return uint(i)
}
