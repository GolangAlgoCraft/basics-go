package slices

func CheckSorted[T int32 | int64 | float32 | float64 | string](values []T) bool {
	if len(values) <= 1 {
		return true
	}
	for i := 1; i < len(values); i++ {
		if values[i] < values[i-1] {
			return false
		}
	}
	return true
}