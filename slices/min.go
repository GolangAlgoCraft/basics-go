package slices

// Returns the minimum value from a slice of integers
func Min(slice []int) int {
	if len(slice) == 0 {
		// return 0 if the slice is empty
		return 0
	}

	// initialize the minimum value to the first element of the slice
	min := slice[0]
	for _, value := range slice[1:] {
		// update the minimum value if a smaller value is found
		if value < min {
			min = value
		}
	}

	// return the minimum value
	return min
}
