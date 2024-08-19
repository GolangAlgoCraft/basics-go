package slices

// Returns the sum of all elements in the slice
func Sum[T int | int32 | int64 | float32 | float64 | string](values []T) T {
	var sum T
	for _, value := range values {
		// Add the value to the sum
		sum += value
	}

	// Return the sum
	return sum
}
