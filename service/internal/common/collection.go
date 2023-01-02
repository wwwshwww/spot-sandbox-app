package common

func Contain[T comparable](src []T, elem T) bool {
	for _, e := range src {
		if e == elem {
			return true
		}
	}
	return false
}
