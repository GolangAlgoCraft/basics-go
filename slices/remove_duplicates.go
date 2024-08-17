package slices

// Removes duplicate values from a slice.
func RemoveDuplicates[T comparable](slice []T) []T {
	// Create a map to store the encountered values.
	encountered := map[T]bool{}
	// Create a slice to store the result.
	result := []T{}

	for v := range slice {
		// If the value has not been encountered before, add it to the map and the result slice.
		if !encountered[slice[v]] {
			encountered[slice[v]] = true
			result = append(result, slice[v])
		}
	}

	// Return the result.
	return result
}
