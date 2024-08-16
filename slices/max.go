package slices

// Returns the maximum value from a slice of integers
func Max(slice []int) int {
	// initialize the maximum value with 0, So that if the slice is empty, it will return 0
	max := 0
	for _, value := range slice {
		// update the maximum value if the current value is greater than the maximum value
		if value > max {
			max = value
		}
	}

	// return the maximum value
	return max
}
