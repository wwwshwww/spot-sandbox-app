package common

type number interface {
	~int | ~int8 | ~int16 | ~int32 | int64 | ~uint | ~uint32 | ~uint64 | ~float32 | ~float64
}

func Max[T number](n1, n2 T) T {
	if n1 > n2 {
		return n1
	} else {
		return n2
	}
}

func Min[T number](n1, n2 T) T {
	if n1 < n2 {
		return n1
	} else {
		return n2
	}
}
