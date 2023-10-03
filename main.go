package helpers

func Map[T any, U any](input []T, fn func(T) U) []U {
	result := make([]U, len(input))

	for i, v := range input {
		result[i] = fn(v)
	}

	return result
}
