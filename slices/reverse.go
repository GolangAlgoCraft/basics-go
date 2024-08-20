package slices

// Reverse reverses the order of elements in a slice.
func Reverse[T int | int32 | int64 | float32 | float64 | string](slice []T) []T {
	for start, end := 0, len(slice)-1; start < end; start, end = start+1, end-1 {
		// swap the elements in the slice using the last element at the end with the start on the front
		slice[start], slice[end] = slice[end], slice[start]
	}

	return slice
}
