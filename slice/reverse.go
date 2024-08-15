package slice


// Reverse reverses the order of elements in a slice of ints.

func ReverseElements(slice []int) {
 for start, end := 0, len(slice)-1; start < end; start, end = start+1, end-1 {

	// swap the elements in the slice using the last element at the end with the start on the front
	 slice[start], slice[end] = slice[end], slice[start]
 }
}
