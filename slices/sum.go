package slices

// Returns the sum of all elements in the slice
func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		// Add the number to the sum
		sum += number
	}

	// Return the sum
	return sum
}
