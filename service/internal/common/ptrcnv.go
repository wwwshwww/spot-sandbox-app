package common

func P2V[T any](in *T) T {
	var out T
	if in == nil {
		return out
	} else {
		out = *in
		return out
	}
}
