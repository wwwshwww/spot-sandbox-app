package common

func Map[I, O any](fn func(I) O, input []I) []O {
	result := make([]O, len(input))
	for i := range input {
		result[i] = fn(input[i])
	}
	return result
}
