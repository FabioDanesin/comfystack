package functional

func Map[T, V any](array []T, fn func(T) V) []V {
	arrlen := len(array)
	result := make([]V, arrlen)

	for i, t := range array {
		result[i] = fn(t)
	}

	return result
}

func Fold[T, V any](array []T, fn func(V, T) V, initial V) V {
	if len(array) == 0 {
		return initial
	}

	accumulator := initial

	for _, val := range array {
		accumulator = fn(accumulator, val)
	}

	return accumulator
}
