package helpers

func Map[T any, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))

	for i, v := range input {
		result[i] = fn(v)
	}

	return result
}

func ForEachT[T any](input []T, fn func(T)) {
	for _, v := range input {
		fn(v)
	}
}
