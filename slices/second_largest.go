package slices

// SecondLargest returns the second largest element in a slice of ints or floats.
func SecondLargest[T int32 | int64 | float32 | float64](values []T) T {
	if len(values) < 2 {
		var zero T
		// Return the zero value for the type if there aren't enough elements return default value
		return zero 
	}
	
	firstLargest := values[0]
	secondLargest := values[0]

	for _, value := range values {
		if value > firstLargest {
			secondLargest = firstLargest
			firstLargest = value
		} else if value > secondLargest && value != firstLargest {
			secondLargest = value
		}
	}
	
	// Handle edge case where all elements are the same
	if firstLargest == secondLargest {
		var zero T
		return zero
	}

	return secondLargest
}
